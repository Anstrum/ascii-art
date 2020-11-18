package Typo

func GetShadow(char string) [][]string {
	var toReturn [][]string
	space := [][]string{
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
	}
	switch char {
	case " ":
		toReturn = space
	}

	return toReturn
}
