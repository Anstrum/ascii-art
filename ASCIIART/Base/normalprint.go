package Base

import "../Typo"

func NormalPrint(params Params, PrintFunc func(string)) {
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
	for _, val := range params.text {
		toAdd = Typo.GetChar(params.fontStyleValue, string(val))
		toPrint = FuseTable(toPrint, toAdd)
	}
	for i := range toPrint { //Print the sentense characters one by one
		for _, val := range toPrint[i] {
			PrintFunc(val)
		}
		PrintFunc("\n")
	}
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
