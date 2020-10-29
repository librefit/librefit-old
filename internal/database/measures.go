package database

import (
	"time"

	"gorm.io/gorm"
)

type Measurement struct {
	gorm.Model `binding:"-"`
	Type       string    `json:"type" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
	Value      float32   `gorm:"type:FLOAT" json:"value,string" binding:"required"`
	UserID     uint      `binding:"-"`
}

func FindManyMeasurement(UserID uint) ([]Measurement, error) {
	var m []Measurement
	err := DB.Find(&m, "user_id = ?", UserID).Error
	return m, err
}

func DeleteMeasurement(id string) error {
	var m Measurement
	if err := DB.First(&m, "id = ?", id).Error; err != nil {
		return err
	}
	return DB.Unscoped().Delete(m).Error
}

func FindOneMeasurement(UserID string) (Measurement, error) {
	var m Measurement
	err := DB.Find(&m, UserID).Error
	return m, err
}
