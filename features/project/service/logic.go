package service

import (
	"errors"
	"zidan/clean-arch/features/project"
)

type projectService struct {
	projectData project.ProjectDataInterface
}

// dependency injection
func New(repo project.ProjectDataInterface) project.ProjectServiceInterface {
	return &projectService{
		projectData: repo,
	}
}

func (service *projectService) Create(input project.Core) error {
	// logic validation
	if input.Name == "" {
		return errors.New("[validation] nama project harus diisi")
	}
	err := service.projectData.Insert(input)
	return err
}

func (service *projectService) GetAll() ([]project.Core, error) {
	// logic
	// memanggil func yg ada di data layer
	results, err := service.projectData.SelectAll()
	return results, err
}

func (service *projectService) GetProjectsByUserID(userID int) ([]project.Core, error) {
	results, err := service.projectData.SelectByUserID(userID)
	return results, err
}

// GetDetail implements project.ProjectServiceInterface.
func (service *projectService) GetDetail(id int) (project.Core, error) {
	results, err := service.projectData.GetDetail(id)
	return results, err
}

func (service *projectService) IsUserAuthorizedToUpdate(userID, projectID int) (bool, error) {
	// Logika validasi hak akses pengguna
	result, err := service.projectData.SelectByProjectID(projectID)
	if err != nil {
		return false, err
	}

	return result.UserID == uint(userID), nil
}

func (service *projectService) Update(id int, input project.Core) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	//validasi inputan
	// ...
	err := service.projectData.Update(id, input)
	return err
}

func (service *projectService) IsUserAuthorizedToDelete(userID, projectID int) (bool, error) {
	// Logika validasi hak akses pengguna
	result, err := service.projectData.SelectByProjectID(projectID)
	if err != nil {
		return false, err
	}

	return result.UserID == uint(userID), nil
}

// Delete implements project.ProjectServiceInterface.
func (service *projectService) Delete(input []project.Core, id int) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	//validasi inputan
	// ...

	err := service.projectData.Delete(input, id)
	return err
}
