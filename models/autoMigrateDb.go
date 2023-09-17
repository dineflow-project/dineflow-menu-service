package models

import (
	"dineflow-menu-services/configs"
)

// AutoMigrateDB automatically migrates the database to match the struct models.
func AutoMigrateDB() error {
	// Auto-migrate the Canteen model
	if err := configs.Db.AutoMigrate(&Canteen{}); err != nil {
		return err
	}

	// Auto-migrate the Vendor model
	if err := configs.Db.AutoMigrate(&Vendor{}); err != nil {
		return err
	}

	return nil
}
