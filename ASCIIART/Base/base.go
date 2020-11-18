package Base

import (
	"../Colors"
	"../Justify"
	"../Output"
)

type Params struct {
	useColor  bool
	output    bool
	reverse   bool
	justify   bool
	fontStyle bool

	colorValue      string
	outputFileName  string
	justifyValue    string
	fontStyleValue  string
	reverseFileName string

	text string
}

func PrintASCII(args []string) {
	params := Params{}
	params = SetParams(args, params)

	PrintFunc := Colors.GetColor(params.colorValue)
	if params.output {
		Output.Output(params.outputFileName, params.text, params.fontStyleValue)
		return
	}
	if params.reverse {
		return
	}
	if params.justify {
		Justify.Justify(params.justifyValue, params.text, params.fontStyleValue, PrintFunc)
		return
	}
	NormalPrint(params, PrintFunc)
}
