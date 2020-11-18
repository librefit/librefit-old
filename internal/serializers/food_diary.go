package serializers

import (
	"strings"
	"time"

	db "github.com/librefitness/librefitness/internal/database"
)

type FoodDiaryRes struct {
	ID                     uint                   `json:"id"`
	MealTypeID             string                 `json:"meal_type"`
	Quantity               float32                `json:"quantity"`
	Date                   string                 `json:"date"`
	QuantityUnitID         uint                   `json:"quantity_unit"`
	Comments               string                 `json:"comments"`
	FoodInventoryID        uint                   `json:"food_inventory_id"`
	ComputedConsumptionRes ComputedConsumptionRes `json:"computed_consuption"`
	FoodInventory          FoodInventoryRes       `json:"food_inventory"`
}

type FoodDiaryStatsRes struct {
	Date     string `json:"date"`
	Carbs    int    `json:"carbs"`
	Calories int    `json:"calories"`
	Proteins int    `json:"proteins"`
	Fat      int    `json:"fat"`
	Fibers   int    `json:"fibers"`
}

type ComputedConsumptionRes struct {
	Calories           int `json:"calories"`
	FatTotal           int `json:"fat_total"`
	FatSaturated       int `json:"fat_saturated"`
	FatPolyunsaturated int `json:"fat_polyunsaturated"`
	FatMonounsaturated int `json:"fat_monounsaturated"`
	FatTrans           int `json:"fat_trans"`
	FatCholesterol     int `json:"fat_cholesterol"`
	Sodium             int `json:"sodium"`
	Potassium          int `json:"potassium"`
	Carbs              int `json:"carbs"`
	Fibers             int `json:"fibers"`
	Sugars             int `json:"sugars"`
	Proteins           int `json:"proteins"`
}

func FoodDiaryResponse(f *db.FoodDiary) FoodDiaryRes {
	return FoodDiaryRes{
		ID:                     f.ID,
		MealTypeID:             f.MealTypeID,
		Quantity:               f.Quantity,
		QuantityUnitID:         f.QuantityUnitID,
		Date:                   f.Date.Format(time.RFC3339),
		Comments:               f.Comments,
		ComputedConsumptionRes: computeConsumption(f),
		FoodInventoryID:        f.FoodInventoryID,
		FoodInventory: FoodInventoryRes{
			ID:                 f.FoodInventory.ID,
			OffCode:            f.FoodInventory.OffCode,
			ProductName:        f.FoodInventory.ProductName,
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
			Proteins:           f.FoodInventory.Proteins,
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

func FoodDiaryStatsResponse(f *[]db.FoodDiary) []FoodDiaryStatsRes {
	var r []FoodDiaryStatsRes

	for _, v := range *f {
		res := FoodDiaryResponse(&v)
		date := strings.Split(res.Date, "T")[0]
		if exist, index := contains(r, date); exist {
			r[index].Carbs = r[index].Carbs + res.ComputedConsumptionRes.Carbs
			r[index].Calories = r[index].Calories + res.ComputedConsumptionRes.Calories
			r[index].Fat = r[index].Fat + res.ComputedConsumptionRes.FatTotal
			r[index].Proteins = r[index].Fat + res.ComputedConsumptionRes.Proteins
			r[index].Fibers = r[index].Fat + res.ComputedConsumptionRes.Fibers
		} else {
			e := FoodDiaryStatsRes{
				Date:     date,
				Carbs:    res.ComputedConsumptionRes.Carbs,
				Calories: res.ComputedConsumptionRes.Calories,
				Fat:      res.ComputedConsumptionRes.FatTotal,
				Proteins: res.ComputedConsumptionRes.Proteins,
				Fibers:   res.ComputedConsumptionRes.Fibers,
			}
			r = append(r, e)
		}
	}

	// To handle empty slices treated as null: https://github.com/golang/go/issues/37711
	if r == nil {
		r = []FoodDiaryStatsRes{}
	}

	return r
}

func computeConsumption(f *db.FoodDiary) ComputedConsumptionRes {
	m := factorMultiplier(f.QuantityUnitID)

	r := ComputedConsumptionRes{
		Calories:           int(m * f.Quantity * f.FoodInventory.Calories),
		FatTotal:           int(m * f.Quantity * f.FoodInventory.FatTotal),
		FatSaturated:       int(m * f.Quantity * f.FoodInventory.FatSaturated),
		FatPolyunsaturated: int(m * f.Quantity * f.FoodInventory.FatPolyunsaturated),
		FatMonounsaturated: int(m * f.Quantity * f.FoodInventory.FatMonounsaturated),
		FatTrans:           int(m * f.Quantity * f.FoodInventory.FatTrans),
		FatCholesterol:     int(m * f.Quantity * f.FoodInventory.FatCholesterol),
		Sodium:             int(m * f.Quantity * f.FoodInventory.Sodium),
		Potassium:          int(m * f.Quantity * f.FoodInventory.Potassium),
		Carbs:              int(m * f.Quantity * f.FoodInventory.Carbs),
		Fibers:             int(m * f.Quantity * f.FoodInventory.Fibers),
		Sugars:             int(m * f.Quantity * f.FoodInventory.Sugars),
		Proteins:           int(m * f.Quantity * f.FoodInventory.Proteins),
	}

	return r
}

func factorMultiplier(f uint) float32 {
	switch f {
	case 1: // 100 grams
		return 1
	case 2: // 1 grams -> 1/100
		return 0.01
	case 3: // 3.5 ounces: it's really 99.22 grams but lets round to 100 grams)
		return 1
	case 4: // 1 ounce -> 1/28.3496
		return 0.035273
	}
	return 0
}

func contains(s []FoodDiaryStatsRes, date string) (bool, int) {
	for i, a := range s {
		if a.Date == date {
			return true, i
		}
	}
	return false, -1
}
