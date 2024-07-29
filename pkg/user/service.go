package user

import (
	"sam/pkg/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(registerDTO *UserRegisterDTO, repo repository.Repository[User]) (*User, error) {
	u := &User{
		Email:    registerDTO.Email,
		Password: hashPassword(registerDTO.Password),
		Username: registerDTO.Username,
	}

	err := repo.Create(u)

	return u, err
}

func GetUser(id int, repo repository.Repository[User]) (*User, error) {
	return repo.Find(id)
}

func GetUserByUsernameAndPlainPassword(username, password string, repo repository.Repository[User]) (*User, error) {
	fs := []repository.Field{
		{
			Column: "username",
			Value:  username,
		},
	}

	u, err := repo.FindFirstBy(fs...)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, err
	}

	return u, nil
}

func GenerateTokenAndStoreUser(u *User, jwtTokenSecret string, repo repository.Repository[User]) (*JWTResponse, error) {

	expiresAt := time.Now().Add(time.Minute * 15)

	claims := &UserClaims{
		Username: u.Username,
		Email:    u.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}
	refreshClaims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 48).Unix(),
	}

	token, _ := (jwt.NewWithClaims(jwt.SigningMethodHS256, claims)).SignedString([]byte(jwtTokenSecret))
	refreshToken, _ := (jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)).SignedString([]byte(jwtTokenSecret))

	err := repo.Update(u, []repository.Field{
		{
			Column: "token",
			Value:  token,
		},
		{
			Column: "valid_token_until",
			Value:  expiresAt,
		},
	}...)

	if err != nil {
		return nil, err
	}

	return &JWTResponse{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}

func hashPassword(plain string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)

	return string(hash)
}
