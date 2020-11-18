package Base

func VerifArgs(args []string) bool {
	hasFS := false
	hasColor := false
	hasOutput := false
	hasReverse := false
	hasJustify := false

	if args[0] == "" { //Contiens du texte a imprimer
		return false
	}

	var parts []string
	for _, val := range args[1:] { //Vérifie si les arguments sont valides
		parts = []string{}
		if !(Parse(val, &parts)) {
			if parts[0] == "standard" || parts[0] == "shadow" || parts[0] == "thinkertoy" {
				if hasFS {
					return Error("Argument en doublon")
				}
				hasFS = true
			} else {
				return Error("Argument inconnu: " + parts[0])
			}
		}

		switch parts[0] {
		case "color":
			if TestColorValues(parts[1]) {
				if hasColor {
					return Error("Couleur définie deux fois")
				}
				hasColor = true
			} else {
				return Error("Couleur non valide: " + parts[1])
			}
		case "output":
			if parts[1][len(parts[1])-4:] == ".txt" && len(parts[1]) > 4 {
				if hasOutput {
					return Error("Output défini deux fois")
				}
				hasOutput = true
			} else {
				return Error("Nom de fichier incorrect: " + parts[1])
			}
		case "align":
			if parts[1] == "center" || parts[1] == "left" || parts[1] == "right" || parts[1] == "justify" {
				if hasJustify {
					return Error("Argument en doublon")
				}
				hasJustify = true
			}
		case "reverse":
			if parts[1][len(parts[1])-4:] == ".txt" && len(parts[1]) > 4 {
				if hasReverse {
					return Error("Output défini deux fois")
				}
				hasReverse = true
			} else {
				return Error("Nom de fichier incorrect: " + parts[1])
			}
		}
	}

	if hasFS && hasReverse { //Traitement des conflits des paramètres
		return Error("FS ne marche pas avec Reverse")
	}
	if hasOutput {
		if hasColor || hasReverse || hasJustify {
			return Error("Output ne marche qu'avec FS")
		}
	}

	return true
}

func TestColorValues(s string) bool {
	colors := []string{
		"red",
		"blue",
		"green",
		"yellow",
		"purple",
		"cyan",
		"white",
	}
	for _, val := range colors {
		if s == val {
			return true
		}
	}
	return false
}
