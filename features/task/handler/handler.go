package handler

import (
	"net/http"
	"strconv"
	"zidan/clean-arch/features/task"
	"zidan/clean-arch/utils/responses"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService task.TaskServiceInterface
}

func New(service task.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		taskService: service,
	}
}

func (handler *TaskHandler) CreateTask(c echo.Context) error {
	newTask := TaskRequest{}
	errBind := c.Bind(&newTask) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	//mapping dari request ke core
	taskCore := RequestToCore(newTask)

	errInsert := handler.taskService.CreateTask(taskCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data"+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *TaskHandler) DeleteTask(c echo.Context) error {
	taskID := c.Param("taskid")

	taskID_int, errConv := strconv.Atoi(taskID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	result, errRead := handler.taskService.GetTaskByID(taskID_int)
	if errRead != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errRead.Error(), nil))
	}

	errDel := handler.taskService.Delete(result, taskID_int)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data. "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}

func (handler *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("taskid")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	taskStatus := TaskRequest{}
	errBind := c.Bind(&taskStatus) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	//mapping dari request ke core
	taskStatusCore := RequestToCore(taskStatus)

	err := handler.taskService.Update(idParam, taskStatusCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}
