package controllers

import (
	"auth-go/dto"
	"auth-go/services"
	"auth-go/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	userController := &UserController{
		userService: _userService,
	}
	return userController

}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(*dto.RegisterUserRequestDTO)
	if err := u.userService.CreateUser(payload.Name, payload.Email, payload.Password); err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Failed to create user.", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully.", nil)
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	if userId == "" {
		userId = r.Context().Value("userId").(string)
	}
	user, err := u.userService.GetUserById(userId)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON write failed.", err)
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully.", user)
}

func (u *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON write failed.", err)
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Users fetched successfully.", users)
}

func (u *UserController) DeleteById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	if userId == "" {
		userId = r.Context().Value("userId").(string)
	}
	if err := u.userService.DeleteUserById(userId); err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON write failed.", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User deleted successfully.", nil)
}
func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("payload").(*dto.LoginUserRequestDTO)
	jwtToken, err := u.userService.LoginUser(payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON write failed.", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully.", jwtToken)
}
