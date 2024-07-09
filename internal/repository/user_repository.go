package repository

import (
    "database/sql"
    "github.com/bidyutmondal/twitter-clone/auth-srv/internal/models"
)

type UserRepository struct {
    DB *sql.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {
    _, err := r.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
    return err
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
    user := &models.User{}
    err := r.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        return nil, err
    }
    return user, nil
}