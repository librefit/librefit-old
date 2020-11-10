package serializers

import (
	"fmt"
	"time"

	db "github.com/librefitness/librefitness/internal/database"
)

type FluidRes struct {
	ID    uint    `json:"id"`
	Date  string  `json:"date"`
	Type  string  `json:"type"`
	Value float32 `json:"value"`
}

type FluidStats struct {
	Date  string  `json:"date"`
	Value float32 `json:"value"`
}

func FluidResponse(f *db.Fluid) FluidRes {
	return FluidRes{
		ID:    f.ID,
		Date:  f.Date.Format(time.RFC3339),
		Type:  f.Type,
		Value: f.Value,
	}
}

func FluidsResponse(m *[]db.Fluid) []FluidRes {
	var r []FluidRes

	for _, v := range *m {
		res := FluidResponse(&v)
		r = append(r, res)
	}

	// To handle empty slices treated as null: https://github.com/golang/go/issues/37711
	if r == nil {
		r = []FluidRes{}
	}

	return r
}

// Return the amount of fluid per day
func StatsTotalFluidsPerDay(f *[]db.Fluid) ([]FluidStats, error) {
	var r []FluidStats

	for _, v := range *f {
		date := time.Date(v.Date.Year(), v.Date.Month(), v.Date.Day(), 0, 0, 0, 0, v.Date.Location())

		k, found := fluidsFind(r, fmt.Sprint(date.Format("2006-01-02")))
		if !found {
			e := FluidStats{
				Date:  fmt.Sprint(date.Format("2006-01-02")),
				Value: v.Value,
			}

			r = append(r, e)
			continue
		}

		r[k].Value = r[k].Value + v.Value
	}

	return r, nil
}

func fluidsFind(slice []FluidStats, date string) (int, bool) {
	for i, item := range slice {
		if item.Date == date {
			return i, true
		}
	}
	return -1, false
}
