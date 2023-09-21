package models

import (
	"dineflow-menu-services/configs"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Menu struct {
	ID          int     `json:"id"`
	VendorID    int     `json:"vendor_id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	ImagePath   string  `json:"image_path"`
	Description string  `json:"description"`
    IsAvailable bool    `json:"is_available"`
}

func GetAllMenus() ([]Menu, error) {
	var menus []Menu
	if err := configs.Db.Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}

func GetAllMenusByVendorID(vendorID string) ([]Menu, error) {
	var menus []Menu
	if err := configs.Db.Where("vendor_id = ?", vendorID).Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}


func GetMenuByID(menuID string) (Menu, error) {
	var menu Menu
	result := configs.Db.Where("id = ?", menuID).First(&menu)
	if result.RowsAffected == 0 {
		return Menu{}, fmt.Errorf("the menu id could not be founded")
	}
	if result.Error != nil {
		return Menu{}, result.Error
	}
	return menu, nil
}

// Function to create a menu item
type VendorNotFoundError struct {
    VendorID int
}

func (e VendorNotFoundError) Error() string {
    return fmt.Sprintf("Vendor with ID %d does not exist", e.VendorID)
}

func IsVendorNotFoundError(err error) bool {
    _, ok := err.(VendorNotFoundError)
    return ok
}

func CreateMenu(menu Menu) error {
    // Check if the vendor exists
    var vendorCount int64
    if err := configs.Db.Model(&Vendor{}).Where("id = ?", menu.VendorID).Count(&vendorCount).Error; err != nil {
        return err
    }

    if vendorCount == 0 {
        // Return a specific error indicating that the vendor does not exist
        return VendorNotFoundError{VendorID: menu.VendorID}
    }

    // The vendor exists, so create the menu item
    err := configs.Db.Create(&menu).Error
    return err
}


func DeleteMenuByID(menuID string) error {
	result := configs.Db.Delete(&Menu{}, menuID)

	if result.RowsAffected == 0 {
		return fmt.Errorf("the menu id could not be founded")
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateMenuByID(menuID string, updatedMenu Menu) error {
	result := configs.Db.Model(&Menu{}).Where("id = ?", menuID).Updates(updatedMenu)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("the menu id could not be founded")
	}

	return nil
}
