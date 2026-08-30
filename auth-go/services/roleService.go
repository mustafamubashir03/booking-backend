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
}

type RoleServiceImp struct {
	roleRepository db.RoleRepository
}

func NewRoleService(_roleRepository db.RoleRepository) RoleService {
	return &RoleServiceImp{roleRepository: _roleRepository}
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
