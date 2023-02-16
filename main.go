package main

import (
	"flag"
	"fmt"

	"github.com/Computarfant/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
// var fahr float64
// var celsius float64
// var kelvin float64
var out string
var temp float64
var unit string

//var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	// Du må selv definere flag-variablene for "C" og "K"

	// flag.Float64Var(&fahr, "F", 0.0, "temperature in degrees Fahrenheit")
	// flag.Float64Var(&celsius, "C", 0.0, "temperature in degrees Celsius")
	// flag.Float64Var(&kelvin, "K", 0.0, "temperature in Kelvin")
	flag.StringVar(&unit, "U", "", "Unit of tempratur")
	flag.Float64Var(&temp, "T", 0, "Temperatur in degrees")
	flag.StringVar(&out, "out", "", "calculate temperature in C - celsius, F - fahrenheit, K- kelvin")
}

func main() {
	flag.Parse()
	if unit == "" || out == "" {
		fmt.Println("Parametere U og out må ha verdi")
		return
	}

	switch unit {
	case "F":
		switch out {
		case "C":
			fmt.Printf("%g°F is %.2f°C", temp, conv.FahrenheitToCelsius(temp))
		case "K":
			fmt.Printf("%g°F is %.2f°K", temp, conv.FahrenheitToKelvin(temp))
		default:
			fmt.Printf("%s is not a valid -out value or can't be same as -unit value ", out)
		}
	case "C":
		switch out {
		case "F":
			fmt.Printf("%g°C is %.2f°F", temp, conv.CelsiusToFahrenheit(temp))
		case "K":
			fmt.Printf("%g°C is %.2f°K", temp, conv.CelsiusToKelvin(temp))
		default:
			fmt.Printf("%s is not a valid -out value or can't be same as -unit value", out)
		}
	case "K":
		switch out {
		case "F":
			fmt.Printf("%g°K is %.2f°F", temp, conv.KelvinToFahrenheit(temp))
		case "C":
			fmt.Printf("%g°K is %.2f°C", temp, conv.KelvinToCelsius(temp))
		default:
			fmt.Printf("%s is not a valid -out value or can't be same as -unit value", out)
		}
	default:
		fmt.Printf("%s is not a valid unit", unit)
	}

	/*
			Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
			pakkene implementeres.

			Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
			av flagg. flag-pakken har funksjoner som man kan bruke for å teste
			hvor mange flagg og argumenter er spesifisert på kommandolinje.

			fmt.Println("len(flag.Args())", len(flag.Args())
			fmt.Println("flag.NFlag()", flag.NFlag())

			Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
		    brukes for å utelukke ugyldige kombinasjoner:
		    -F, -C, -K kan ikke brukes samtidig
			disse tre kan brukes med -out, men ikke med -funfacts
		    -funfacts kan brukes kun med -t
		    ...
			Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
		    implementert i flag-pakken og at den vil skrive ut "Usage" med
			beskrivelsene av flagg-variablene, som angitt i parameter fire til
		    funksjonene Float64Var og StringVar
	*/

	// 	if (fahr != 0 && celsius != 0) || (fahr != 0 && kelvin != 0) || (celsius != 0 && kelvin != 0) {
	// 		fmt.Println("error: only one temperature argument should be specified")
	// 		return
	// 	}

	// 	if (fahr != 0 && out == "F") || (celsius != 0 && out == "C") || (kelvin != 0 && out == "K") {
	// 		fmt.Println("error: the specified temperature argument and the -out argument should not be the same")
	// 		return
	// 	}
	// 	// funker ikke for input 0
	// 	if fahr != 0  {
	// 		celsiusResult := conv.FahrenheitToCelsius(fahr)
	// 		kelvinResult := conv.FahrenheitToKelvin(fahr)
	// 		if out == "C" {
	// 			fmt.Printf("%.2f°F is %.2f°C\n", fahr, celsiusResult)
	// 		} else if out == "K" {
	// 			fmt.Printf("%.2f°F is %.2fK\n", fahr, kelvinResult)
	// 		}
	// 	} else if celsius != 0 {
	// 		fahrResult := conv.CelsiusToFahrenheit(celsius)
	// 		kelvinResult := conv.CelsiusToKelvin(celsius)
	// 		if out == "F" {
	// 			fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrResult)
	// 		} else if out == "K" {
	// 			fmt.Printf("%.2f°C is %.2fK\n", celsius, kelvinResult)
	// 		}

	// 	} else if kelvin != 0 {
	// 		fahrResult := conv.KelvinToFahrenheit(kelvin)
	// 		celsiusResult := conv.KelvinToCelsius(kelvin)
	// 		if out == "C" {
	// 			fmt.Printf("%.2fK is %.2f°C\n", kelvin, celsiusResult)
	// 		} else if out == "F" {
	// 			fmt.Printf("%.2fK is %.2f°F\n", kelvin, fahrResult)
	// 		}
	// 	} else {
	// 		fmt.Println("error: one of the following temperature arguments must be specified: -F, -C, -K")
	// 		return
	// 	}
}

// 	// Her er noen eksempler du kan bruke i den manuelle testingen
// 	//fmt.Println(fahr, out, funfacts)

// 	fmt.Println("len(flag.Args())", len(flag.Args()))
// 	fmt.Println("flag.NFlag()", flag.NFlag())

// 	fmt.Println(isFlagPassed("out"))

// 	// Eksempel på enkel logikk
// 	if out == "C" && isFlagPassed("F") {
// 		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
// 		// skal returnere °C
// 		fmt.Println("0°F er -17.78°C")
// 	}

// }

// // Funksjonen sjekker om flagget er spesifisert på kommandolinje
// // Du trenger ikke å bruke den, men den kan hjelpe med logikken
// func isFlagPassed(name string) bool {
// 	found := false
// 	flag.Visit(func(f *flag.Flag) {
// 		if f.Name == name {
// 			found = true
// 		}
// 	})
// 	return found
//}
