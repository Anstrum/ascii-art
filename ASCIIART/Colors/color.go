package Colors

import "github.com/gookit/color"

func GetColor(myColor string) func(string) {
	var ColorFunc func(string)
	switch myColor {
	case "red":
		ColorFunc = func(val string) {
			color.Red.Print(val)
		}
	case "green":
		ColorFunc = func(val string) {
			color.Green.Print(val)
		}
	case "yellow":
		ColorFunc = func(val string) {
			color.Yellow.Print(val)
		}
	case "blue":
		ColorFunc = func(val string) {
			color.Blue.Print(val)
		}
	case "purple":
		ColorFunc = func(val string) {
			color.FgMagenta.Print(val)
		}
	case "cyan":
		ColorFunc = func(val string) {
			color.Cyan.Print(val)
		}
	case "white":
		ColorFunc = func(val string) {
			color.White.Print(val)
		}
	default:
		ColorFunc = func(val string) {
			print(val)
		}
	}
	return ColorFunc
}
