package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter creează un nou router
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	return router
}
