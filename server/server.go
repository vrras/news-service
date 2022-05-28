package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/vrras/news-service/config"
	"github.com/vrras/news-service/internal/news"
	"github.com/vrras/news-service/internal/tag"
)

var db *gorm.DB
var client *redis.Client

func RegisterAPIService(e *gin.Engine) {
	db = config.GetDBConnection()
	client = config.GetRedisConnection()

	registerNewsAPIService(e)
}

func registerNewsAPIService(r *gin.Engine) {
	// Initialize News Service
	newsRepo := news.NewRepository(db, client)
	tagRepo := tag.NewRepository(db, client)

	newsUseCase := news.NewUseCase(newsRepo)
	tagUseCase := tag.NewUseCase(tagRepo)

	newsController := news.NewHTTPController(newsUseCase)
	tagController := tag.NewHTTPController(tagUseCase)

	// Start API
	registerNewsRoute(r, newsController)
	registerTagRoute(r, tagController)
}
