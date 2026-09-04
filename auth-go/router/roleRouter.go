package router

import (
	"auth-go/controllers"
	"auth-go/dto"
	"auth-go/middlewares"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	roleController *controllers.RoleController
}

func NewRoleRouter(_roleController *controllers.RoleController) Router {
	roleRouter := &RoleRouter{
		roleController: _roleController,
	}
	return roleRouter
}

func (roleRouter *RoleRouter) Register(r chi.Router) {
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin"), middlewares.ValidateRequest(&dto.CreateRoleRequestDTO{})).Post("/role", roleRouter.roleController.CreateRole)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin")).Get("/role", roleRouter.roleController.GetAllRoles)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin", "user")).Get("/role/name/{rolename}", roleRouter.roleController.GetRoleByName)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin", "user")).Get("/role/{id}", roleRouter.roleController.GetRoleById)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin"), middlewares.ValidateRequest(&dto.UpdateRoleRequestDTO{})).Put("/role/{id}", roleRouter.roleController.UpdateRoleById)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin")).Delete("/role/{id}", roleRouter.roleController.DeleteRoleById)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin", "user"), middlewares.ValidateRequest(&dto.GetRolesPermissionsRequestDTO{})).Get("/role/{id}/permissions", roleRouter.roleController.GetRolesPermissions)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin"), middlewares.ValidateParams(&dto.AssignPermissionToRoleRequestDTO{})).Post("/role/{roleId}/permission/{permissionId}", roleRouter.roleController.AssignPermissionToRole)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin"), middlewares.ValidateParams(&dto.RemovePermissionFromRoleRequestDTO{})).Delete("/role/{roleId}/permission/{permissionId}", roleRouter.roleController.RemovePermissionFromRole)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin"), middlewares.ValidateParams(&dto.RevokeAllPermissionsFromRoleRequestDTO{})).Delete("/role/{roleId}/permissions", roleRouter.roleController.RevokeAllPermissionsFromRole)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin")).Delete("/role/permissions", roleRouter.roleController.RevokeAllRolesPermissions)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin"), middlewares.ValidateParams(&dto.AssignRoleToUserRequestDTO{})).Post("/user/{userId}/role/{roleId}", roleRouter.roleController.AssignRoleToUser)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin"), middlewares.ValidateParams(&dto.UnAssignRoleFromUserRequestDTO{})).Delete("/user/{userId}/role/{roleId}", roleRouter.roleController.UnAssignRoleFromUser)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin", "user"), middlewares.ValidateParams(&dto.GetUserRolesRequestDTO{})).Get("/user/{userId}/roles", roleRouter.roleController.GetUserRoles)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin", "user"), middlewares.ValidateParams(&dto.GetUserPermissionsRequestDTO{})).Get("/user/{userId}/permissions", roleRouter.roleController.GetUserPermissions)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin", "user"), middlewares.ValidateParams(&dto.HasPermissionRequestDTO{})).Get("/user/{userId}/permission/{permissionName}", roleRouter.roleController.HasPermission)
	r.With(middlewares.AuthMiddleware, middlewares.RequireAllRolesMiddleware("admin", "user"), middlewares.ValidateParams(&dto.HasAllRolesRequestDTO{})).Get("/user/{userId}/roles/{roleNames}", roleRouter.roleController.HasAllRoles)
}
