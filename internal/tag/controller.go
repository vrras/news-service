package tag

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

func NewHTTPController(tagUseCase UseCase) HTTPController {
	return &httpController{
		tagUseCase: tagUseCase,
	}
}

type httpController struct {
	tagUseCase UseCase
}

func (controller *httpController) FindAll(c *gin.Context) {
	tag, err := controller.tagUseCase.FindAll(c.Request.Context())

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, tag)
}

func (controller *httpController) FindByID(c *gin.Context) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.tagUseCase.FindByID(c.Request.Context(), reqID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, result)
}

func (controller *httpController) Add(c *gin.Context) {
	var spec Tag
	err := c.ShouldBindJSON(&spec)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.tagUseCase.Add(c.Request.Context(), spec)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusCreated, result)
}

func (controller *httpController) Update(c *gin.Context) {
	var spec Tag
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

	result, err := controller.tagUseCase.Update(c.Request.Context(), spec)
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

	err = controller.tagUseCase.Delete(c.Request.Context(), reqID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, nil)
}
