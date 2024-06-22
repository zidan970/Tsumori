package project

import (
	"time"
	"zidan/clean-arch/features/user/data"
)

type Core struct {
	ID        uint      `json:"id"`
	Name      string    `gorm:"not null" json:"name" form:"name"`
	UserID    uint      `gorm:"not null" json:"user_id" form:"user_id"`
	User      data.User `gorm:"foreignKey:UserID" json:"user" form:"user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// interface untuk Data Layer
type ProjectDataInterface interface {
	Insert(input Core) error
	SelectByUserID(id int) ([]Core, error)
	GetDetail(int) (Core, error)
	SelectByProjectID(int) (Core, error)
	Update(id int, input Core) error
	SelectAll() ([]Core, error)
	Delete(input []Core, id int) error
}

// interface untuk Service Layer
type ProjectServiceInterface interface {
	Create(input Core) error
	GetProjectsByUserID(id int) ([]Core, error)
	GetDetail(int) (Core, error)
	IsUserAuthorizedToUpdate(int, int) (bool, error)
	Update(id int, input Core) error
	GetAll() ([]Core, error)
	IsUserAuthorizedToDelete(int, int) (bool, error)
	Delete(input []Core, id int) error
}
