package task

import (
	"time"
)

type Core struct {
	ID        uint   `json:"id"`
	Name      string `gorm:"not null" json:"name" form:"name"`
	ProjectID uint   `json:"project_id" form:"project_id"`
	Status    string `gorm:"not null" json:"status" form:"status"`
	//Project   data.Project `gorm:"foreignKey:ProjectID" json:"project" form:"project"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// interface untuk Data Layer
type TaskDataInterface interface {
	InsertTask(input Core) error
	Update(id int, input Core) error
	GetTaskByID(id int) (Core, error)
	Delete(Core, int) error
}

// interface untuk Service Layer
type TaskServiceInterface interface {
	CreateTask(input Core) error
	Update(id int, input Core) error
	GetTaskByID(id int) (Core, error)
	Delete(Core, int) error
}
