package dto

type RegisterUserRequestDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginUserRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type DeleteUserRequestDTO struct {
	Id string `param:"id" validate:"required"`
}

type GetUserByIdRequestDTO struct {
	Id string `param:"id" validate:"required"`
}
