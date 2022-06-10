package calculate

import (
	"math"

	models "github.com/az/models"
)

//example chickens amount are 50 then trucks needed are 50/17 = 3 trucks needed , 17 is truck operational time in 1 month
func Calculate(chicken models.Chicken, truck models.Truck) float64 {
	var amountTruck float64
	const weeksInMonth int64 = 4
	if chicken.Amount == 1 && truck.MaintenaceTime == 1 {
		amountTruck = 1
	}
	//case: one truck only can fit ten chickens
	if chicken.Amount > 1 {
		amountOP := float64(31 - (truck.OperationalTime * weeksInMonth) - truck.MaintenaceTime)
		amountTruck = math.Ceil(float64(chicken.Amount) / amountOP)
	}
	return amountTruck
}
