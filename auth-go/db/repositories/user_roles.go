package db

import (
	"auth-go/models"
	"database/sql"
	"fmt"
)

type UserRoleRepository interface {
	GetUserRoles(userId int) ([]models.Role, error)
	AssignRole(userId int, roleId int) error
	UnAssignRole(userId int, roleId int) error
	GetUserPermissions(userId int) ([]models.Permission, error)
	HasPermission(userId int, permissionName string) (bool, error)
	HasRole(userId int, roleName string) (bool, error)
}

type UserRoleRepositoryImp struct {
	db *sql.DB
}

func NewUserRoleRepositoryImp(db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImp{db: db}
}

func (userRoleRepo *UserRoleRepositoryImp) GetUserRoles(userId int) ([]models.Role, error) {
	query := "SELECT r.id, ur.user_id, ur.role_id, ur.created_at, ur.updated_at FROM user_roles ur INNER JOIN roles r ON r.id = ur.role_id WHERE ur.user_id = ?"
	rows, err := userRoleRepo.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	roles := []models.Role{}
	for rows.Next() {
		role := &models.Role{}
		err := rows.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No rows found")
				return nil, err
			}
			return nil, fmt.Errorf("error scanning database: %w", err)
		}
		roles = append(roles, *role)
	}
	fmt.Println("Roles fetched", roles)
	return roles, nil
}

func (userRoleRepo *UserRoleRepositoryImp) GetUserPermissions(userId int) ([]models.Permission, error) {
	query := "SELECT p.id, p.name, p.resource, p.action, p.created_at, p.updated_at FROM user_roles ur INNER JOIN role_permissions rp ON ur.role_id = rp.role_id INNER JOIN permissions p ON rp.permission_id = p.id WHERE ur.user_id = ?"
	rows, err := userRoleRepo.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	permissions := []models.Permission{}
	for rows.Next() {
		permission := &models.Permission{}
		err := rows.Scan(&permission.Id, &permission.Name, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
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

func (userRoleRepo *UserRoleRepositoryImp) HasPermission(userId int, permissionName string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM user_roles ur JOIN role_permissions rp ON ur.role_id = rp.role_id JOIN permissions p ON rp.permission_id = p.id WHERE ur.user_id = ? AND p.name = ?)"
	rows, err := userRoleRepo.db.Query(query, userId, permissionName)
	if err != nil {
		return false, fmt.Errorf("error executing query: %w", err)
	}
	var exists bool
	for rows.Next() {
		err := rows.Scan(&exists)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No rows found")
				return false, err
			}
			return false, fmt.Errorf("error scanning database: %w", err)
		}
	}
	fmt.Println("Has permission", exists)
	return exists, nil
}

func (userRoleRepo *UserRoleRepositoryImp) HasRole(userId int, roleName string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM user_roles ur JOIN roles r ON ur.role_id = r.id WHERE ur.user_id = ? AND r.name = ?)"
	rows, err := userRoleRepo.db.Query(query, userId, roleName)
	if err != nil {
		return false, fmt.Errorf("error executing query: %w", err)
	}
	var exists bool
	for rows.Next() {
		err := rows.Scan(&exists)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No rows found")
				return false, err
			}
			return false, fmt.Errorf("error scanning database: %w", err)
		}
	}
	fmt.Println("Has role", exists)
	return exists, nil
}

func (userRoleRepo *UserRoleRepositoryImp) AssignRole(userId int, roleId int) error {
	query := "INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)"
	_, err := userRoleRepo.db.Exec(query, userId, roleId)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	return nil
}

func (userRoleRepo *UserRoleRepositoryImp) UnAssignRole(userId int, roleId int) error {
	query := "DELETE FROM user_roles WHERE user_id = ? AND role_id = ?"
	_, err := userRoleRepo.db.Exec(query, userId, roleId)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	fmt.Println("Role unassigned")
	return nil
}
