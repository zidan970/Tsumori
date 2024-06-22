package handler

import (
	"zidan/clean-arch/features/project"
)

type ProjectResponse struct {
	ID         uint   `json:"id" form:"id"`
	Name       string `gorm:"not null" json:"name" form:"name"`
	UserCoreID uint   `gorm:"not null" json:"user_id" form:"user_id"`
}

func CoreToResponse(data project.Core) ProjectResponse {
	return ProjectResponse{
		ID:         uint(data.ID),
		Name:       data.Name,
		UserCoreID: uint(data.UserID),
	}
}

func CoreToResponseList(data []project.Core) []ProjectResponse {
	var results []ProjectResponse
	for _, v := range data {
		results = append(results, CoreToResponse(v))
	}
	return results
}
