package entity

import (
	db "github.com/conorsheppard/user-api-go/internal/db/sqlc"
	"time"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country" binding:"required,oneof=IE UK"` // todo: country codes
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Country   string `json:"country" binding:"required,oneof=IE UK"` // todo: country codes
}

type UserResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Country   string    `json:"country"`
	UpdatedAt time.Time `json:"password_changed_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Nickname:  user.Nickname,
		Country:   user.Country,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}
}

func GetAllUsersResponse(users []db.User) []UserResponse {
	userResponse := []UserResponse{}
	for _, user := range users {
		userResponse = append(userResponse, UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Nickname:  user.Nickname,
			Country:   user.Country,
			UpdatedAt: user.UpdatedAt,
			CreatedAt: user.CreatedAt,
		})
	}
	return userResponse
}
