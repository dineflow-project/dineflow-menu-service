package models

import (
	"dineflow-menu-services/configs"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Menu struct {
	ID           int     `json:"id"`
	VendorID     int     `json:"vendor_id"`
	Name         string  `json:"name"`
	Price        float32 `json:"price"`
	Image_path   string  `json:"image_path"`
	Description  string  `json:"description"`
	Is_available int     `json:"is_available"`
}

func GetAllMenus(canteenId, vendorId int, minprice, maxprice float64) ([]Menu, error) {
	var menus []Menu
	query := configs.Db.Table("menus").
		Select("menus.*").
		Joins("join vendors on menus.vendor_id = vendors.id").
		Joins("join canteens on vendors.canteen_id = canteens.id")

	if canteenId > 0 {
		query = query.Where("canteens.id = ?", canteenId)
	}
	if vendorId > 0 {
		query = query.Where("menus.vendor_id = ?", vendorId)
	}
	if minprice > 0 {
		query = query.Where("menus.price >= ?", minprice)
	}
	if maxprice > 0 {
		query = query.Where("menus.price <= ?", maxprice)
	}

	if err := query.Find(&menus).Error; err != nil {
		return nil, err
	}

	return menus, nil
}

func GetAllMenusByVendorID(vendorID string) ([]Menu, error) {
	var menus []Menu
	result := configs.Db.Where("vendor_id = ?", vendorID).Find(&menus)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("the vendor id could not be founded")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return menus, nil
}

func GetMenuByID(menuID string) (Menu, error) {
	var menu Menu
	menuIDint, err := strconv.Atoi(menuID)
	if err != nil {
		fmt.Println("Error:", err)
		return Menu{}, err
	}
	result := configs.Db.Where("id = ?", menuIDint).First(&menu)
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

func CreateMenu(menu Menu) (Menu, error) {
	// Check if the vendor exists
	var vendorCount int64
	if err := configs.Db.Model(&Vendor{}).Where("id = ?", menu.VendorID).Count(&vendorCount).Error; err != nil {
		return Menu{}, err
	}

	if vendorCount == 0 {
		// Return a specific error indicating that the vendor does not exist
		return Menu{}, VendorNotFoundError{VendorID: menu.VendorID}
	}

	// The vendor exists, so create the menu item
	err := configs.Db.Create(&menu).Error
	return menu, err
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
	var existingMenu Menu
	find_result := configs.Db.First(&existingMenu, "ID = ?", menuID)
	if find_result.RowsAffected == 0 {
		return fmt.Errorf("the menu id could not be found")
	}
	if find_result.Error != nil {
		return find_result.Error
	}
	fmt.Println(existingMenu)
	fmt.Println(updatedMenu)
	result := configs.Db.Model(&Menu{}).Where("id = ?", menuID).Updates(updatedMenu)
	// fmt.Println("result", result)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
