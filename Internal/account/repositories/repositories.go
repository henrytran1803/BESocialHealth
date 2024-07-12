package accountrepositories

import (
	"BESocialHealth/component/ws"
	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

type DashboardRepository struct {
	DB      *gorm.DB
	Manager *ws.WebSocketManager
}

func NewDashboardRepositoryy(db *gorm.DB, manager *ws.WebSocketManager) *DashboardRepository {
	return &DashboardRepository{DB: db, Manager: manager}
}
