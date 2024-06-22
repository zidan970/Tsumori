package data

import (
	"errors"
	"zidan/clean-arch/features/task"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) task.TaskDataInterface {
	return &taskQuery{
		db: db,
	}
}

// InsertTask implements task.TaskDataInterface.
func (repo *taskQuery) InsertTask(input task.Core) error {
	// simpan ke DB
	newTaskGorm := Task{
		Name:      input.Name,
		ProjectID: input.ProjectID,
		Status:    input.Status,
	}

	tx := repo.db.Create(&newTaskGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetTaskByID implements task.TaskDataInterface.
func (repo *taskQuery) GetTaskByID(id int) (task.Core, error) {
	var singleTaskGorm Task
	tx := repo.db.First(&singleTaskGorm, id)
	if tx.Error != nil {
		return task.Core{}, tx.Error
	}

	singleTaskCore := ModelToCore(singleTaskGorm)

	return singleTaskCore, nil
}

// Delete implements task.TaskDataInterface.
func (repo *taskQuery) Delete(input task.Core, id int) error {
	taskGorm := CoreToModel(input)

	txDel := repo.db.Delete(&taskGorm, id)
	if txDel.Error != nil {
		return txDel.Error
	}

	if txDel.RowsAffected == 0 {
		return errors.New("user's not found")
	}

	return nil
}

// Update implements task.TaskDataInterface.
func (repo *taskQuery) Update(id int, input task.Core) error {
	dataGorm := CoreToModel(input)
	tx := repo.db.Model(&Task{}).Where("id = ?", id).Updates(dataGorm)
	if tx.Error != nil {
		// fmt.Println("err:", tx.Error)
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}
