package db

import (
	"auth-go/models"
	"database/sql"
	"fmt"
)

type RoleRepository interface {
	GetRoleById(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]models.Role, error)
	DeleteRoleById(id int) error
	UpdateRoleById(id int, name string, description string) (*models.Role, error)
	CreateRole(name string, description string) error
	GetRolesPermissions(roleId int) ([]models.Permission, error)
}

type RoleRepositoryImp struct {
	db *sql.DB
}

func NewRoleRepository(_db *sql.DB) RoleRepository {
	roleRepository := &RoleRepositoryImp{
		db: _db,
	}
	return roleRepository
}

func (roleRepo *RoleRepositoryImp) GetRoleById(id int) (*models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE id = ?"
	row := roleRepo.db.QueryRow(query, id)
	role := &models.Role{}
	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			return nil, err
		}
		return nil, fmt.Errorf("error scanning database: %w", err)
	}
	fmt.Println("Role fetched", role)
	return role, nil
}

func (roleRepo *RoleRepositoryImp) GetRoleByName(name string) (*models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE name = ?"
	row := roleRepo.db.QueryRow(query, name)
	role := &models.Role{}
	err := row.Scan(&role.Id, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			return nil, err
		}
		return nil, fmt.Errorf("error scanning database: %w", err)
	}
	fmt.Println("Role fetched", role)
	return role, nil
}

func (roleRepo *RoleRepositoryImp) GetAllRoles() ([]models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles"
	rows, err := roleRepo.db.Query(query)
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

func (roleRepo *RoleRepositoryImp) DeleteRoleById(id int) error {
	query := "DELETE FROM roles WHERE id = ?"
	result, err := roleRepo.db.Exec(query, id)
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

func (roleRepo *RoleRepositoryImp) UpdateRoleById(id int, name string, description string) (*models.Role, error) {
	query := "UPDATE roles SET name = ?, description = ? WHERE id = ?"
	result, err := roleRepo.db.Exec(query, name, description, id)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error getting rows affected: %w", err)
	}
	fmt.Println("Rows affected", rowsAffected)
	return &models.Role{Id: int64(id), Name: name, Description: description}, nil
}

func (roleRepo *RoleRepositoryImp) CreateRole(name string, description string) error {
	query := "INSERT INTO roles (name, description) VALUES (?, ?)"
	result, err := roleRepo.db.Exec(query, name, description)
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

func (roleRepo *RoleRepositoryImp) GetRolesPermissions(roleId int) ([]models.Permission, error) {
	query := "SELECT p.id, p.name, p.resource, p.action, p.created_at, p.updated_at FROM role_permissions rp INNER JOIN permissions p ON rp.permission_id = p.id WHERE rp.role_id = ?"
	rows, err := roleRepo.db.Query(query, roleId)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	permissions := []models.Permission{}
	for rows.Next() {
		permission := &models.Permission{}
		err := rows.Scan(&permission.Id, &permission.Name, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
			return nil, fmt.Errorf("error scanning database: %w", err)
		}
		permissions = append(permissions, *permission)
	}
	return permissions, nil
}
