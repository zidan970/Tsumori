package data

import (
	"errors"
	"zidan/clean-arch/features/project"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) project.ProjectDataInterface {
	return &projectQuery{
		db: db,
	}
}

func (repo *projectQuery) Insert(input project.Core) error {
	// simpan ke DB
	newProjectGorm := CoreToModel(input)

	tx := repo.db.Create(&newProjectGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (repo *projectQuery) SelectByUserID(userID int) ([]project.Core, error) {
	var projectsDataGorm []Project
	tx := repo.db.Where("user_id = ?", userID).Find(&projectsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Mapping proyek ke model Core
	allProjectCore := ModelToCoreGorm(projectsDataGorm)

	return allProjectCore, nil
}

// GetDetail implements project.ProjectDataInterface.
func (repo *projectQuery) GetDetail(id int) (project.Core, error) {
	// Query untuk mendapatkan detail proyek berdasarkan ID
	var projectCup Project
	if err := repo.db.Preload("Task").First(&projectCup, id).Error; err != nil {
		if err != nil {
			return project.Core{}, err
		}
	}

	projectCupCore := ModelToCore(projectCup)

	return projectCupCore, nil
}

func (repo *projectQuery) SelectByProjectID(projectID int) (project.Core, error) {
	var projectGorm Project
	tx := repo.db.Where("ID = ?", projectID).First(&projectGorm)
	if tx.Error != nil {
		return project.Core{}, tx.Error
	}

	projectCore := ModelToCore(projectGorm)

	return projectCore, nil
}

func (repo *projectQuery) Update(projectID int, updatedProject project.Core) error {
	newUpdateGorm := CoreToModel(updatedProject)

	txUpdates := repo.db.Model(&Project{}).Where("id = ?", projectID).Updates(newUpdateGorm)
	if txUpdates.Error != nil {
		return txUpdates.Error
	}

	return nil
}

func (repo *projectQuery) SelectAll() ([]project.Core, error) {
	var projectsDataGorm []Project
	tx := repo.db.Find(&projectsDataGorm) // select * from users;
	if tx.Error != nil {
		return nil, tx.Error
	}

	//mapping
	allProjectCore := ModelToCoreGorm(projectsDataGorm)

	return allProjectCore, nil
}

func (repo *projectQuery) Delete(input []project.Core, id int) error {
	allProjectGorm := CoretoModelGorm(input)

	txDel := repo.db.Delete(&allProjectGorm, id)
	if txDel.Error != nil {
		return txDel.Error
	}

	if txDel.RowsAffected == 0 {
		return errors.New("user's not found")
	}

	return nil
}
