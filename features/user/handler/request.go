package handler

import "zidan/clean-arch/features/user"

type UserRequest struct {
	Name        string `gorm:"not null" json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	Address     string `gorm:"not null" json:"address" form:"address"`
	PhoneNumber string `gorm:"not null" json:"phone" form:"phone"`
	Password    string `gorm:"not null" json:"password" form:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(input UserRequest) user.Core {
	return user.Core{
		Name:        input.Name,
		Email:       input.Email,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}
}
