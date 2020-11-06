package serializers

import (
	db "github.com/librefitness/librefitness/internal/database"
)

type FoodInventoryRes struct {
	ID        uint   `json:"id"`
	OffCode   string `json:"off_code"`
	Favourite bool   `json:"favourite"`
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
