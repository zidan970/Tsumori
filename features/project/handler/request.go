package handler

import "zidan/clean-arch/features/project"

type ProjectRequest struct {
	Name   string `gorm:"not null" json:"name" form:"name"`
	UserID uint   `gorm:"not null" json:"user_id" form:"user_id"`
}

func RequestToCore(input ProjectRequest) project.Core {
	return project.Core{
		Name:   input.Name,
		UserID: input.UserID,
	}
}
