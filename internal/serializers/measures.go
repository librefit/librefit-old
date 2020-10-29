package serializers

import (
	db "github.com/librefitness/librefitness/internal/database"
)

type MeasurementRes struct {
	ID    uint    `json:"id"`
	Date  string  `json:"date"`
	Value float32 `json:"value"`
}

func MeasurementResponse(m *db.Measurement) MeasurementRes {
	return MeasurementRes{
		ID:    m.ID,
		Date:  m.Date.Format("2006-01-02"),
		Value: m.Value,
	}
}

func MeasurementsResponse(m *[]db.Measurement) []MeasurementRes {
	var r []MeasurementRes

	for _, v := range *m {
		res := MeasurementResponse(&v)
		r = append(r, res)
	}

	// To handle empty slices treated as null: https://github.com/golang/go/issues/37711
	if r == nil {
		r = []MeasurementRes{}
	}

	return r
}
