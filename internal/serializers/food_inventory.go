package serializers

import (
	db "github.com/librefitness/librefitness/internal/database"
)

type FoodInventoryRes struct {
	ID                 uint    `json:"id"`
	OffCode            string  `json:"off_code"`
	Favourite          bool    `json:"favourite"`
	Calories           float64 `json:"calories"`
	FatTotal           float64 `json:"fat_total"`
	FatSaturated       float64 `json:"fat_saturated"`
	FatPolyunsaturated float64 `json:"fat_polyunsaturated"`
	FatMonounsaturated float64 `json:"fat_monounsaturated"`
	FatTrans           float64 `json:"fat_trans"`
	FatCholesterol     float64 `json:"fat_cholesterol"`
	Sodium             float64 `json:"sodium"`
	Potassium          float64 `json:"potassium"`
	Carbs              float64 `json:"carbs"`
	Fibers             float64 `json:"fibers"`
	Sugars             float64 `json:"sugars"`
	Protein            float64 `json:"proteins"`
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
