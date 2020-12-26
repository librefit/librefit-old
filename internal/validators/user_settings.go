package validators

import (
	"fmt"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	db "github.com/librefitness/librefitness/internal/database"
)

type UserSettingValidator struct {
	UserSetting struct {
		Country       string  `json:"country"`
		Timezone      string  `json:"timezone"`
		Language      string  `json:"language"`
		Theme         string  `json:"theme"`
		FullName      string  `json:"full_name"`
		Email         string  `json:"email"`
		Birthday      string  `json:"birthday"`
		UseMetric     bool    `json:"use_metric"`
		Sex           string  `json:"sex"`
		Height        float32 `json:"height"`
		BodyFat       string  `json:"body_fat"`
		ActivityLevel string  `json:"activity_level"`
	} `json:"user_settings"`
	UserSettingDb db.UserSetting `json:"-"`
}

func NewUserSettingsValidatorFillWith(u db.UserSetting) UserSettingValidator {
	usv := NewUserSettingsValidator()

	usv.UserSetting.Country = u.Country
	usv.UserSetting.Timezone = u.Timezone
	usv.UserSetting.Language = u.Language
	usv.UserSetting.Theme = u.Theme
	usv.UserSetting.FullName = u.FullName
	usv.UserSetting.Email = u.Email
	usv.UserSetting.Birthday = fmt.Sprintf("%s", u.Birthday.Format("2006-01-02"))
	usv.UserSetting.UseMetric = u.UseMetric
	usv.UserSetting.Sex = u.Sex
	usv.UserSetting.Height = u.Height
	usv.UserSetting.BodyFat = u.BodyFat
	usv.UserSetting.ActivityLevel = u.ActivityLevel

	return usv
}

func (self *UserSettingValidator) Bind(c *gin.Context) error {
	claims := jwt.ExtractClaims(c)

	if err := c.ShouldBindJSON(&self.UserSetting); err != nil {
		return err
	}

	birthday, err := time.Parse("2006-01-02", self.UserSetting.Birthday)
	if err != nil {
		return err
	}

	self.UserSettingDb.Country = self.UserSetting.Country
	self.UserSettingDb.Timezone = self.UserSetting.Timezone
	self.UserSettingDb.Language = self.UserSetting.Language
	self.UserSettingDb.Theme = self.UserSetting.Theme
	self.UserSettingDb.FullName = self.UserSetting.FullName
	self.UserSettingDb.Email = self.UserSetting.Email
	self.UserSettingDb.Birthday = birthday
	self.UserSettingDb.UseMetric = self.UserSetting.UseMetric
	self.UserSettingDb.Sex = self.UserSetting.Sex
	self.UserSettingDb.Height = self.UserSetting.Height
	self.UserSettingDb.BodyFat = self.UserSetting.BodyFat
	self.UserSettingDb.ActivityLevel = self.UserSetting.ActivityLevel
	self.UserSettingDb.UserID = uint(claims["UserID"].(float64))

	return nil
}

func NewUserSettingsValidator() UserSettingValidator {
	usv := UserSettingValidator{}
	return usv
}
