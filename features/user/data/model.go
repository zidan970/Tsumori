package data

import (
	"zidan/clean-arch/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name" form:"name"`
	Email       string `gorm:"unique" json:"email" form:"email"`
	Address     string `gorm:"not null" json:"address" form:"address"`
	PhoneNumber string `gorm:"not null" json:"phone" form:"phone"`
	Password    string `gorm:"not null" json:"password" form:"password"`
}

func CoreToModel(input user.Core) User {
	return User{
		Name:        input.Name,
		Email:       input.Email,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}
}

func CoretoModelGorm(data []user.Core) []User {
	var usersDataGorm []User
	for _, value := range data {
		var userGorm = User{
			Name:     value.Name,
			Email:    value.Email,
			Password: value.Password,
		}
		usersDataGorm = append(usersDataGorm, userGorm)
	}

	return usersDataGorm
}

func ModelToCore(u User) user.Core {
	return user.Core{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Address:     u.Address,
		PhoneNumber: u.PhoneNumber,
		Password:    u.Password,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func ModelToCoreGorm(u []User) []user.Core {
	var usersDataCore []user.Core
	for _, value := range u {
		var userCore = user.Core{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Address:     value.Address,
			PhoneNumber: value.PhoneNumber,
			Password:    value.Password,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		}
		usersDataCore = append(usersDataCore, userCore)
	}

	return usersDataCore
}
