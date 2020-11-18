package Justify

import (
	"../Typo"
	"github.com/nathan-fiscaletti/consolesize-go"
)

func Justify(align, text, typo string, Print func(string)) {
	cols, _ := consolesize.GetConsoleSize()

	switch align {
	case "left":
		left(text, typo, cols, Print)
	case "right":
		right(text, typo, cols, Print)
	case "center":
		center(text, typo, cols, Print)
	case "justify":
		justify(text, typo, cols, Print)
	default:
	}
}

func left(s, typo string, cols int, Print func(string)) {
	var toAdd [][]string
	toPrint := [][]string{
		{""},
		{""},
		{""},
		{""},
		{""},
		{""},
		{""},
		{""},
	}

	for _, val := range s {
		toAdd = Typo.GetChar(typo, string(val))
		toPrint = FuseTable(toPrint, toAdd)
	}

	for i := range toPrint { //Print the sentense characters one by one
		for _, val := range toPrint[i] {
			Print(val)
		}
		Print("\n")
	}
}
func right(s, typo string, cols int, Print func(string)) {
	var toAdd [][]string
	toPrint := [][]string{
		{""},
		{""},
		{""},
		{""},
		{""},
		{""},
		{""},
		{""},
	}

	for _, val := range s {
		toAdd = Typo.GetChar(typo, string(val))
		toPrint = FuseTable(toPrint, toAdd)
	}

	for i := range toPrint { //Print the sentense characters one by one
		for _, val := range toPrint[i] {
			Print(val)
		}
		Print("\n")
	}
}
func center(s, typo string, cols int, Print func(string)) {
}
func justify(s, typo string, cols int, Print func(string)) {
}

func FuseTable(table, toAdd [][]string) [][]string { //add letter table to full sentense table
	for i := range table {
		line := toAdd[i]
		for _, val := range line {
			table[i] = append(table[i], val)
		}
	}
	return table
}
