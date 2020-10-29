package serializers

import (
	"time"

	db "github.com/librefitness/librefitness/internal/database"
)

type FluidRes struct {
	ID    uint    `json:"id"`
	Date  string  `json:"date"`
	Type  string  `json:"type"`
	Value float32 `json:"value"`
}

func FluidResponse(f *db.Fluid) FluidRes {
	return FluidRes{
		ID: f.ID,
		// Date:  f.Date.Format("2006-01-02"),
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
