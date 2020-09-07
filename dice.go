// Package dice contains utilities for dealing
// with dungeons and dragons dice
package dice

import (
	"fmt"
	"regexp"
)

// Die holds the number of sides on the die
type Die int

// ParseDice takes a Dungeons and Dragons dice string
// and returns a die type and number of dice
func ParseDice(diceString string) (die Die, number int, err error) {
	regex, err := regexp.Compile(`(?P<nDice>\d*)d(?P<nSides>\d+)`)
	if err != nil {
		return Die(-1), -1, err
	}

	matches := regex.FindStringSubmatch(diceString)
	names := regex.SubexpNames()
	for i := range matches {
		if i != 0 {
			fmt.Println(names[i], matches[i])
		}
	}
	return Die(-1), -1, err
}
