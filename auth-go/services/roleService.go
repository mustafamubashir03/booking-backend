package services

import (
	db "auth-go/db/repositories"
	"auth-go/models"
)

type RoleService interface {
	GetRoleById(id int) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]models.Role, error)
	DeleteRoleById(id int) error
	UpdateRoleById(id int, name string, description string) (*models.Role, error)
	CreateRole(name string, description string) error
	GetRolesPermissions(roleId int) ([]models.Permission, error)
	AssignPermissionToRole(roleId int, permissionId int) error
	RemovePermissionFromRole(roleId int, permissionId int) error
	RevokeAllPermissionsFromRole(roleId int) error
	RevokeAllRolesPermissions() error
	AssignRoleToUser(userId int, roleId int) error
	UnAssignRoleFromUser(userId int, roleId int) error
	GetUserRoles(userId int) ([]models.Role, error)
	GetUserPermissions(userId int) ([]models.Permission, error)
	HasPermission(userId int, permissionName string) (bool, error)
	HasAllRoles(userId int, roleNames []string) (bool, error)
}

type RoleServiceImp struct {
	roleRepository       db.RoleRepository
	permissionRepository db.PermissionRepository
	userRoleRepository   db.UserRoleRepository
}

func NewRoleService(_roleRepository db.RoleRepository, _permissionRepository db.PermissionRepository, _userRoleRepository db.UserRoleRepository) RoleService {
	return &RoleServiceImp{roleRepository: _roleRepository, permissionRepository: _permissionRepository, userRoleRepository: _userRoleRepository}
}

func (roleService *RoleServiceImp) GetRoleById(id int) (*models.Role, error) {
	return roleService.roleRepository.GetRoleById(id)
}

func (roleService *RoleServiceImp) GetRoleByName(name string) (*models.Role, error) {
	return roleService.roleRepository.GetRoleByName(name)
}

func (roleService *RoleServiceImp) GetAllRoles() ([]models.Role, error) {
	return roleService.roleRepository.GetAllRoles()
}

func (roleService *RoleServiceImp) DeleteRoleById(id int) error {
	return roleService.roleRepository.DeleteRoleById(id)
}

func (roleService *RoleServiceImp) UpdateRoleById(id int, name string, description string) (*models.Role, error) {
	return roleService.roleRepository.UpdateRoleById(id, name, description)
}

func (roleService *RoleServiceImp) CreateRole(name string, description string) error {
	return roleService.roleRepository.CreateRole(name, description)
}

func (roleService *RoleServiceImp) GetRolesPermissions(roleId int) ([]models.Permission, error) {
	return roleService.roleRepository.GetRolesPermissions(roleId)
}

func (roleService *RoleServiceImp) AssignPermissionToRole(roleId int, permissionId int) error {
	return roleService.roleRepository.AssignPermissionToRole(roleId, permissionId)
}

func (roleService *RoleServiceImp) RemovePermissionFromRole(roleId int, permissionId int) error {
	return roleService.roleRepository.RemovePermissionFromRole(roleId, permissionId)
}

func (roleService *RoleServiceImp) RevokeAllPermissionsFromRole(roleId int) error {
	return roleService.roleRepository.RevokeAllPermissionsFromRole(roleId)
}

func (roleService *RoleServiceImp) RevokeAllRolesPermissions() error {
	return roleService.roleRepository.RevokeAllRolesPermissions()
}

func (roleService *RoleServiceImp) AssignRoleToUser(userId int, roleId int) error {
	return roleService.userRoleRepository.AssignRoleToUser(userId, roleId)
}

func (roleService *RoleServiceImp) UnAssignRoleFromUser(userId int, roleId int) error {
	return roleService.userRoleRepository.UnAssignRoleFromUser(userId, roleId)
}

func (roleService *RoleServiceImp) GetUserRoles(userId int) ([]models.Role, error) {
	return roleService.userRoleRepository.GetUserRoles(userId)
}

func (roleService *RoleServiceImp) GetUserPermissions(userId int) ([]models.Permission, error) {
	return roleService.userRoleRepository.GetUserPermissions(userId)
}

func (roleService *RoleServiceImp) HasPermission(userId int, permissionName string) (bool, error) {
	return roleService.userRoleRepository.HasPermission(userId, permissionName)
}

func (roleService *RoleServiceImp) HasAllRoles(userId int, roleNames []string) (bool, error) {
	return roleService.userRoleRepository.HasAllRoles(userId, roleNames)
}
