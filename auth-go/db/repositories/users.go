package db

import (
	"auth-go/models"
	"database/sql"
	"fmt"
	"log"
)

type UserRepository interface {
	Create(name string, email string, hashedPassword string) error
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]models.User, error)
	DeleteById(id string) error
}

type UserRepositoryImp struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	userRepository := &UserRepositoryImp{
		db: _db,
	}
	return userRepository

}
func (userRepo *UserRepositoryImp) Create(name string, email string, hashedPassword string) error {
	fmt.Println("user repository hit")
	query := "INSERT INTO users(username, email, password) VALUES (?,?,?)"
	result, err := userRepo.db.Exec(query, name, email, hashedPassword)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	rowsAffected, rowErr := result.RowsAffected()
	if rowErr != nil {
		log.Fatalf("Error getting rows affected: %v", rowErr)
	}
	if rowsAffected == 0 {
		fmt.Println("No user created")
		return nil
	}
	fmt.Println("Rows affected", rowsAffected)

	return nil
}

func (userRepo *UserRepositoryImp) GetById(id string) (*models.User, error) {
	fmt.Println("getting user by id hit")
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE ID = ?"
	row := userRepo.db.QueryRow(query, id)
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			return nil, err
		}
		log.Fatalf("Error scanning database: %v", err)
	}
	fmt.Println("User fetched", user)
	return user, err
}
func (userRepo *UserRepositoryImp) GetByEmail(email string) (*models.User, error) {
	fmt.Println("getting user by email hit")
	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE email = ?"
	row := userRepo.db.QueryRow(query, email)
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			return nil, err
		}
		log.Fatalf("Error scanning database: %v", err)
	}
	fmt.Println("User fetched", user)
	return user, err
}

func (userRepo *UserRepositoryImp) GetAll() ([]models.User, error) {
	fmt.Println("getting user by id hit")
	query := "SELECT id, username, email, created_at, updated_at FROM users"
	rows, err := userRepo.db.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	users := []models.User{}
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No rows found")
				return nil, err
			}
			log.Fatalf("Error scanning database: %v", err)
		}
		users = append(users, *user)
	}
	fmt.Println("User fetched", users)
	return users, err
}

func (userRepo *UserRepositoryImp) DeleteById(id string) error {
	fmt.Println("delete user by id hit")
	query := "DELETE FROM users WHERE id = ?"
	result, err := userRepo.db.Exec(query, id)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	rowsAffected, rowsErr := result.RowsAffected()
	if rowsErr != nil {
		log.Fatalf("Error getting rows affected: %v", rowsErr)
	}
	if rowsAffected == 0 {
		fmt.Println("No user deleted")
		return nil
	}
	fmt.Println("Rows affected", rowsAffected)
	return nil

}
