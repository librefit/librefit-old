package serializers

import (
	"time"

	db "github.com/librefitness/librefitness/internal/database"
)

type FoodDiaryRes struct {
	ID              uint             `json:"id"`
	MealTypeID      string           `json:"meal_type"`
	Quantity        float32          `json:"quantity"`
	Date            string           `json:"date"`
	QuantityUnitID  uint             `json:"quantity_unit"`
	Comments        string           `json:"comments"`
	FoodInventoryID uint             `json:"food_inventory_id"`
	FoodInventory   FoodInventoryRes `json:"food_inventory"`
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
		FoodInventory: FoodInventoryRes{
			ID:                 f.FoodInventory.ID,
			OffCode:            f.FoodInventory.OffCode,
			Favourite:          f.FoodInventory.Favourite,
			Calories:           f.FoodInventory.Calories,
			FatTotal:           f.FoodInventory.FatTotal,
			FatSaturated:       f.FoodInventory.FatSaturated,
			FatPolyunsaturated: f.FoodInventory.FatPolyunsaturated,
			FatMonounsaturated: f.FoodInventory.FatMonounsaturated,
			FatTrans:           f.FoodInventory.FatTrans,
			FatCholesterol:     f.FoodInventory.FatCholesterol,
			Sodium:             f.FoodInventory.Sodium,
			Potassium:          f.FoodInventory.Potassium,
			Carbs:              f.FoodInventory.Carbs,
			Fibers:             f.FoodInventory.Fibers,
			Sugars:             f.FoodInventory.Sugars,
			Protein:            f.FoodInventory.Proteins,
		},
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
