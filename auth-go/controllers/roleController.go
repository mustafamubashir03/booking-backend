package controllers

import (
	"auth-go/dto"
	"auth-go/services"
	"auth-go/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{roleService: roleService}
}

func (r *RoleController) CreateRole(w http.ResponseWriter, req *http.Request) {
	payload := req.Context().Value("payload").(dto.CreateRoleRequestDTO)
	err := r.roleService.CreateRole(payload.Name, payload.Description)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Failed to create role.", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role created successfully.", nil)

}

func (r *RoleController) GetAllRoles(w http.ResponseWriter, req *http.Request) {
	roles, err := r.roleService.GetAllRoles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Roles fetched successfully.", roles)
}

func (r *RoleController) GetRoleById(w http.ResponseWriter, req *http.Request) {
	roleId := chi.URLParam(req, "id")
	id, _ := strconv.Atoi(roleId)
	role, err := r.roleService.GetRoleById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role fetched successfully.", role)
}

// func (r *RoleController) GetRolesPermissions(w http.ResponseWriter, req *http.Request) {
// 	roleId := req.URL.Query().Get("roleId")
// 	id, _ := strconv.Atoi(roleId)
// 	rolePermissions, err := r.roleService.GetRolePermissions(id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role permissions fetched successfully.", rolePermissions)
// }

// func (r *RoleController) AddPermissionToRole(w http.ResponseWriter, req *http.Request) {
// 	permissionId := req.URL.Query().Get("permissionId")
// 	roleId := req.URL.Query().Get("roleId")
// 	id, _ := strconv.Atoi(permissionId)
// 	rid, _ := strconv.Atoi(roleId)
// 	err := r.roleService.AddPermissionToRole(rid, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Permission assigned to role successfully.", nil)
// }

// func (r *RoleController) RemovePermissionFromRole(w http.ResponseWriter, req *http.Request) {
// 	permissionId := req.URL.Query().Get("permissionId")
// 	roleId := req.URL.Query().Get("roleId")
// 	id, _ := strconv.Atoi(permissionId)
// 	rid, _ := strconv.Atoi(roleId)
// 	err := r.roleService.RemovePermissionFromRole(rid, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Permission removed from role successfully.", nil)
// }

func (r *RoleController) GetRoleByName(w http.ResponseWriter, req *http.Request) {
	roleName := chi.URLParam(req, "rolename")
	role, err := r.roleService.GetRoleByName(roleName)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Failed to get role by name.", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role fetched successfully.", role)
}

func (r *RoleController) DeleteRoleById(w http.ResponseWriter, req *http.Request) {
	roleId := chi.URLParam(req, "id")
	id, _ := strconv.Atoi(roleId)
	err := r.roleService.DeleteRoleById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role deleted successfully.", nil)
}
func (r *RoleController) UpdateRoleById(w http.ResponseWriter, req *http.Request) {
	roleId := chi.URLParam(req, "id")
	id, _ := strconv.Atoi(roleId)
	payload := req.Context().Value("payload").(dto.UpdateRoleRequestDTO)
	role, err := r.roleService.UpdateRoleById(id, payload.Name, payload.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role updated successfully.", role)
}

// func (r *RoleController) GetRolePermissions(w http.ResponseWriter, req *http.Request) {
// 	roleId := req.URL.Query().Get("roleId")
// 	id, _ := strconv.Atoi(roleId)
// 	rolePermissions, err := r.roleService.GetRolePermissions(id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role permissions fetched successfully.", rolePermissions)
// }
// func (r *RoleController) GetAllPermissions(w http.ResponseWriter, req *http.Request) {
// 	permissions, err := r.roleService.GetAllPermissions()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Permissions fetched successfully.", permissions)
// }
