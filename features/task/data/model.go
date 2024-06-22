package data

import (
	"zidan/clean-arch/features/project/data"
	"zidan/clean-arch/features/task"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name      string       `gorm:"not null" json:"name" form:"name"`
	ProjectID uint         `json:"project_id" form:"project_id"`
	Status    string       `gorm:"not null" json:"status" form:"status"`
	Project   data.Project `gorm:"foreignKey:ProjectID" json:"project" form:"project"`
}

func CoreToModel(input task.Core) Task {
	return Task{
		Name:      input.Name,
		ProjectID: input.ProjectID,
		Status:    input.Status,
	}
}

func ModelToCore(data Task) task.Core {
	allProjectCore := task.Core{
		Name:      data.Name,
		ProjectID: data.ProjectID,
		Status:    data.Status,
	}

	return allProjectCore
}

func ModelToCoreGorm(data []Task) []task.Core {
	var allProjectCore []task.Core
	for _, value := range data {
		var projectCore = task.Core{
			Name:      value.Name,
			ProjectID: value.ProjectID,
			Status:    value.Status,
		}
		allProjectCore = append(allProjectCore, projectCore)
	}

	return allProjectCore
}
