package handler

import (
	"net/http"
	"strconv"
	"zidan/clean-arch/app/middlewares"
	"zidan/clean-arch/features/project"
	"zidan/clean-arch/utils/responses"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService project.ProjectServiceInterface
}

func New(service project.ProjectServiceInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService: service,
	}
}

// insert project
func (handler *ProjectHandler) CreateProject(c echo.Context) error {
	// Mendapatkan ID pengguna dari token JWT
	userID := middlewares.ExtractTokenUserId(c)

	newProject := ProjectRequest{}
	errBind := c.Bind(&newProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	// Menggunakan ID pengguna untuk membuat proyek
	newProject.UserID = uint(userID)

	projectCore := RequestToCore(newProject)

	errCreate := handler.projectService.Create(projectCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data"+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *ProjectHandler) GetAllProjects(c echo.Context) error {
	// Mengekstrak ID pengguna
	userID := middlewares.ExtractTokenUserId(c)

	// Mengambil semua proyek milik pengguna dengan ID tersebut
	result, errFind := handler.projectService.GetProjectsByUserID(userID)
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errFind.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", result))
}

func (handler *ProjectHandler) GetDetailProject(c echo.Context) error {
	// Mendapatkan parameter ID proyek dari URL
	projectID, err := strconv.Atoi(c.Param("projectid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid project ID"})
	}

	result, err := handler.projectService.GetDetail(projectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", result))
}

// update data project
func (handler *ProjectHandler) UpdateProject(c echo.Context) error {
	// Mendapatkan ID pengguna dari token JWT
	userID := middlewares.ExtractTokenUserId(c)

	// Mendapatkan ID proyek dari parameter URL
	projectID, err := strconv.Atoi(c.Param("projectid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("invalid project ID", nil))
	}

	// Mendapatkan data project dari body request
	updatedProject := ProjectRequest{}
	errBind := c.Bind(&updatedProject)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	// Validasi apakah pengguna memiliki hak untuk mengupdate proyek
	isAuthorized, err := handler.projectService.IsUserAuthorizedToUpdate(userID, projectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error checking authorization", nil))
	}
	if !isAuthorized {
		return c.JSON(http.StatusForbidden, responses.WebResponse("unauthorized to update project", nil))
	}

	updatedProjectCore := RequestToCore(updatedProject)

	// Mengupdate proyek
	errUpdate := handler.projectService.Update(projectID, updatedProjectCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error updating project. "+errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success updating project", nil))
}

func (handler *ProjectHandler) DeleteProject(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)

	ProjectID := c.Param("projectid")

	ProjectID_int, errConv := strconv.Atoi(ProjectID)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error convert id param", nil))
	}

	result, errRead := handler.projectService.GetAll()
	if errRead != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errRead.Error(), nil))
	}

	// checking authorization
	isAuthorized, err := handler.projectService.IsUserAuthorizedToDelete(userID, ProjectID_int)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error checking authorization", nil))
	}
	if !isAuthorized {
		return c.JSON(http.StatusForbidden, responses.WebResponse("unauthorized to delete project", nil))
	}

	errDel := handler.projectService.Delete(result, ProjectID_int)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data. "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}
