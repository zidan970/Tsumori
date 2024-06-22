package data

import (
	"errors"
	"zidan/clean-arch/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	// proses mapping dari struct entities core ke model gorm
	userInputGorm := CoreToModel(input)
	// simpan ke DB
	tx := repo.db.Create(&userInputGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (data *user.Core, err error) {
	var userGorm User
	tx := repo.db.Where("email = ? and password = ?", email, password).First(&userGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}
	result := ModelToCore(userGorm)
	return &result, nil
}

func (repo *userQuery) GetSingle(id int) (user.Core, error) {
	var singleProductGorm User
	tx := repo.db.First(&singleProductGorm, id)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	singleProductCore := ModelToCore(singleProductGorm)

	return singleProductCore, nil
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(id int, input user.Core) error {
	dataGorm := CoreToModel(input)
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(dataGorm)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// SelectAll implements user.UserDataInterface.
func (repo *userQuery) SelectAll() ([]user.Core, error) {
	var usersDataGorm []User
	tx := repo.db.Find(&usersDataGorm) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	usersDataCore := ModelToCoreGorm(usersDataGorm)

	return usersDataCore, nil
}

func (repo *userQuery) DeleteUser(input []user.Core, id int) error {
	// proses mapping dari struct gorm model ke struct core
	usersDataGorm := CoretoModelGorm(input)

	result := repo.db.Delete(&usersDataGorm, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user's not found")
	}

	return nil
}
