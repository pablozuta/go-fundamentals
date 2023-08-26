package main

import "fmt"

type ZodiacSign struct {
	Name          string
	Compatibility []string
}

var zodiacSignsCompatibility = map[string]ZodiacSign{
	"Aries": {
		Name:          "Aries",
		Compatibility: []string{"Leo", "Sagittarius", "Gemini"},
	},
	"Taurus": {
		Name:          "Taurus",
		Compatibility: []string{"Virgo", "Capricorn", "Pisces"},
	},
	"Gemini": {
		Name:          "Gemini",
		Compatibility: []string{"Libra", "Aquarius", "Aries"},
	},
	"Cancer": {
		Name:          "Cancer",
		Compatibility: []string{"Scorpio", "Pisces", "Taurus"},
	},
	"Leo": {
		Name:          "Leo",
		Compatibility: []string{"Aries", "Sagittarius", "Gemini"},
	},
	"Virgo": {
		Name:          "Virgo",
		Compatibility: []string{"Taurus", "Capricorn", "Cancer"},
	},
	"Libra": {
		Name:          "Libra",
		Compatibility: []string{"Gemini", "Aquarius", "Leo"},
	},
	"Scorpio": {
		Name:          "Scorpio",
		Compatibility: []string{"Cancer", "Pisces", "Virgo"},
	},
	"Sagittarius": {
		Name:          "Sagittarius",
		Compatibility: []string{"Leo", "Aries", "Libra"},
	},
	"Capricorn": {
		Name:          "Capricorn",
		Compatibility: []string{"Taurus", "Virgo", "Scorpio"},
	},
	"Aquarius": {
		Name:          "Aquarius",
		Compatibility: []string{"Gemini", "Libra", "Sagittarius"},
	},
	"Pisces": {
		Name:          "Pisces",
		Compatibility: []string{"Cancer", "Scorpio", "Taurus"},
	},
}

func checkCompatibility(sign1, sign2 string) bool {
	compatibility := zodiacSignsCompatibility[sign1].Compatibility

	for _, sign := range compatibility {
		if sign == sign2 {
			return true
		}
	}
	return false
}

func main() {
	sign1 := "Aries"
	sign2 := "Leo"
	compatible := checkCompatibility(sign1, sign2)
	if compatible {
		fmt.Printf("%s and %s are compatible!\n", sign1, sign2)
	} else {
		fmt.Printf("%s and %s are not compatible.\n", sign1, sign2)
	}

}
