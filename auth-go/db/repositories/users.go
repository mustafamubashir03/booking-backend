package db

import (
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImp struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	userRepository := &UserRepositoryImp{
		db: nil,
	}
	return userRepository

}
func (user *UserRepositoryImp) Create() error {
	fmt.Println("user repository hit")
	return nil
}
