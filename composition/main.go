package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	// composition
	Vehicle Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of passnegers:", v.NumberOfPassengers)
	fmt.Println("Number of wheels:", v.NumberOfWheels)
	fmt.Println()
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Is electric:", c.IsElectric)
	fmt.Println("Is hybrid:", c.IsHybrid)
	// execute the method of the Vehicle type
	c.Vehicle.showDetails()

}

func main() {
	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 6,
	}

	volvoXC90 := Car{
		Make:       "Volvo",
		Model:      "XC90 T8",
		Year:       2021,
		IsElectric: false,
		IsHybrid:   true,
		Vehicle:    suv,
	}

	volvoXC90.show()

	teslaModelX := Car{
		Make:       "Tesla",
		Model:      "Model X",
		Year:       2021,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	teslaModelX.Vehicle.NumberOfPassengers = 7

	teslaModelX.show()

}
