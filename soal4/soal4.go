package soal4

import (
	"fmt"
	"math"
)

func Main() {
	fmt.Println(convertTemperature(200, "celcius", "fahrenheit"))
}

func convertTemperature(initialTemperature float64, initialScale string, targetScale string) string {

	var finalTemperature float64

	switch initialScale {
	case "celcius":
		if targetScale == "reamur" {
			finalTemperature = (initialTemperature * 4) / 5
		} else if targetScale == "kelvin" {
			finalTemperature = initialTemperature + 273
		} else if targetScale == "fahrenheit" {
			finalTemperature = (9/5)*initialTemperature + 32
		}
	case "reamur":
		if targetScale == "celcius" {
			finalTemperature = (5 / 4) * initialTemperature
		} else if targetScale == "kelvin" {
			finalTemperature = (5/4)*initialTemperature + 32
		} else if targetScale == "fahrenheit" {
			finalTemperature = (9/4)*initialTemperature + 32
		}
	case "fahrenheit":
		if targetScale == "celcius" {
			finalTemperature = (5.0 / 9.0) * (initialTemperature - 32)
		} else if targetScale == "reamur" {
			finalTemperature = (4.0 / 9.0) * (initialTemperature - 32)
		} else if targetScale == "kelvin" {

			finalTemperature = (5.0/9.0)*(initialTemperature-32) + 273
		}
	case "kelvin":
		if targetScale == "celcius" {
			finalTemperature = initialTemperature - 273
		} else if targetScale == "reamur" {
			finalTemperature = (4.0 / 5.0) * (initialTemperature - 273)
		} else if targetScale == "fahrenheit" {

			finalTemperature = (9/5)*(initialTemperature-273) + 32
		}
	}

	return fmt.Sprintf("convert from temperature %.0f° from scale %s to scale %s is %.0f°", initialTemperature, initialScale, targetScale, math.Floor(finalTemperature))
}
