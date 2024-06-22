package handler

import "zidan/clean-arch/features/user"

type UserResponse struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func CoreToResponse(data user.Core) UserResponse {
	return UserResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
	}
}

func CoreToResponseList(data []user.Core) []UserResponse {
	var results []UserResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
