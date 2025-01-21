package routes

import (
	"main/src/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	watchlistHandler := handler.NewWatchlistHandler()

	// Watchlist routes
	watchlistGroup := router.Group("/watchlist")
	{
		watchlistGroup.POST("/multi-adg", watchlistHandler.AdgWatchlistAction)
	}

	return router
}
