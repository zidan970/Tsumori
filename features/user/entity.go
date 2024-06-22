package user

import "time"

type Core struct {
	ID          uint
	Name        string `gorm:"not null" json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	Address     string `gorm:"not null" json:"address" form:"address"`
	PhoneNumber string `gorm:"not null" json:"phone" form:"phone"`
	Password    string `gorm:"not null" json:"password" form:"password"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// interface untuk Data Layer
type UserDataInterface interface {
	Insert(input Core) error
	Login(email, password string) (data *Core, err error)
	GetSingle(productID_int int) (Core, error)
	Update(id int, input Core) error
	SelectAll() ([]Core, error)
	DeleteUser(input []Core, id int) error
}

// interface untuk Service Layer
type UserServiceInterface interface {
	Create(input Core) error
	Login(email, password string) (data *Core, token string, err error)
	GetSingle(productID_int int) (Core, error)
	Update(id int, input Core) error
	GetAll() ([]Core, error)
	DeleteUser(input []Core, id int) error
}
