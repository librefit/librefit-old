package validators

import (
	"github.com/gin-gonic/gin"
	db "github.com/librefitness/librefitness/internal/database"
)

type UserValidator struct {
	User struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		OldPassword string `json:"old_password"`
		IsAdmin     bool   `json:"is_admin"`
	} `json:"user"`
	UserDb db.User `json:"-"`
}

func NewUserValidatorFillWith(u db.User) UserValidator {
	uv := NewUserValidator()
	uv.User.Username = u.Username
	uv.User.Password = u.Password
	uv.User.IsAdmin = u.IsAdmin
	uv.UserDb.UserSettingID = u.UserSettingID

	return uv
}

func (self *UserValidator) Bind(c *gin.Context) error {
	if err := c.ShouldBindJSON(&self.User); err != nil {
		return err
	}

	self.UserDb.Username = self.User.Username
	password, err := db.PasswordHash(self.User.Password)
	if err != nil {
		return err
	}
	self.UserDb.Password = password
	self.UserDb.IsAdmin = self.User.IsAdmin

	return nil
}

func NewUserValidator() UserValidator {
	return UserValidator{}
}
