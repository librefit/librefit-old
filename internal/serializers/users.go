package serializers

import (
	"time"

	db "github.com/librefitness/librefitness/internal/database"
)

type UserRes struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	IsAdmin   bool      `json:"is_admin"`
	Language  string    `json:"language"`
}

func UserResponse(u *db.User) UserRes {
	return UserRes{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		Username:  u.Username,
		IsAdmin:   u.IsAdmin,
		Language:  u.Language,
	}
}
