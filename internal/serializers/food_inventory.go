package serializers

import (
	db "github.com/librefitness/librefitness/internal/database"
)

type FoodInventoryRes struct {
	ID                 uint    `json:"id"`
	OffCode            string  `json:"off_code"`
	ProductName        string  `json:"product_name"`
	Favourite          bool    `json:"favourite"`
	Calories           float32 `json:"calories"`
	FatTotal           float32 `json:"fat_total"`
	FatSaturated       float32 `json:"fat_saturated"`
	FatPolyunsaturated float32 `json:"fat_polyunsaturated"`
	FatMonounsaturated float32 `json:"fat_monounsaturated"`
	FatTrans           float32 `json:"fat_trans"`
	FatCholesterol     float32 `json:"fat_cholesterol"`
	Sodium             float32 `json:"sodium"`
	Potassium          float32 `json:"potassium"`
	Carbs              float32 `json:"carbs"`
	Fibers             float32 `json:"fibers"`
	Sugars             float32 `json:"sugars"`
	Proteins           float32 `json:"proteins"`
}

func FoodInventoryResponse(f *db.FoodInventory) FoodInventoryRes {
	return FoodInventoryRes{
		ID:        f.ID,
		OffCode:   f.OffCode,
		Favourite: f.Favourite,
	}
}

func FoodInventoriesResponse(f *[]db.FoodInventory) []FoodInventoryRes {
	var r []FoodInventoryRes

	for _, v := range *f {
		res := FoodInventoryResponse(&v)
		r = append(r, res)
	}

	// To handle empty slices treated as null: https://github.com/golang/go/issues/37711
	if r == nil {
		r = []FoodInventoryRes{}
	}

	return r
}
