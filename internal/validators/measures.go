package validators

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
)

type MeasurementValidator struct {
	Measurement struct {
		Type  string `json:"type" binding:"required"`
		Date  string `json:"date" binding:"required"`
		Value string `json:"value" binding:"required"`
	} `json:"measure"`
	MeasurementDb db.Measurement `json:"-"`
}

func NewMeasurementValidatorFillWith(m db.Measurement) MeasurementValidator {
	mv := NewMeasurementValidator()
	mv.Measurement.Type = m.Type
	mv.Measurement.Date = fmt.Sprintf("%s", m.Date)
	mv.Measurement.Value = fmt.Sprintf("%f", m.Value)
	return mv
}

func (self *MeasurementValidator) Bind(c *gin.Context) error {
	claims := jwt.ExtractClaims(c)

	if err := c.ShouldBindJSON(&self.Measurement); err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", self.Measurement.Date)
	if err != nil {
		return err
	}

	value, err := strconv.ParseFloat(self.Measurement.Value, 32)
	if err != nil {
		return err
	}

	self.MeasurementDb.Type = self.Measurement.Type
	self.MeasurementDb.Date = date
	self.MeasurementDb.Value = float32(value)
	self.MeasurementDb.UserID = uint(claims["UserID"].(float64))

	return nil
}

func NewMeasurementValidator() MeasurementValidator {
	MeasurementValidator := MeasurementValidator{}
	return MeasurementValidator
}
