package handlers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"src/internal/auth"
	"src/internal/database"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		db: db,
	}
}

type registerRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerResponse struct {
	ID        uint              `json:"id"`
	Email     string            `json:"email"`
	Role      database.UserRole `json:"role"`
	CreatedAt time.Time         `json:"created_at"`
}

var emailValidator = validator.New()

func (handler *AuthHandler) Register(ctx *gin.Context) {
	var request registerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	email := strings.ToLower(strings.TrimSpace(request.Email))
	if err := emailValidator.Var(email, "required,email"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email address",
		})
		return
	}

	if err := auth.ValidatePassword(request.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "password must contain at least 8 characters",
		})
		return
	}

	passwordHash, err := auth.HashPassword(request.Password)
	request.Password = "" // plaintext password not needed from this point on -> drop ASAP
	if err != nil {
		if errors.Is(err, auth.ErrPasswordTooLong) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "password is too long",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to process password",
		})
		return
	}

	user := &database.User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	if err := database.CreateUser(handler.db, user); err != nil {
		if errors.Is(err, database.ErrDuplicateEmail) {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": "email address already registered",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, registerResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	})
}
