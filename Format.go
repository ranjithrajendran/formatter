package formatter

import (
	"fmt"
	"math"
	"strings"
)

// ReadableIndianCash - code convert normal float amount to readable coma seperated format
func ReadableIndianCash(num float64) string {

	// seperate decimal and integer part
	main, fraction := math.Modf(num)

	// Round offing the fraction to 2 decimal ponit
	fractionString := fmt.Sprintf("%.2f", fraction)[1:]

	// convert float to string
	mainString := fmt.Sprintf("%d", int(main))

	var parts []string

	switch {

	case len(mainString) > 12:
		parts = append(parts, mainString[:len(mainString)-12])
		parts = append(parts, mainString[len(mainString)-12:len(mainString)-10])
		parts = append(parts, mainString[len(mainString)-10:len(mainString)-7])
		parts = append(parts, mainString[len(mainString)-7:len(mainString)-5])
		parts = append(parts, mainString[len(mainString)-5:len(mainString)-3])
		parts = append(parts, mainString[len(mainString)-3:])

	case len(mainString) > 10:
		parts = append(parts, mainString[:len(mainString)-10])
		parts = append(parts, mainString[len(mainString)-10:len(mainString)-7])
		parts = append(parts, mainString[len(mainString)-7:len(mainString)-5])
		parts = append(parts, mainString[len(mainString)-5:len(mainString)-3])
		parts = append(parts, mainString[len(mainString)-3:])

	case len(mainString) > 7:
		parts = append(parts, mainString[:len(mainString)-7])
		parts = append(parts, mainString[len(mainString)-7:len(mainString)-5])
		parts = append(parts, mainString[len(mainString)-5:len(mainString)-3])
		parts = append(parts, mainString[len(mainString)-3:])

	case len(mainString) > 5:
		parts = append(parts, mainString[:len(mainString)-5])
		parts = append(parts, mainString[len(mainString)-5:len(mainString)-3])
		parts = append(parts, mainString[len(mainString)-3:])

	case len(mainString) > 3:
		parts = append(parts, mainString[:len(mainString)-3])
		parts = append(parts, mainString[len(mainString)-3:])

	default:
		parts = append(parts, mainString)

	}
	return strings.Join(parts, ",") + fractionString
}
