package user

import (
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username        string
	Password        string `json:"-"`
	Email           string
	Token           string
	ValidTokenUntil time.Time
}

type UserRegisterDTO struct {
	Username string
	Password string
	Email    string
}

type UserLoginDTO struct {
	Username string
	Password string
}

type UserClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type JWTResponse struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}
