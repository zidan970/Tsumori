package handler

import (
	"net/http"
	"zidan/clean-arch/app/middlewares"
	"zidan/clean-arch/features/user"
	"zidan/clean-arch/utils/responses"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	//mapping dari request ke core
	userCore := RequestToCore(newUser)

	errInsert := handler.userService.Create(userCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data"+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *UserHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}
	result, token, err := handler.userService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error login "+err.Error(), nil))
	}
	responseData := map[string]any{
		"token": token,
		"nama":  result.Name,
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success login", responseData))
}

func (handler *UserHandler) GetProfile(c echo.Context) error {
	// Menggunakan fungsi ExtractTokenUserId untuk mendapatkan ID pengguna dari token JWT
	userId := middlewares.ExtractTokenUserId(c)

	result, errFirst := handler.userService.GetSingle(userId)
	if errFirst != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errFirst.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", result))
}

func (handler *UserHandler) Update(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	var userData = UserRequest{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	userCore := RequestToCore(userData)
	err := handler.userService.Update(userId, userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data"+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *UserHandler) Delete(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	result, errFind := handler.userService.GetAll()
	if errFind != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errFind.Error(), nil))
	}

	errDel := handler.userService.DeleteUser(result, userId)
	if errDel != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data. "+errDel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}
