package main

import (
	"fmt"

	"github.com/helloujin/go_mod_8.8.2/auto"
)

func main() {

	// реализация для Додж
	dodge := auto.NewAuto(auto.CM, "Dodge", "RAM TRX", 5780, 2217, 1950, 190, 880)
	fmt.Println("Brand", dodge.Brand(), "Model", dodge.Model(), "EnginePower", dodge.EnginePower(), "MaxSpeed", dodge.MaxSpeed())
	dodgeDimensions := dodge.Dimensions()
	fmt.Println("Height", dodgeDimensions.Height().Get(auto.Inch), "Width", dodgeDimensions.Width().Get(auto.Inch), "Length", dodgeDimensions.Length().Get(auto.Inch))

	// реализация для БМВ
	bmw := auto.NewAuto(auto.CM, "BMW", "M5", 2000.5, 500.4, 360.5, 320, 650)
	fmt.Println("\n\nBrand", bmw.Brand(), "Model", bmw.Model(), "EnginePower", bmw.EnginePower(), "MaxSpeed", bmw.MaxSpeed())
	bmwDimensions := bmw.Dimensions()
	fmt.Println("Height", bmwDimensions.Height().Get(auto.CM), "Width", bmwDimensions.Width().Get(auto.CM), "Length", bmwDimensions.Length().Get(auto.CM))

	// реализация для Мерседес
	mercedes := auto.NewAuto(auto.CM, "Mercedes", "S63 AMG", 2200.2, 490.3, 360.4, 320, 640)
	fmt.Println("\n\nBrand", mercedes.Brand(), "Model", mercedes.Model(), "EnginePower", mercedes.EnginePower(), "MaxSpeed", mercedes.MaxSpeed())
	mercedesDimensions := mercedes.Dimensions()
	fmt.Println("Height", mercedesDimensions.Height().Get(auto.CM), "Width", mercedesDimensions.Width().Get(auto.CM), "Length", mercedesDimensions.Length().Get(auto.CM))
}
