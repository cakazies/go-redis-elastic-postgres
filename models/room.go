package models

import (
	"github.com/cakazies/go-redis-elastic-postgres/configs"
	"github.com/jinzhu/gorm"
)

type (
	// RoomModel struct for setting this file like class
	RoomModel struct{}

	// Room struct for table room
	Room struct {
		Name  string `json:"name,omitempty" validate:"required"`
		Quota int    `json:"quota,omitempty" validate:"required"`
		gorm.Model
	}
)

// GetRooms function get all data room
func (RoomModel) GetRooms() ([]Room, error) {
	result := make([]Room, 0)
	err := configs.GetDB.Model(&result).Find(&result).Error
	return result, err
}

// GetRoom function for get data detail from table room
func (RoomModel) GetRoom(id string) (Room, error) {
	var result Room
	err := configs.GetDB.Model(&result).Where("id = ?", id).Scan(&result).Error
	return result, err
}

// InsertRoom function insert data in table room
func (RoomModel) InsertRoom(data Room) error {
	err := configs.GetDB.Create(&data).Error
	return err
}

// UpdateRoom function for update data in table room
func (RoomModel) UpdateRoom(id string, data Room) error {
	err := configs.GetDB.Model(&data).Where("id = ?", id).Updates(&data).Error
	return err
}

// DeleteRoom function delete data in table
func (RoomModel) DeleteRoom(data Room) error {
	err := configs.GetDB.Delete(data).Error
	return err
}
