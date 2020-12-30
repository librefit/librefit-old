package serializers

import (
	"fmt"

	db "github.com/librefitness/librefitness/internal/database"
)

type UserSettingRes struct {
	ID            uint    `json:"id"`
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
	Avatar        string  `json:"avatar"`
	ActivityLevel string  `json:"activity_level"`
}

func UserSettingResponse(s *db.UserSetting) UserSettingRes {
	return UserSettingRes{
		ID:            s.ID,
		Country:       s.Country,
		Timezone:      s.Timezone,
		Language:      s.Language,
		Theme:         s.Theme,
		FullName:      s.FullName,
		Email:         s.Email,
		Birthday:      fmt.Sprintf("%s", s.Birthday.Format("2006-01-02")),
		UseMetric:     s.UseMetric,
		Sex:           s.Sex,
		Height:        s.Height,
		BodyFat:       s.BodyFat,
		Avatar:        s.Avatar,
		ActivityLevel: s.ActivityLevel,
	}
}
