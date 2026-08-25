package controllers

import (
	"auth-go/dto"
	"auth-go/middlewares"
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
	payload := r.Context().Value(
		middlewares.RequestDTOKey,
	).(*dto.RegisterUserRequestDTO)
	u.userService.CreateUser(payload.Name, payload.Email, payload.Password)
	w.Write([]byte("created user successful"))
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
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
	if err := u.userService.DeleteUserById(userId); err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON write failed.", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User deleted successfully.", nil)
}
func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value(middlewares.RequestDTOKey).(*dto.LoginUserRequestDTO)
	jwtToken, err := u.userService.LoginUser(payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON write failed.", err)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully.", jwtToken)
}
