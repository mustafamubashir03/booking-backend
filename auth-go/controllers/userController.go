package controllers

import (
	"auth-go/dto"
	"auth-go/services"
	"auth-go/utils"
	"fmt"
	"net/http"
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
	fmt.Println("user controller hit")
	u.userService.CreateUser("Mustafa", "mustafamubashir87@gmail.com", "12345")
	w.Write([]byte("created user successful"))
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user controller hit")
	u.userService.GetUserById()
	w.Write([]byte("got user successful"))
}

func (u *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user controller hit")
	u.userService.GetAllUsers()
	w.Write([]byte("got all users successful"))
}

func (u *UserController) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user controller hit")
	u.userService.DeleteUserById()
	w.Write([]byte("deleted user successful"))
}
func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginUserRequestDTO
	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		err := utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON parse error.", jsonErr)
		if err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON write failed.", err)
		}
		return
	}

	fmt.Printf("[PARSED PAYLOAD] Email: %q | Password: %q\n", payload.Email, payload.Password)

	if validationErr := utils.Validate.Struct(payload); validationErr != nil {
		err := utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed.", validationErr)
		if err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON write failed.", err)
		}
		return
	}
	fmt.Println("user controller hit")
	jwtToken, err := u.userService.LoginUser(&payload)
	if err != nil {
		fmt.Println(err)
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON write failed.", err)
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully.", jwtToken)
}
