package database

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

var (
	ErrNilUser       = errors.New("user cannot be nil")
	ErrEmptyEmail    = errors.New("email cannot be empty")
	ErrEmptyPassword = errors.New("password hash cannot be empty")
	ErrInvalidRole   = errors.New("invalid user role")
)

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Email        string    `gorm:"not null;uniqueIndex;size:255"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         UserRole  `gorm:"type:varchar(10);not null;default:user"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

func MigrateUsers(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func CreateUser(db *gorm.DB, user *User) error {
	if user == nil {
		return ErrNilUser
	}

	user.Email = normalizeEmail(user.Email)

	if user.Email == "" {
		return ErrEmptyEmail
	}

	if user.PasswordHash == "" {
		return ErrEmptyPassword
	}

	/*
		This only assigns the 'user' role if no role has been provided by the calling function.
		If 'admin' as been assigned in the passed struct, the created user will be an admin.
		FIXME: Keep this potential security risk in mind for later implementation and authorization checks.
	*/
	if user.Role == "" {
		user.Role = RoleUser
	}

	if !isValidUserRole(user.Role) {
		return ErrInvalidRole
	}

	return db.Create(user).Error
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User

	email = normalizeEmail(email)
	if email == "" {
		return nil, ErrEmptyEmail
	}

	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByID(db *gorm.DB, id uint) (*User, error) {
	if id == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var user User

	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func isValidUserRole(role UserRole) bool {
	switch role {
	case RoleUser, RoleAdmin:
		return true
	default:
		return false
	}
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
