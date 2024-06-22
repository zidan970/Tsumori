package routers

import (
	"zidan/clean-arch/app/middlewares"

	_dataUser "zidan/clean-arch/features/user/data"
	_userHandler "zidan/clean-arch/features/user/handler"
	_userService "zidan/clean-arch/features/user/service"

	_dataProject "zidan/clean-arch/features/project/data"
	_projectHandler "zidan/clean-arch/features/project/handler"
	_projectService "zidan/clean-arch/features/project/service"

	_dataTask "zidan/clean-arch/features/task/data"
	_taskHandler "zidan/clean-arch/features/task/handler"
	_taskService "zidan/clean-arch/features/task/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {

	userData := _dataUser.New(db)
	// userData := _userData.NewRaw(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	projectData := _dataProject.New(db)
	projectService := _projectService.New(projectData)
	projectHandlerAPI := _projectHandler.New(projectService)

	taskData := _dataTask.New(db)
	taskService := _taskService.New(taskData)
	taskHandlerAPI := _taskHandler.New(taskService)

	// User routes
	e.POST("/users", userHandlerAPI.CreateUser)
	e.POST("/users/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.GetProfile, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())

	// Project routes
	e.POST("/projects", projectHandlerAPI.CreateProject, middlewares.JWTMiddleware())
	e.GET("/projects", projectHandlerAPI.GetAllProjects, middlewares.JWTMiddleware())
	e.GET("/projects/:projectid", projectHandlerAPI.GetDetailProject, middlewares.JWTMiddleware())
	e.PUT("/projects/:projectid", projectHandlerAPI.UpdateProject, middlewares.JWTMiddleware())
	e.DELETE("/projects/:projectid", projectHandlerAPI.DeleteProject, middlewares.JWTMiddleware())

	// Task routes
	e.POST("/tasks", taskHandlerAPI.CreateTask, middlewares.JWTMiddleware())
	e.DELETE("/tasks/:taskid", taskHandlerAPI.DeleteTask, middlewares.JWTMiddleware())
	e.PUT("/tasks/:taskid", taskHandlerAPI.UpdateTask, middlewares.JWTMiddleware())
}
