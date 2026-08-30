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
	r.With(middlewares.AuthMiddleware, middlewares.ValidateRequest(&dto.CreateRoleRequestDTO{})).Post("/role", roleRouter.roleController.CreateRole)
	r.With(middlewares.AuthMiddleware).Get("/role", roleRouter.roleController.GetAllRoles)
	r.With(middlewares.AuthMiddleware).Get("/role/name/{rolename}", roleRouter.roleController.GetRoleByName)
	r.With(middlewares.AuthMiddleware).Get("/role/{id}", roleRouter.roleController.GetRoleById)
	r.With(middlewares.AuthMiddleware, middlewares.ValidateRequest(&dto.UpdateRoleRequestDTO{})).Put("/role/{id}", roleRouter.roleController.UpdateRoleById)
	r.With(middlewares.AuthMiddleware).Delete("/role/{id}", roleRouter.roleController.DeleteRoleById)
}
