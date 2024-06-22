package service

import (
	"errors"
	"zidan/clean-arch/features/task"
)

type taskService struct {
	taskData task.TaskDataInterface
}

// dependency injection
func New(repo task.TaskDataInterface) task.TaskServiceInterface {
	return &taskService{
		taskData: repo,
	}
}

// CreateTask implements task.TaskServiceInterface.
func (service *taskService) CreateTask(input task.Core) error {
	// logic validation
	if input.Name == "" {
		return errors.New("[validation] nama task harus diisi")
	}
	err := service.taskData.InsertTask(input)
	return err
}

// GetTaskByID implements task.TaskServiceInterface.
func (service *taskService) GetTaskByID(id int) (task.Core, error) {
	if id == 0 {
		return task.Core{}, errors.New("invalid id")
	}
	//validasi inputan
	// ...
	res, err := service.taskData.GetTaskByID(id)
	return res, err
}

// Delete implements task.TaskServiceInterface.
func (service *taskService) Delete(input task.Core, id int) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	//validasi inputan
	// ...
	err := service.taskData.Delete(input, id)
	return err
}

// Update implements task.TaskServiceInterface.
func (service *taskService) Update(id int, input task.Core) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	//validasi inputan
	// ...
	err := service.taskData.Update(id, input)
	return err
}
