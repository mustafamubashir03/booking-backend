package db

import (
	"auth-go/models"
	"database/sql"
	"fmt"
)

type RolePermissionRepository interface {
	GetRolePermissions(roleId int) ([]models.Permission, error)
	GetRolePermissionById(id int) (models.RolePermission, error)
	AddPermissionToRole(roleId int, permissionId int) error
	RemovePermissionFromRole(roleId int, permissionId int) error
	GetAllPermissions() ([]models.Permission, error)
	GetPermissionByRole(roleId int) ([]models.Permission, error)
}

type RolePermissionRepositoryImp struct {
	db *sql.DB
}

func NewRolePermissionRepository(db *sql.DB) *RolePermissionRepositoryImp {
	return &RolePermissionRepositoryImp{db: db}
}

func (rolePermissionRepo *RolePermissionRepositoryImp) GetRolePermissions(roleId int) ([]models.Permission, error) {
	query := "SELECT p.id, p.name, p.resource, p.action, p.created_at, p.updated_at FROM role_permissions rp INNER JOIN permissions p ON rp.permission_id = p.id WHERE rp.role_id = ?"
	rows, err := rolePermissionRepo.db.Query(query, roleId)
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

func (rolePermissionRepo *RolePermissionRepositoryImp) GetRolesPermissionsById(roleId int) ([]models.RolePermission, error) {
	query := "SELECT p.id, p.name, p.resource, p.action, p.created_at, p.updated_at FROM role_permissions rp INNER JOIN permissions p ON rp.permission_id = p.id WHERE rp.role_id = ?"
	rows, err := rolePermissionRepo.db.Query(query, roleId)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	rolePermission := []models.RolePermission{}
	for rows.Next() {
		rolePermission := &models.RolePermission{}
		err := rows.Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId, &rolePermission.CreatedAt, &rolePermission.UpdatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("No rows found")
				return []models.RolePermission{}, err
			}
			return []models.RolePermission{}, fmt.Errorf("error scanning database: %w", err)
		}
		rolePermission = rolePermission
	}
	fmt.Println("Role permissions fetched", rolePermission)
	return rolePermission, nil
}

func (rolePermissionRepo *RolePermissionRepositoryImp) RemovePermissionFromRole(roleId int, permissionId int) error {
	query := "DELETE FROM role_permissions WHERE role_id = ? AND permission_id = ?"
	_, err := rolePermissionRepo.db.Exec(query, roleId, permissionId)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	fmt.Println("Permission unassigned")
	return nil
}

func (rolePermissionRepo *RolePermissionRepositoryImp) AddPermissionToRole(roleId int, permissionId int) error {
	query := "INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?)"
	_, err := rolePermissionRepo.db.Exec(query, roleId, permissionId)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	fmt.Println("Permission assigned")
	return nil
}

func (rolePermissionRepo *RolePermissionRepositoryImp) GetAllPermissions() ([]models.Permission, error) {
	query := "SELECT p.id, p.name, p.resource, p.action, p.created_at, p.updated_at FROM permissions p"
	rows, err := rolePermissionRepo.db.Query(query)
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

func (rolePermissionRepo *RolePermissionRepositoryImp) GetPermissionByRole(roleId int) ([]models.Permission, error) {
	query := "SELECT p.id, p.name, p.resource, p.action, p.created_at, p.updated_at FROM permissions p INNER JOIN role_permissions rp ON rp.permission_id = p.id WHERE rp.role_id = ?"
	rows, err := rolePermissionRepo.db.Query(query, roleId)
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
