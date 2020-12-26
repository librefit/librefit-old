package validators

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
)

type FluidValidator struct {
	Fluid struct {
		Type  string `json:"type" binding:"required"`
		Date  string `json:"date" binding:"required"`
		Value string `json:"value" binding:"required"`
	}
	FluidDb db.Fluid `json:"-"`
}

type FluidSearchValidator struct {
	Fluid struct {
		From string `json:"from"`
		To   string `json:"to"`
	}
	FluidDb db.Fluid `json:"-"`
}

func (self *FluidSearchValidator) Bind(c *gin.Context, fs *db.FluidSearch) error {
	claims := jwt.ExtractClaims(c)

	if err := c.ShouldBindJSON(&self.Fluid); err != nil {
		return err
	}

	from, err := time.Parse("2006-01-02", self.Fluid.From)
	if err != nil {
		return err
	}

	to, err := time.Parse("2006-01-02", self.Fluid.To)
	if err != nil {
		return err
	}

	fs.From = from
	fs.To = to
	fs.UserID = uint(claims["UserID"].(float64))

	return nil
}

func NewFluidValidatorFillWith(f *db.Fluid) FluidValidator {
	fv := NewFluidValidator()
	fv.Fluid.Type = f.Type
	fv.Fluid.Date = fmt.Sprintf("%s", f.Date)
	fv.Fluid.Value = fmt.Sprintf("%f", f.Value)
	return fv
}

func (self *FluidValidator) Bind(c *gin.Context) error {
	claims := jwt.ExtractClaims(c)

	if err := c.ShouldBindJSON(&self.Fluid); err != nil {
		return err
	}

	date, err := time.Parse(time.RFC3339, self.Fluid.Date)
	if err != nil {
		return err
	}

	value, err := strconv.ParseFloat(self.Fluid.Value, 32)
	if err != nil {
		return err
	}

	self.FluidDb.Type = self.Fluid.Type
	self.FluidDb.Date = date
	self.FluidDb.Value = float32(value)
	self.FluidDb.UserID = uint(claims["UserID"].(float64))

	return nil
}

func NewFluidValidator() FluidValidator {
	FluidValidator := FluidValidator{}
	return FluidValidator
}
