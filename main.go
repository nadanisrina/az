package main

import (
	"fmt"

	calculate "github.com/az/calculate"
	models "github.com/az/models"
)

func main() {
	var chicken int64
	var truck int64
	var maintenace int64
	var operational int64

	fmt.Println("chicken : ")
	fmt.Scanf("%d", &chicken)
	fmt.Println("truck : ")
	fmt.Scanf("%d", &truck)
	fmt.Println("maintenace : ")
	fmt.Scanf("%d", &maintenace)
	fmt.Println("operational : ")
	fmt.Scanf("%d", &operational)

	chickenArg := models.Chicken{
		Amount: chicken,
	}
	truckArg := models.Truck{
		Amount: truck,
	}

	fmt.Println(calculate.Calculate(chickenArg, truckArg))
}
