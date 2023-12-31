package models

import (
	"dineflow-menu-services/configs"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Canteen struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Image_path string `json:"image_path"`
}

func GetAllCanteens() ([]Canteen, error) {
	var canteens []Canteen
	if err := configs.Db.Find(&canteens).Error; err != nil {
		return nil, err
	}

	return canteens, nil
}

func GetCanteenByID(canteenID string) (Canteen, error) {
	var canteen Canteen
	result := configs.Db.Where("id = ?", canteenID).First(&canteen)
	// fmt.Println(result)
	if result.RowsAffected == 0 {
		return Canteen{}, fmt.Errorf("the canteen id could not be found")
	}
	if result.Error != nil {
		return Canteen{}, result.Error
	}
	return canteen, nil
}

func CreateCanteen(newCanteen Canteen) error {
	result := configs.Db.Create(&newCanteen)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteCanteenByID(canteenID string) error {
	result := configs.Db.Where("ID = ?", canteenID).Delete(&Canteen{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("the canteen id could not be found")
	}
	return nil
}

func UpdateCanteenByID(canteenID string, updatedCanteen Canteen) error {
	// Check if the canteen with the given ID exists
	var existingCanteen Canteen
	find_result := configs.Db.First(&existingCanteen, "ID = ?", canteenID)
	if find_result.RowsAffected == 0 {
		return fmt.Errorf("the canteen id could not be found")
	}
	if find_result.Error != nil {
		return find_result.Error
	}
	fmt.Println(updatedCanteen)

	result := configs.Db.Model(&Canteen{}).Where("ID = ?", canteenID).Updates(updatedCanteen)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
