package Base

func SetParams(args []string, params Params) Params {
	params.text = args[0]

	var parts []string
	if len(args) > 1 {
		for _, val := range args[1:] {
			if Parse(val, &parts) {
				switch parts[0] {
				case "output":
					params.outputFileName = parts[1]
					params.output = true
				case "reverse":
					params.reverseFileName = parts[1]
					params.reverse = true
				case "align":
					params.justifyValue = parts[1]
					params.justify = true
				case "color":
					params.colorValue = parts[1]
				}

			} else {
				switch parts[0] {
				case "standard":
					params.fontStyleValue = parts[0]
				case "shadow":
					params.fontStyleValue = parts[0]
				case "thinkertoy":
					params.fontStyleValue = parts[0]
				}
			}
		}
	}
	return params
}
