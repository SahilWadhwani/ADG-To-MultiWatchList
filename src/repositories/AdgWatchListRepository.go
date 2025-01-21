package repositories

import (
	"fmt"
	"main/src/database"
	"main/src/models"

	"errors"

	"gorm.io/gorm"
)

type WatchlistRepository struct {
	db *gorm.DB
}

func NewWatchlistRepository() *WatchlistRepository {
	return &WatchlistRepository{
		db: database.DB,
	}
}

func (r *WatchlistRepository) AddScriptToWatchlists(scriptID int64, watchlistIDs []int64, userID int64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		// Veryfying if Script already exists
		var script models.ScriptMaster
		if err := tx.First(&script, scriptID).Error; err != nil {
			return errors.New("script not found")
		}

		// Verify watchlists belong to user - took a little bit of GPT help for gorm queries
		for _, watchlistID := range watchlistIDs {
			var watchlist models.Watchlist
			if err := tx.Where("id = ? AND user_id = ?", watchlistID, userID).First(&watchlist).Error; err != nil {
				return errors.New("watchlist not found")
			}

			// Check if script already exists in watchlist
			var exists models.WatchlistScript
			result := tx.Where("watchlist_id = ? AND scripts_id = ?", watchlistID, scriptID).First(&exists)
			if result.Error == nil {
				continue
			}

			// Add scrip to watchlist
			watchlistScrip := models.WatchlistScript{
				WatchlistID: watchlistID,
				ScriptsID:   scriptID,
			}
			if err := tx.Create(&watchlistScrip).Error; err != nil {
				return err
			}

		}
		return nil
	})
}

func (r *WatchlistRepository) DeleteScriptFromWatchlists(scripID int64, watchlistIDs []int64, userID int64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, watchlistID := range watchlistIDs {

			// Verify watchlist belongs to user
			var watchlist models.Watchlist
			if err := tx.Where("id = ?", watchlistID).First(&watchlist).Error; err != nil {
				return errors.New("watchlist not found")
			}

			// Delete script from watchlist
			result := tx.Where("watchlist_id = ? AND scripts_id = ?", watchlistID, scripID).
				Delete(&models.WatchlistScript{})
			if result.Error != nil {
				return result.Error
			}

		}
		return nil
	})
}

func (r *WatchlistRepository) GetWatchlistScrips(watchlistIDs []int64, userID int64) ([]models.ScriptMaster, error) {
	var scripts []models.ScriptMaster

	err := r.db.
		Select("DISTINCT script_master.*").
		Table("script_master").
		Joins("JOIN watchlist_scripts ON watchlist_scripts.scripts_id = script_master.id").
		Joins("JOIN watchlists ON watchlists.id = watchlist_scripts.watchlist_id").
		Where("watchlists.user_id = ? AND watchlist_scripts.watchlist_id IN ?", userID, watchlistIDs).
		Find(&scripts).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch scrips: %w", err)
	}

	return scripts, nil
}

func (r *WatchlistRepository) GetWatchlistsByScript(scriptID int64, userID int64) ([]models.Watchlist, error) {
	var watchlists []models.Watchlist

	err := r.db.
		Select("DISTINCT watchlists.*").
		Joins("JOIN watchlist_scripts ON watchlist_scripts.watchlist_id = watchlists.id").
		Where("watchlists.user_id = ? AND watchlist_scripts.scripts_id = ?", userID, scriptID).
		Find(&watchlists).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch watchlists: %w", err)
	}

	return watchlists, nil
}
