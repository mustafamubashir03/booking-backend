package db

import (
	"auth-go/models"
	"database/sql"
	"fmt"
	"strings"
)

type UserRoleRepository interface {
	GetUserRoles(userId int) ([]models.Role, error)
	AssignRoleToUser(userId int, roleId int) error
	UnAssignRoleFromUser(userId int, roleId int) error
	GetUserPermissions(userId int) ([]models.Permission, error)
	HasPermission(userId int, permissionName string) (bool, error)
	HasAllRoles(userId int, roleNames []string) (bool, error)
	HasAnyRole(userId int, roleNames []string) (bool, error)
}

type UserRoleRepositoryImp struct {
	db *sql.DB
}

func NewUserRoleRepository(db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImp{db: db}
}

func (userRoleRepo *UserRoleRepositoryImp) GetUserRoles(userId int) ([]models.Role, error) {
	query := `
		SELECT
			r.id,
			r.name,
			r.description,
			r.created_at,
			r.updated_at
		FROM user_roles ur
		INNER JOIN roles r ON r.id = ur.role_id
		WHERE ur.user_id = ?
	`

	rows, err := userRoleRepo.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	roles := []models.Role{}

	for rows.Next() {
		role := models.Role{}

		if err := rows.Scan(
			&role.Id,
			&role.Name,
			&role.Description,
			&role.CreatedAt,
			&role.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning database: %w", err)
		}

		roles = append(roles, role)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating roles: %w", err)
	}

	return roles, nil
}

func (userRoleRepo *UserRoleRepositoryImp) GetUserPermissions(userId int) ([]models.Permission, error) {
	query := "SELECT p.id, p.name, p.resource, p.action, p.created_at, p.updated_at FROM user_roles ur INNER JOIN role_permissions rp ON ur.role_id = rp.role_id INNER JOIN permissions p ON rp.permission_id = p.id WHERE ur.user_id = ?"
	rows, err := userRoleRepo.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()
	permissions := []models.Permission{}
	for rows.Next() {
		permission := &models.Permission{}
		if err := rows.Scan(
			&permission.Id,
			&permission.Name,
			&permission.Resource,
			&permission.Action,
			&permission.CreatedAt,
			&permission.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning database: %w", err)
		}
		permissions = append(permissions, *permission)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}
	fmt.Println("Permissions fetched", permissions)
	return permissions, nil
}

func (userRoleRepo *UserRoleRepositoryImp) HasPermission(userId int, permissionName string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM user_roles ur
			JOIN role_permissions rp ON ur.role_id = rp.role_id
			JOIN permissions p ON rp.permission_id = p.id
			WHERE ur.user_id = ?
			AND p.name = ?
		)
	`

	var exists bool

	err := userRoleRepo.db.QueryRow(
		query,
		userId,
		permissionName,
	).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("error checking permission: %w", err)
	}

	return exists, nil
}

func (userRoleRepo *UserRoleRepositoryImp) HasAllRoles(userId int, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return false, nil
	}

	placeholders := strings.TrimRight(strings.Repeat("?,", len(roleNames)), ",")

	query := fmt.Sprintf(`
		SELECT COUNT(DISTINCT r.name) = ?
		FROM user_roles ur
		JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ?
		AND r.name IN (%s)
	`, placeholders)

	args := make([]interface{}, 0, len(roleNames)+2)

	args = append(args, len(roleNames))
	args = append(args, userId)

	for _, roleName := range roleNames {
		args = append(args, roleName)
	}

	var hasAll bool

	err := userRoleRepo.db.QueryRow(query, args...).Scan(&hasAll)
	if err != nil {
		return false, fmt.Errorf("error executing query: %w", err)
	}

	fmt.Println("Has all roles:", hasAll)

	return hasAll, nil
}

func (userRoleRepo *UserRoleRepositoryImp) HasAnyRole(userId int, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return false, nil
	}

	placeholders := strings.TrimRight(strings.Repeat("?,", len(roleNames)), ",")

	query := fmt.Sprintf(`
		SELECT EXISTS (
			SELECT 1
			FROM user_roles ur
			JOIN roles r ON ur.role_id = r.id
			WHERE ur.user_id = ?
			AND r.name IN (%s)
		)
	`, placeholders)

	args := make([]interface{}, 0, len(roleNames)+1)
	args = append(args, userId)

	for _, roleName := range roleNames {
		args = append(args, roleName)
	}

	var exists bool

	err := userRoleRepo.db.QueryRow(query, args...).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error executing query: %w", err)
	}

	fmt.Println("Has any role:", exists)

	return exists, nil
}

func (userRoleRepo *UserRoleRepositoryImp) AssignRoleToUser(userId int, roleId int) error {
	query := "INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)"
	_, err := userRoleRepo.db.Exec(query, userId, roleId)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	return nil
}

func (userRoleRepo *UserRoleRepositoryImp) UnAssignRoleFromUser(userId int, roleId int) error {
	query := `
		DELETE FROM user_roles
		WHERE user_id = ?
		AND role_id = ?
	`

	result, err := userRoleRepo.db.Exec(query, userId, roleId)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking deleted role: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("role is not assigned to this user")
	}

	return nil
}
