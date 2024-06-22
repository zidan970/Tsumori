package service

import (
	"errors"
	"zidan/clean-arch/app/middlewares"
	"zidan/clean-arch/features/user"
)

type userService struct {
	userData user.UserDataInterface
}

// dependency injection
func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	// logic validation
	if input.Email == "" || input.Password == "" {
		return errors.New("[validation] email dan password harus diisi")
	}
	err := service.userData.Insert(input)
	return err
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	if email == "" || password == "" {
		return nil, "", errors.New("email dan password wajib diisi")
	}

	data, err = service.userData.Login(email, password)
	if err != nil {
		return nil, "", err
	}
	// log.Println("id user:", data.ID)
	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}
	return data, token, err
}

func (service *userService) GetSingle(id int) (user.Core, error) {
	// logic
	// memanggil func yg ada di data layer
	results, err := service.userData.GetSingle(id)
	return results, err
}

// Update implements user.UserServiceInterface.
func (service *userService) Update(id int, input user.Core) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	//validasi inputan
	// ...
	err := service.userData.Update(id, input)
	return err
}

// GetAll implements user.UserServiceInterface.
func (service *userService) GetAll() ([]user.Core, error) {
	// logic
	// memanggil func yg ada di data layer
	results, err := service.userData.SelectAll()
	return results, err
}

// Update implements user.UserServiceInterface.
func (service *userService) DeleteUser(input []user.Core, id int) error { //validasi
	if id == 0 {
		return errors.New("invalid id")
	}
	//validasi inputan
	// ...
	err := service.userData.DeleteUser(input, id)
	return err
}
