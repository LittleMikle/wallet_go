package models

import "gorm.io/gorm"

type Wallets struct {
	ID      uint    `gorm:"primary key;autoIncrement" json:"id"`
	Owner   *string `json:"owner"`
	Balance *int    `json:"balance"`
}

func MigrateWallets(db *gorm.DB) error {
	err := db.AutoMigrate(&Wallets{})
	return err
}
