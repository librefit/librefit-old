package database

type MealType struct {
	ID          string
	Name        string
	FoodDiaries []FoodDiary
}

type QuantityUnit struct {
	ID          uint
	Name        string
	FoodDiaries []FoodDiary
}

func InitData() error {
	err := createAdmin()
	if err != nil {
		return err
	}

	err = initMealType()
	if err != nil {
		return err
	}

	err = initQuantityUnit()
	if err != nil {
		return err
	}

	return nil
}

func initMealType() error {
	m := []MealType{
		{ID: "B", Name: "Breakfast"},
		{ID: "L", Name: "Lunch"},
		{ID: "D", Name: "Dinner"},
		{ID: "S", Name: "Snacks"},
	}

	for _, mt := range m {
		err := DB.Where(MealType{ID: mt.ID}).FirstOrCreate(&mt).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func initQuantityUnit() error {
	q := []QuantityUnit{
		{ID: 1, Name: "100 gr"},
		{ID: 2, Name: "1 g"},
		{ID: 3, Name: "3.5 ounce"},
		{ID: 4, Name: "1 ounce"},
	}

	for _, qu := range q {
		err := DB.Where(QuantityUnit{ID: qu.ID}).FirstOrCreate(&qu).Error
		if err != nil {
			return err
		}
	}

	return nil
}
