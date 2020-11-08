package serializers

import (
	"time"

	db "github.com/librefitness/librefitness/internal/database"
)

type FoodDiaryRes struct {
	ID              uint    `json:"id"`
	MealTypeID      string  `json:"meal_type"`
	Quantity        float32 `json:"quantity"`
	Date            string  `json:"date"`
	QuantityUnitID  uint    `json:"quantity_unit"`
	Comments        string  `json:"comments"`
	FoodInventoryID uint    `json:"food_inventory_id"`
}

func FoodDiaryResponse(f *db.FoodDiary) FoodDiaryRes {
	return FoodDiaryRes{
		ID:              f.ID,
		MealTypeID:      f.MealTypeID,
		Quantity:        f.Quantity,
		QuantityUnitID:  f.QuantityUnitID,
		Date:            f.Date.Format(time.RFC3339),
		Comments:        f.Comments,
		FoodInventoryID: f.FoodInventoryID,
	}
}

func FoodDiariesResponse(f *[]db.FoodDiary) []FoodDiaryRes {
	var r []FoodDiaryRes

	for _, v := range *f {
		res := FoodDiaryResponse(&v)
		r = append(r, res)
	}

	// To handle empty slices treated as null: https://github.com/golang/go/issues/37711
	if r == nil {
		r = []FoodDiaryRes{}
	}

	return r
}
