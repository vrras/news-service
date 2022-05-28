package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vrras/news-service/internal/news"
)

func registerNewsRoute(r *gin.Engine, newsController news.HTTPController) {
	newsRouter := r.Group("/v1/news")
	newsRouter.GET("/", newsController.FindAll)
	newsRouter.GET("/:id", newsController.FindByID)
	newsRouter.POST("/", newsController.Add)
	newsRouter.PUT("/:id", newsController.Update)
	newsRouter.DELETE("/:id", newsController.Delete)
}

func registerTagRoute(r *gin.Engine, tagController news.HTTPController) {
	tagRouter := r.Group("/v1/tags")
	tagRouter.GET("/", tagController.FindAll)
	tagRouter.GET("/:id", tagController.FindByID)
	tagRouter.POST("/", tagController.Add)
	tagRouter.PUT("/:id", tagController.Update)
	tagRouter.DELETE("/:id", tagController.Delete)
}
