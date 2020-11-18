package Typo

func GetChar(typo, val string) [][]string {
	var toReturn [][]string

	switch typo {
	case "standard":
		toReturn = GetStandard(val)
	case "thinkertoy":
		toReturn = GetThinkerToy(val)
	case "shadow":
		toReturn = GetShadow(val)
	default:
		toReturn = GetStandard(val)
	}

	return toReturn
}
