package handler

import "zidan/clean-arch/features/task"

type TaskRequest struct {
	Name      string `gorm:"not null" json:"name" form:"name"`
	ProjectID uint   `json:"project_id" form:"project_id"`
	Status    string `gorm:"not null" json:"status" form:"status"`
}

func RequestToCore(input TaskRequest) task.Core {
	return task.Core{
		Name:      input.Name,
		ProjectID: input.ProjectID,
		Status:    input.Status,
	}
}
