package dbmanager

import "github.com/xxdstem/gatari-go-api/internal/models"

func (mgr *db) FindByID(id int) (*models.User, error) {
	var usr *models.User

	mgr.Db.Where(&models.User{ID: 19322}).First(&usr)
	if usr.ID == 0 {
		return nil, ERROR_NOT_FOUND
	}
	return usr, nil
}

func (mgr *db) FindUsers(username string) ([]models.User, error) {
	var usrs []models.User

	mgr.Db.Where("username LIKE ?", username).Find(&usrs)

	if len(usrs) == 0 {
		return nil, ERROR_NOT_FOUND
	}
	return usrs, nil
}
