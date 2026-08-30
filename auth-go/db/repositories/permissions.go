package db

import (
	"auth-go/models"
	"database/sql"
	"fmt"
)

type PermissionRepository interface {
	GetPermissionById(id int) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]models.Permission, error)
	DeletePermissionById(id int) error
	UpdatePermissionById(id int, name string, description string, resource string, action string) (*models.Permission, error)
	CreatePermission(name string, description string, resource string, action string) error
}

type PermissionRepositoryImp struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImp{db: db}
}

func (permRepo *PermissionRepositoryImp) GetPermissionById(id int) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE id = ?"
	row := permRepo.db.QueryRow(query, id)
	permission := &models.Permission{}
	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			return nil, err
		}
		return nil, fmt.Errorf("error scanning database: %w", err)
	}
	fmt.Println("Permission fetched", permission)
	return permission, nil
}

func (permRepo *PermissionRepositoryImp) GetPermissionByName(name string) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE name = ?"
	row := permRepo.db.QueryRow(query, name)
	permission := &models.Permission{}
	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			return nil, err
		}
		return nil, fmt.Errorf("error scanning database: %w", err)
	}
	fmt.Println("Permission fetched", permission)
	return permission, nil
}

func (permRepo *PermissionRepositoryImp) GetAllPermissions() ([]models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions"
	rows, err := permRepo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	permissions := []models.Permission{}
	for rows.Next() {
		permission := &models.Permission{}
		err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No rows found")
				return nil, err
			}
			return nil, fmt.Errorf("error scanning database: %w", err)
		}
		permissions = append(permissions, *permission)
	}
	fmt.Println("Permissions fetched", permissions)
	return permissions, nil
}

func (permRepo *PermissionRepositoryImp) DeletePermissionById(id int) error {
	query := "DELETE FROM permissions WHERE id = ?"
	result, err := permRepo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	fmt.Println("Rows affected", rowsAffected)
	return nil
}

func (permRepo *PermissionRepositoryImp) UpdatePermissionById(id int, name string, description string, resource string, action string) (*models.Permission, error) {
	query := "UPDATE permissions SET name = ?, description = ?, resource = ?, action = ? WHERE id = ?"
	result, err := permRepo.db.Exec(query, name, description, resource, action, id)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error getting rows affected: %w", err)
	}
	fmt.Println("Rows affected", rowsAffected)
	return &models.Permission{Id: int64(id), Name: name, Description: description, Resource: resource, Action: action}, nil
}

func (permRepo *PermissionRepositoryImp) CreatePermission(name string, description string, resource string, action string) error {
	query := "INSERT INTO permissions (name, description, resource, action) VALUES (?, ?, ?, ?)"
	result, err := permRepo.db.Exec(query, name, description, resource, action)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	fmt.Println("Rows affected", rowsAffected)
	return nil
}
