package Base

func Parse(s string, parts *[]string) bool {
	var temp string = ""
	var arrayTemp []string

	if s[0:2] == "--" {
		for i, val := range s[2:] {
			if string(val) == "=" {
				arrayTemp = append(arrayTemp, temp)
				temp = s[i+3:]
				break
			} else {
				temp += string(val)
			}
		}
		if temp != "" {
			arrayTemp = append(arrayTemp, temp)
			*parts = arrayTemp
			return true
		}
		println("Aucune valeurs après l'argument: " + arrayTemp[0])
	}
	*parts = append(*parts, s)
	return false
}

func Error(errorText string) bool {
	print(errorText)
	return false
}
