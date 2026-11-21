package database

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func openTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	/*
		INTERACTION with TEST DB:

		run
			export TEST_DATABASE_URL="host=localhost user=test_user password=test_password dbname=post_db_test port=5433 sslmode=disable"
		on host machine first for env var to work
	*/
	sourceName := os.Getenv("TEST_DATABASE_URL")
	if sourceName == "" {
		t.Skip("TEST_DATABASE_URL is not set; skipping PostgreSQL integration test")
	}

	db, err := gorm.Open(postgres.Open(sourceName), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("failed to get underlying database connection: %v", err)
	}

	var databaseName string
	if err := db.Raw("SELECT current_database()").Scan(&databaseName).Error; err != nil {
		sqlDB.Close()
		t.Fatalf("failed to determine database name: %v", err)
	}

	/* Safety check: never run these tests against the development database. */
	if !strings.HasSuffix(databaseName, "_test") {
		sqlDB.Close()
		t.Fatalf("refusing to run tests against database %q: database name must end in _test", databaseName)
	}

	if err := MigrateUsers(db); err != nil {
		sqlDB.Close()
		t.Fatalf("failed to migrate test database: %v", err)
	}

	t.Cleanup(func() {
		if err := sqlDB.Close(); err != nil {
			t.Errorf("failed to close test database connection: %v", err)
		}
	})

	return db
}

func testEmail(t *testing.T) string {
	t.Helper()

	name := strings.ReplaceAll(t.Name(), "/", "-")
	return fmt.Sprintf("%s-%d@example.test", name, time.Now().UnixNano())
}

func createTestUser(t *testing.T, db *gorm.DB) *User {
	t.Helper()

	user := &User{
		Email:        testEmail(t),
		PasswordHash: "test-password-hash",
	}

	if err := CreateUser(db, user); err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	t.Cleanup(func() {
		if err := db.Unscoped().Delete(&User{}, user.ID).Error; err != nil {
			t.Errorf("failed to clean up test user: %v", err)
		}
	})

	return user
}

func TestCreateUser(t *testing.T) {
	db := openTestDB(t)

	email := testEmail(t)
	user := &User{
		Email:        " " + strings.ToUpper(email) + " ",
		PasswordHash: "test-password-hash",
	}

	if err := CreateUser(db, user); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	t.Cleanup(func() {
		if err := db.Unscoped().Delete(&User{}, user.ID).Error; err != nil {
			t.Errorf("failed to clean up test user: %v", err)
		}
	})

	if user.ID == 0 {
		t.Error("expected user ID to be populated")
	}

	if user.Email != strings.ToLower(email) {
		t.Errorf("expected normalized email %q, got %q", strings.ToLower(email), user.Email)
	}

	if user.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be populated")
	}

	if user.UpdatedAt.IsZero() {
		t.Error("expected UpdatedAt to be populated")
	}
}

func TestCreateUserRejectsInvalidInput(t *testing.T) {
	db := openTestDB(t)

	tests := []struct {
		name    string
		user    *User
		wantErr error
	}{
		{
			name:    "nil user",
			user:    nil,
			wantErr: ErrNilUser,
		},
		{
			name: "empty email",
			user: &User{
				PasswordHash: "test-hash",
			},
			wantErr: ErrEmptyEmail,
		},
		{
			name: "whitespace email",
			user: &User{
				Email:        "   ",
				PasswordHash: "test-hash",
			},
			wantErr: ErrEmptyEmail,
		},
		{
			name: "empty password hash",
			user: &User{
				Email: testEmail(t),
			},
			wantErr: ErrEmptyPassword,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := CreateUser(db, testCase.user)

			if !errors.Is(err, testCase.wantErr) {
				t.Errorf("expected error %v, got %v", testCase.wantErr, err)
			}
		})
	}
}

func TestCreateUserRejectsDuplicateEmail(t *testing.T) {
	db := openTestDB(t)

	firstUser := createTestUser(t, db)

	duplicate := &User{
		Email:        " " + strings.ToUpper(firstUser.Email) + " ",
		PasswordHash: "another-test-hash",
	}

	err := CreateUser(db, duplicate)
	if err == nil {
		t.Fatal("expected duplicate email insertion to fail")
	}
}

func TestGetUserByEmail(t *testing.T) {
	db := openTestDB(t)

	createdUser := createTestUser(t, db)

	email := " " + strings.ToUpper(createdUser.Email) + " "

	foundUser, err := GetUserByEmail(db, email)
	if err != nil {
		t.Fatalf("failed to retrieve user by email: %v", err)
	}

	if foundUser.ID != createdUser.ID {
		t.Errorf("expected user ID %d, got %d", createdUser.ID, foundUser.ID)
	}

	if foundUser.Email != createdUser.Email {
		t.Errorf("expected email %q, got %q", createdUser.Email, foundUser.Email)
	}
}

func TestGetUserByID(t *testing.T) {
	db := openTestDB(t)

	createdUser := createTestUser(t, db)

	foundUser, err := GetUserByID(db, createdUser.ID)
	if err != nil {
		t.Fatalf("failed to retrieve user by ID: %v", err)
	}

	if foundUser.ID != createdUser.ID {
		t.Errorf("expected user ID %d, got %d", createdUser.ID, foundUser.ID)
	}

	if foundUser.Email != createdUser.Email {
		t.Errorf("expected email %q, got %q", createdUser.Email, foundUser.Email)
	}
}

func TestGetUserByEmailNotFound(t *testing.T) {
	db := openTestDB(t)

	_, err := GetUserByEmail(db, testEmail(t))

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected gorm.ErrRecordNotFound, got %v", err)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	db := openTestDB(t)

	var maxID uint
	if err := db.Model(&User{}).
		Select("COALESCE(MAX(id), 0)").
		Scan(&maxID).Error; err != nil {
		t.Fatalf("failed to determine maximum user ID: %v", err)
	}

	_, err := GetUserByID(db, maxID+1)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected gorm.ErrRecordNotFound, got %v", err)
	}
}

func TestGetUserByIDZero(t *testing.T) {
	db := openTestDB(t)

	user, err := GetUserByID(db, 0)

	if user != nil {
		t.Errorf("expected nil user, got %v", user)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected gorm.ErrRecordNotFound, got %v", err)
	}
}

func TestGetUserByEmailRejectsEmptyEmail(t *testing.T) {
	db := openTestDB(t)

	_, err := GetUserByEmail(db, "   ")

	if !errors.Is(err, ErrEmptyEmail) {
		t.Errorf("expected ErrEmptyEmail, got %v", err)
	}
}
