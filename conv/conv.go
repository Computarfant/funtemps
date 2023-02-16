package conv

func FahrenheitToCelsius(value float64) float64 {

	return (value - 32) * 5 / 9
}

func KelvinToCelsius(value float64) float64 {
	//Celsius := value - 273.15
	return value - 273.15
}

func CelsiusToFahrenheit(value float64) float64 {
	//Farhrenheit := value*(9/5) + 32
	return value*9/5 + 32
}

func KelvinToFahrenheit(value float64) float64 {
	//Farhrenheit := (value-273.15)*(9/5) + 32
	return (value-273.15)*9/5 + 32
}

func FahrenheitToKelvin(value float64) float64 {
	//Kelvin := (value-32)*(5/9) + 273.15
	return (value-32)*5/9 + 273.15
}

func CelsiusToKelvin(value float64) float64 {
	//Kelvin := value + 273.15
	return value + 273.15
}

// func Funtemps(value float64, unit string, out string) string {
// 	var result float64
// 	var resultUnit string
// 	switch unit {
// 	case "F":
// 		if out == "C" {
// 			result = FahrenheitToCelsius(value)
// 			resultUnit = "°C"
// 		} else if out == "K" {
// 			result = FahrenheitToKelvin(value)
// 			resultUnit = "K"
// 		}
// 	case "C":
// 		if out == "F" {
// 			result = CelsiusToFahrenheit(value)
// 			resultUnit = "°F"
// 		} else if out == "K" {
// 			result = CelsiusToKelvin(value)
// 			resultUnit = "K"
// 		}
// 	case "K":
// 		if out == "C" {
// 			result = KelvinToCelsius(value)
// 			resultUnit = "°C"
// 		} else if out == "F" {
// 			result = KelvinToFahrenheit(value)
// 			resultUnit = "°F"
// 		}
// 	}
// 	return strconv.FormatFloat(value, 'f', 2, 64) + unit + " er " + strconv.FormatFloat(result, 'f', 2, 64) + resultUnit
// }
