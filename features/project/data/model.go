package data

import (
	"zidan/clean-arch/features/project"
	"zidan/clean-arch/features/user/data"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name   string    `gorm:"not null" json:"name" form:"name"`
	UserID uint      `gorm:"not null" json:"user_id" form:"user_id"`
	User   data.User `gorm:"foreignKey:UserID" json:"user" form:"user"`
}

func CoreToModel(input project.Core) Project {
	return Project{
		Name:   input.Name,
		UserID: input.UserID,
	}
}

func CoretoModelGorm(data []project.Core) []Project {
	var projectsDataGorm []Project
	for _, value := range data {
		var projectGorm = Project{
			Name:   value.Name,
			UserID: value.UserID,
		}
		projectsDataGorm = append(projectsDataGorm, projectGorm)
	}

	return projectsDataGorm
}

func ModelToCore(input Project) project.Core {
	return project.Core{
		Name:   input.Name,
		UserID: input.UserID,
	}
}

func ModelToCoreGorm(data []Project) []project.Core {
	var allProjectCore []project.Core
	for _, value := range data {
		var projectCore = project.Core{
			ID:        value.ID,
			Name:      value.Name,
			UserID:    value.UserID,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		allProjectCore = append(allProjectCore, projectCore)
	}

	return allProjectCore
}
