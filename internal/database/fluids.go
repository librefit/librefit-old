package database

import (
	"time"

	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

type Fluid struct {
	gorm.Model `binding:"-"`
	Type       string    `json:"type" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
	Value      float32   `gorm:"type:FLOAT" json:"value,string" binding:"required"`
	UserID     uint      `binding:"-"`
}

type FluidSearch struct {
	UserID uint
	From   time.Time
	To     time.Time
}

func FindFluids(f *FluidSearch) ([]Fluid, error) {
	var r []Fluid

	tx := DB.Where("user_id = ?", f.UserID)

	if !f.From.IsZero() && !f.To.IsZero() {
		tx.Where("date BETWEEN ? AND ?", f.From, now.With(f.To).EndOfDay())
	}

	err := tx.Find(&r).Error

	return r, err
}

func DeleteFluid(id string) error {
	var f Fluid
	if err := DB.First(&f, "id = ?", id).Error; err != nil {
		return err
	}
	return DB.Unscoped().Delete(f).Error
}

func FindOneFluid(UserID string) (Fluid, error) {
	var f Fluid
	err := DB.Find(&f, UserID).Error
	return f, err
}
