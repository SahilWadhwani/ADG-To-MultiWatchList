package business

import (
	"errors"
	"main/src/models"
	"main/src/repositories"
)

type WatchlistBusiness struct {
	repo *repositories.WatchlistRepository
}

func NewWatchlistBusiness() *WatchlistBusiness {
	return &WatchlistBusiness{
		repo: repositories.NewWatchlistRepository(),
	}
}

func (b *WatchlistBusiness) ProcessMultiWatchlistAction(req models.AdgWatchlistRequest) error {

	switch req.Action {
	case "ADD":
		return b.repo.AddScriptToWatchlists(req.ScriptID, req.WatchlistID, req.UserID)
	case "DELETE":
		return b.repo.DeleteScriptFromWatchlists(req.ScriptID, req.WatchlistID, req.UserID)
	case "GET":
		_, err := b.repo.GetWatchlistsByScript(req.ScriptID, req.UserID)
		return err
	default:
		return errors.New("invalid action")
	}
}

func (b *WatchlistBusiness) GetWatchlistScrips(watchlistIDs []int64, userID int64) ([]models.ScriptMaster, error) {
	// Directly call the repository method with correct types
	scrips, err := b.repo.GetWatchlistScrips(watchlistIDs, userID)
	if err != nil {
		return nil, err
	}
	return scrips, nil
}

func (b *WatchlistBusiness) GetWatchlistsByScript(scripID int64, userID int64) ([]models.Watchlist, error) {
	return b.repo.GetWatchlistsByScript(scripID, userID)
}
