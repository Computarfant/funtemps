package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfacts i en struktur
*/
type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

// GetFunFacts returns an array of fun facts about the specified subject.
func GetFunFacts(about string) []string {
	var facts []string

	switch about {
	case "sun":
		facts = []string{
			"The temperature in the sun's core is ",            //15 million Kelvin
			"The outer layer of the sun has a temperature of ", //5778 Kelvin
		}
	case "luna":
		facts = []string{
			"The temperature on the moon's surface at night is ",       //-183 degrees Celsius
			"The temperature on the moon's surface during the day is ", //106 degrees Celsius
		}
	case "terra":
		facts = []string{
			"The highest temperature ever recorded on Earth's surface is ", //134 degrees Fahrenheit
			"The lowest temperature ever recorded on Earth's surface is ",  //-89.4 degrees Fahrenheit
		}
	}

	return facts
}

//cut from main go func init()
//flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
// Du må selv definere flag-variabelen for -t flagget, som bestemmer
// hvilken temperaturskala skal brukes når funfacts skal vises
