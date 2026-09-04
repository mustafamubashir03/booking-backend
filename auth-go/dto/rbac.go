package dto

type UpdateRoleRequestDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateRoleRequestDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type AssignPermissionToRoleRequestDTO struct {
	PermissionId string `param:"permissionId" validate:"required"`
}

type RemovePermissionFromRoleRequestDTO struct {
	PermissionId string `param:"permissionId" validate:"required"`
}

type GetRolesPermissionsRequestDTO struct {
	RoleId string `param:"id" validate:"required"`
}

type RevokeAllPermissionsFromRoleRequestDTO struct {
	RoleId string `param:"roleId" validate:"required"`
}

type RevokeAllRolesPermissionsRequestDTO struct {
	RoleId string `param:"roleId" validate:"required"`
}

type AssignRoleToUserRequestDTO struct {
	RoleId string `param:"roleId" validate:"required"`
	UserId string `param:"userId" validate:"required"`
}

type UnAssignRoleFromUserRequestDTO struct {
	RoleId string `param:"roleId" validate:"required"`
	UserId string `param:"userId" validate:"required"`
}

type GetUserRolesRequestDTO struct {
	UserId string `param:"userId" validate:"required"`
}

type GetUserPermissionsRequestDTO struct {
	UserId string `param:"userId" validate:"required"`
}

type HasPermissionRequestDTO struct {
	UserId         string `param:"userId" validate:"required"`
	PermissionName string `param:"permissionName" validate:"required"`
}

type HasAllRolesRequestDTO struct {
	UserId    string `param:"userId" validate:"required"`
	RoleNames string `param:"roleNames" validate:"required"`
}
