package service

import (
    "errors"
    "github.com/dgrijalva/jwt-go"
    "github.com/bidyutmondal/twitter-clone/auth-srv/internal/models"
    "github.com/bidyutmondal/twitter-clone/auth-srv/internal/repository"
    "golang.org/x/crypto/bcrypt"
    "time"
)

type AuthService struct {
    UserRepo *repository.UserRepository
    JWTSecret string
}

func (s *AuthService) Register(username, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := &models.User{
        Username: username,
        Password: string(hashedPassword),
    }

    return s.UserRepo.CreateUser(user)
}

func (s *AuthService) Login(username, password string) (string, error) {
    user, err := s.UserRepo.GetUserByUsername(username)
    if err != nil {
        return "", err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    return token.SignedString([]byte(s.JWTSecret))
}