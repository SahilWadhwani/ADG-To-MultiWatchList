package handler

import (
	"main/src/business"
	"main/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WatchlistHandler struct {
	business *business.WatchlistBusiness
}

func NewWatchlistHandler() *WatchlistHandler {
	return &WatchlistHandler{
		business: business.NewWatchlistBusiness(),
	}
}

func (h *WatchlistHandler) AdgWatchlistAction(c *gin.Context) {
	var req models.AdgWatchlistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request",
			"error":   err.Error(),
		})
		return
	}

	// For GET action
	if req.Action == "GET" {
		scrips, err := h.business.GetWatchlistScrips(req.WatchlistID, req.UserID)
		if err != nil {
			status := http.StatusInternalServerError
			if err.Error() == "no watchlists found" || err.Error() == "watchlist not found or unauthorized" {
				status = http.StatusNotFound
			}
			c.JSON(status, gin.H{
				"status":  status,
				"message": "Failed to fetch scrips",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Successfully fetched scrips",
			"data":    scrips,
		})
		return
	}

	// For ADD and DELETE actions
	err := h.business.ProcessMultiWatchlistAction(req)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "scrip not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{
			"status":  status,
			"message": "Operation failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Operation successful",
	})
}
