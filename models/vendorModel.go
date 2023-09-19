package models

import (
	"database/sql"
	"fmt"

	"dineflow-menu-services/configs"

	_ "github.com/go-sql-driver/mysql"
)

type Status string

const (
	OPEN  Status = "Open"
	CLOSE Status = "Close"
)

type Vendor struct {
	// gorm.Model
	ID               int          `json:"id"`
	CanteenID        int          `json:"canteen_id"`
	Name             string       `json:"name"`
	OwnerID          sql.NullBool `json:"owner_id"`
	OpeningTimestamp string       `json:"opening_timestamp"`
	ClosingTimestamp string       `json:"closing_timestamp"`
	Status           Status       `json:"status"`
}

func GetAllVendors() ([]Vendor, error) {
	var vendors []Vendor
	if err := configs.Db.Find(&vendors).Error; err != nil {
		return nil, err
	}

	return vendors, nil
}

func GetVendorByID(vendorID string) (Vendor, error) {
	var vendor Vendor
	result := configs.Db.Where("id = ?", vendorID).First(&vendor)
	if result.RowsAffected == 0 {
		return Vendor{}, fmt.Errorf("the vendor id could not be found")
	}
	if result.Error != nil {
		return Vendor{}, result.Error
	}
	return vendor, nil
}

func CreateVendor(vendor Vendor) error {
	err := configs.Db.Create(&vendor).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteVendorByID(vendorID string) error {
	result := configs.Db.Delete(&Vendor{}, vendorID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("the vendor id could not be found")
	}

	return nil
}

func UpdateVendorByID(vendorID string, updatedVendor Vendor) error {
	result := configs.Db.Model(&Vendor{}).Where("id = ?", vendorID).Updates(updatedVendor)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("the vendor id could not be found")
	}

	return nil
}
