package dto

type UpdateRoleRequestDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateRoleRequestDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
