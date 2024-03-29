package diceprinter

import (
	"fmt"

	dice "github.com/rafariossaa/appliedgocoursedice"
)

func PrintRoll(sides int, comment string) {
	fmt.Printf("%s: %d\n", comment, dice.Roll(sides))
}
