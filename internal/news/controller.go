package news

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vrras/news-service/internal/pkg/common/http/response"
)

type HTTPController interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewHTTPController(newsUseCase UseCase) HTTPController {
	return &httpController{
		newsUseCase: newsUseCase,
	}
}

type httpController struct {
	newsUseCase UseCase
}

func (controller *httpController) FindAll(c *gin.Context) {
	news, err := controller.newsUseCase.FindAll(c.Request.Context(), c.Query("topic"), c.Query("status"))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, news)
}

func (controller *httpController) FindByID(c *gin.Context) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.newsUseCase.FindByID(c.Request.Context(), reqID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, result)
}

func (controller *httpController) Add(c *gin.Context) {
	var spec News
	err := c.ShouldBindJSON(&spec)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.newsUseCase.Add(c.Request.Context(), spec)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusCreated, result)
}

func (controller *httpController) Update(c *gin.Context) {
	var spec News
	var err error
	err = c.ShouldBindJSON(&spec)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	spec.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.newsUseCase.Update(c.Request.Context(), spec)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, result)
}

func (controller *httpController) Delete(c *gin.Context) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	err = controller.newsUseCase.Delete(c.Request.Context(), reqID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, nil)
}
