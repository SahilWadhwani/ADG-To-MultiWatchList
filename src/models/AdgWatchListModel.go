package models

import (
	"time"
)

type Watchlist struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	UserID        int64     `json:"userId"`
	WatchlistName string    `json:"watchlistName"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}

type ScriptMaster struct {
	ID         int64  `json:"id" gorm:"primaryKey"`
	ScriptName string `json:"scriptName"`
}

type WatchlistScript struct {
	ID          int64 `json:"id" gorm:"primaryKey"`
	WatchlistID int64 `json:"watchlistId"`
	ScriptsID   int64 `json:"scriptId"`
}

type AdgWatchlistRequest struct {
	Action      string  `json:"action" validate:"required,oneof=ADD DELETE GET"`
	ScriptID    int64   `json:"scriptId" validate:"required"`
	WatchlistID []int64 `json:"watchListId" validate:"required"`
	UserID      int64   `json:"userId" validate:"required"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
