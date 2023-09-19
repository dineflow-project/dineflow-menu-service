package models

import (
	"dineflow-menu-services/configs"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Available   Status = "Available"
	Unavailable Status = "Unavailable"
)

type Menu struct {
	ID          int     `json:"id"`
	VendorID    int     `json:"vendor_id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	ImagePath   string  `json:"image_path"`
	Description string  `json:"description"`
	Status      Status  `json:"status"`
}

func GetAllMenus() ([]Menu, error) {
	var menus []Menu
	if err := configs.Db.Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}

func GetAllMenusByVendorID(vendorID string) (Menu, error) {
	var menus Menu
	if err := configs.Db.Where("vendor_id = ?", vendorID).Find(&menus).Error; err != nil {
		return Menu{}, err
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

// // Function to check if a vendor with the given ID exists
// func VendorExists(vendorDB *sql.DB, vendorID int) (bool, error) {
// 	var count int
// 	err := vendorDB.QueryRow("SELECT COUNT(*) FROM Vendor WHERE id = ?", vendorID).Scan(&count)
// 	if err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }

// Function to create a menu item
func CreateMenu(menu Menu) error {
	// Check if the vendor with the provided vendor_id exists
	// Convert the integer to a string
	// vendorIdStr := strconv.Itoa(menu.VendorID)

	// _,err := GetVendorByID(vendorIdStr)
	// if err != nil {
	// 	return err
	// }

	// The vendor exists, so we can create the menu item
	errr := configs.Db.Create(&menu).Error
	if errr != nil {
		return errr
	}

	return nil
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
