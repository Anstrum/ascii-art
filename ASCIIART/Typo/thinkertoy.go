package Typo

func GetThinkerToy(char string) [][]string {
	var toReturn [][]string
	space := [][]string{
		{"    "},
		{"    "},
		{"    "},
		{"    "},
		{"    "},
		{"    "},
		{"    "},
		{"    "},
	}
	exPoint := [][]string{
		{"  "},
		{"  "},
		{"o "},
		{"| "},
		{"o "},
		{"  "},
		{"O "},
		{"  "},
	}
	doubleQuote := [][]string{
		{"o o "},
		{"| | "},
		{"    "},
		{"    "},
		{"    "},
		{"    "},
		{"    "},
		{"    "},
	}
	diez := [][]string{
		{"      "},
		{" | |  "},
		{"-O-O- "},
		{" | |  "},
		{"-O-O- "},
		{" | |  "},
		{"      "},
		{"      "},
	}
	dollard := [][]string{
		{"  | |   "},
		{" -O-O-  "},
		{"o | |   "},
		{" -O-O-  "},
		{"  | | o "},
		{" -O-O-  "},
		{"  | |   "},
		{"        "},
	}
	pourcentage := [][]string{
		{"      "},
		{"    O "},
		{"o  /  "},
		{"  /   "},
		{" /  o "},
		{"O     "},
		{"      "},
		{"      "},
	}
	and := [][]string{
		{"     "},
		{"     "},
		{"  o  "},
		{" /|  "},
		{"o-O- "},
		{"  |  "},
		{"     "},
		{"     "},
	}
	simpleQuote := [][]string{
		{"o "},
		{"| "},
		{"  "},
		{"  "},
		{"  "},
		{"  "},
		{"  "},
		{"  "},
	}
	accolLeft := [][]string{
		{"   "},
		{" / "},
		{"o  "},
		{"|  "},
		{"o  "},
		{" \\ "},
		{"   "},
		{"   "},
	}
	accolRight := [][]string{
		{"   "},
		{"\\  "},
		{" o "},
		{" | "},
		{" o "},
		{"/  "},
		{"   "},
		{"   "},
	}
	asterisque := [][]string{
		{"      "},
		{"o | o "},
		{" \\|/  "},
		{"--O-- "},
		{" /|\\  "},
		{"o | o "},
		{"      "},
		{"      "},
	}
	plus := [][]string{
		{"    "},
		{"    "},
		{" |  "},
		{"-o- "},
		{" |  "},
		{"    "},
		{"    "},
		{"    "},
	}
	virgule := [][]string{
		{"  "},
		{"  "},
		{"  "},
		{"  "},
		{"  "},
		{"o "},
		{"| "},
		{"  "},
	}
	tiret := [][]string{
		{"    "},
		{"    "},
		{"    "},
		{"    "},
		{"o-o "},
		{"    "},
		{"    "},
		{"    "},
	}
	point := [][]string{
		{"  "},
		{"  "},
		{"  "},
		{"  "},
		{"  "},
		{"O "},
		{"  "},
		{"  "},
	}
	slash := [][]string{
		{"      "},
		{"    o "},
		{"   /  "},
		{"  o   "},
		{" /    "},
		{"o     "},
		{"      "},
		{"      "},
	}
	zero := [][]string{
		{"      "},
		{" o-o  "},
		{"o  /o "},
		{"| / | "},
		{"o/  o "},
		{" o-o  "},
		{"      "},
		{"      "},
	}
	un := [][]string{
		{"      "},
		{"  0   "},
		{" /|   "},
		{"o |   "},
		{"  |   "},
		{"o-o-o "},
		{"      "},
		{"      "},
	}
	deux := [][]string{
		{"     "},
		{" --  "},
		{"o  o "},
		{"  /  "},
		{" /   "},
		{"o--o "},
		{"     "},
		{"     "},
	}
	trois := [][]string{
		{"     "},
		{"o-o  "},
		{"   | "},
		{" oo  "},
		{"   | "},
		{"o-o  "},
		{"     "},
		{"     "},
	}
	quatre := [][]string{
		{"     "},
		{"o  o "},
		{"|  | "},
		{"o--O "},
		{"   | "},
		{"   o "},
		{"     "},
		{"     "},
	}
	cinq := [][]string{
		{"     "},
		{"o--o "},
		{"|    "},
		{"o-o  "},
		{"   | "},
		{"o-o  "},
		{"     "},
		{"     "},
	}
	six := [][]string{
		{"      "},
		{"  o   "},
		{" /    "},
		{"O--o  "},
		{"o   | "},
		{" o-o  "},
		{"      "},
		{"      "},
	}
	sept := [][]string{
		{"      "},
		{"o---o "},
		{"   /  "},
		{"  o   "},
		{"  |   "},
		{"  o   "},
		{"      "},
		{"      "},
	}
	huit := [][]string{
		{"      "},
		{" o-o  "},
		{"|   | "},
		{" o-o  "},
		{"|   | "},
		{" o-o  "},
		{"      "},
		{"      "},
	}
	neuf := [][]string{
		{"      "},
		{" o-o  "},
		{"|   o "},
		{" o--O "},
		{"   /  "},
		{"  o   "},
		{"      "},
		{"      "},
	}
	doublePoint := [][]string{
		{"  "},
		{"  "},
		{"O "},
		{"  "},
		{"O "},
		{"  "},
		{"  "},
		{"  "},
	}
	pointVirgule := [][]string{
		{"  "},
		{"  "},
		{"o "},
		{"  "},
		{"o "},
		{"| "},
		{"  "},
		{"  "},
	}
	less := [][]string{
		{"    "},
		{"  o "},
		{" /  "},
		{"O   "},
		{" \\  "},
		{"  o "},
		{"    "},
		{"    "},
	}
	equal := [][]string{
		{"     "},
		{"     "},
		{"     "},
		{"o--o "},
		{"o--o "},
		{"     "},
		{"     "},
		{"     "},
	}
	more := [][]string{
		{"    "},
		{"o   "},
		{" \\  "},
		{"  O "},
		{" /  "},
		{"o   "},
		{"    "},
		{"    "},
	}
	questPoint := [][]string{
		{"      "},
		{" o-o  "},
		{"o   o "},
		{"   /  "},
		{"  o   "},
		{"      "},
		{"  O   "},
		{"      "},
	}
	arobase := [][]string{
		{"      "},
		{"  o   "},
		{" / \\  "},
		{"o O-o "},
		{" \\    "},
		{"  o-  "},
		{"      "},
		{"      "},
	}
	A := [][]string{
		{"      "},
		{"  O   "},
		{" / \\  "},
		{"o---o "},
		{"|   | "},
		{"o   o "},
		{"      "},
		{"      "},
	}
	B := [][]string{
		{"      "},
		{"o--o  "},
		{"|   | "},
		{"O--o  "},
		{"|   | "},
		{"o--o  "},
		{"      "},
		{"      "},
	}
	C := [][]string{
		{"      "},
		{"  o-o "},
		{" /    "},
		{"O     "},
		{" \\    "},
		{"  o-o "},
		{"      "},
		{"      "},
	}
	D := [][]string{
		{"      "},
		{"o-o   "},
		{"|  \\  "},
		{"|   O "},
		{"|  /  "},
		{"o-o   "},
		{"      "},
		{"      "},
	}
	E := [][]string{
		{"     "},
		{"o--o "},
		{"|    "},
		{"O-o  "},
		{"|    "},
		{"o--o "},
		{"     "},
		{"     "},
	}
	F := [][]string{
		{"     "},
		{"o--o "},
		{"|    "},
		{"O-o  "},
		{"|    "},
		{"o    "},
		{"     "},
		{"     "},
	}
	G := [][]string{
		{"      "},
		{" o-o  "},
		{"o     "},
		{"|  -o "},
		{"o   | "},
		{" o-o  "},
		{"      "},
		{"      "},
	}
	H := [][]string{
		{"     "},
		{"o  o "},
		{"|  | "},
		{"o  o "},
		{"|  | "},
		{"o  o "},
		{"     "},
		{"     "},
	}
	I := [][]string{
		{"      "},
		{"o-O-o "},
		{"  |   "},
		{"  |   "},
		{"  |   "},
		{"o-O-o "},
		{"      "},
		{"      "},
	}
	J := [][]string{
		{"      "},
		{"    o "},
		{"    | "},
		{"    | "},
		{"\\   o "},
		{" o-o  "},
		{"      "},
		{"      "},
	}
	K := [][]string{
		{"     "},
		{"o  o "},
		{"| /  "},
		{"OO   "},
		{"| \\  "},
		{"o  o "},
		{"     "},
		{"     "},
	}
	L := [][]string{
		{"      "},
		{"o     "},
		{"|     "},
		{"|     "},
		{"|     "},
		{"O---o "},
		{"      "},
		{"      "},
	}
	M := [][]string{
		{"      "},
		{"o   o "},
		{"|\\ /| "},
		{"| O | "},
		{"|   | "},
		{"o   o "},
		{"      "},
		{"      "},
	}
	N := [][]string{
		{"      "},
		{"o   o "},
		{"|\\  | "},
		{"| \\ | "},
		{"|  \\| "},
		{"o   o "},
		{"      "},
		{"      "},
	}
	O := [][]string{
		{"      "},
		{" o-o  "},
		{"o   o "},
		{"|   | "},
		{"o   o "},
		{" o-o  "},
		{"      "},
		{"      "},
	}
	P := [][]string{
		{"      "},
		{"o--o  "},
		{"|   | "},
		{"O--o  "},
		{"|     "},
		{"o     "},
		{"      "},
		{"      "},
	}
	Q := [][]string{
		{"      "},
		{" o-o  "},
		{"o   o "},
		{"|   | "},
		{"o   O "},
		{" o-O\\ "},
		{"      "},
		{"      "},
	}
	R := [][]string{
		{"      "},
		{"o--o  "},
		{"|   | "},
		{"O-Oo  "},
		{"|  \\  "},
		{"o   o "},
		{"      "},
		{"      "},
	}
	S := [][]string{
		{"      "},
		{" o-o  "},
		{"|     "},
		{" o-o  "},
		{"    | "},
		{"o--o  "},
		{"      "},
		{"      "},
	}
	T := [][]string{
		{"      "},
		{"o-O-o "},
		{"  |   "},
		{"  |   "},
		{"  |   "},
		{"  o   "},
		{"      "},
		{"      "},
	}
	U := [][]string{
		{"      "},
		{"o   o "},
		{"|   | "},
		{"|   | "},
		{"|   | "},
		{" o-o  "},
		{"      "},
		{"      "},
	}
	V := [][]string{
		{"      "},
		{"o   o "},
		{"|   | "},
		{"o   o "},
		{" \\ /  "},
		{"  o   "},
		{"      "},
		{"      "},
	}
	W := [][]string{
		{"          "},
		{"o       o "},
		{"|       | "},
		{"o   o   o "},
		{" \\ / \\ /  "},
		{"  o   o   "},
		{"          "},
		{"          "},
	}
	X := [][]string{
		{"      "},
		{"o   o "},
		{" \\ /  "},
		{"  O   "},
		{" / \\  "},
		{"o   o "},
		{"      "},
		{"      "},
	}
	Y := [][]string{
		{"      "},
		{"o   o "},
		{" \\ /  "},
		{"  O   "},
		{"  |   "},
		{"  o   "},
		{"      "},
		{"      "},
	}
	Z := [][]string{
		{"      "},
		{"o---o "},
		{"   /  "},
		{" -O-  "},
		{" /    "},
		{"o---o "},
		{"      "},
		{"      "},
	}
	rightBraquet := [][]string{
		{"    "},
		{"O-o "},
		{"|   "},
		{"|   "},
		{"|   "},
		{"|   "},
		{"O-o "},
		{"    "},
	}
	counterSlash := [][]string{
		{"      "},
		{"o     "},
		{" \\    "},
		{"  o   "},
		{"   \\  "},
		{"    o "},
		{"      "},
		{"      "},
	}
	leftBraquet := [][]string{
		{"    "},
		{"o-O "},
		{"  | "},
		{"  | "},
		{"  | "},
		{"  | "},
		{"o-O "},
		{"    "},
	}
	caret := [][]string{
		{"    "},
		{" o  "},
		{"/ \\ "},
		{"    "},
		{"    "},
		{"    "},
		{"    "},
		{"    "},
	}
	underscore := [][]string{
		{"      "},
		{"      "},
		{"      "},
		{"      "},
		{"      "},
		{"o---o "},
		{"      "},
		{"      "},
	}
	accentGrave := [][]string{
		{"  "},
		{"0 "},
		{"| "},
		{"  "},
		{"  "},
		{"  "},
		{"  "},
		{"  "},
	}
	a := [][]string{
		{"     "},
		{"     "},
		{"     "},
		{" oo  "},
		{"| |  "},
		{"o-o- "},
		{"     "},
		{"     "},
	}
	b := [][]string{
		{"     "},
		{"o    "},
		{"|    "},
		{"O-o  "},
		{"|  | "},
		{"o-o  "},
		{"     "},
		{"     "},
	}
	c := [][]string{
		{"     "},
		{"     "},
		{"     "},
		{" o-o "},
		{"|    "},
		{" o-o "},
		{"     "},
		{"     "},
	}
	d := [][]string{
		{"     "},
		{"   o "},
		{"   | "},
		{" o-O "},
		{"|  | "},
		{" o-o "},
		{"     "},
		{"     "},
	}
	e := [][]string{
		{"    "},
		{"    "},
		{"    "},
		{"o-o "},
		{"|-' "},
		{"o-o "},
		{"    "},
		{"    "},
	}
	f := [][]string{
		{"     "},
		{" o-o "},
		{" |   "},
		{"-O-  "},
		{" |   "},
		{" o   "},
		{"     "},
		{"     "},
	}
	g := [][]string{
		{"     "},
		{"     "},
		{"     "},
		{"o--o "},
		{"|  | "},
		{"o--O "},
		{"   | "},
		{"o--o "},
	}
	h := [][]string{
		{"     "},
		{"o    "},
		{"|    "},
		{"O--o "},
		{"|  | "},
		{"o  o "},
		{"     "},
		{"     "},
	}
	i := [][]string{
		{"  "},
		{"  "},
		{"o "},
		{"  "},
		{"| "},
		{"| "},
		{"  "},
		{"  "},
	}
	j := [][]string{
		{"      "},
		{"      "},
		{"    o "},
		{"      "},
		{"    o "},
		{"    | "},
		{"o   o "},
		{" o-o  "},
	}
	k := [][]string{
		{"     "},
		{"o    "},
		{"| /  "},
		{"OO   "},
		{"| \\  "},
		{"o  o "},
		{"     "},
		{"     "},
	}
	l := [][]string{
		{"  "},
		{"o "},
		{"| "},
		{"| "},
		{"| "},
		{"o "},
		{"  "},
		{"  "},
	}
	m := [][]string{
		{"      "},
		{"      "},
		{"      "},
		{"o-O-o "},
		{"| | | "},
		{"o o o "},
		{"      "},
		{"      "},
	}
	n := [][]string{
		{"     "},
		{"     "},
		{"     "},
		{"o-o  "},
		{"|  | "},
		{"o  o "},
		{"     "},
		{"     "},
	}
	o := [][]string{
		{"    "},
		{"    "},
		{"    "},
		{"o-o "},
		{"| | "},
		{"o-o "},
		{"    "},
		{"    "},
	}
	p := [][]string{
		{"     "},
		{"     "},
		{"     "},
		{"o-o  "},
		{"|  | "},
		{"O-o  "},
		{"|    "},
		{"o    "},
	}
	q := [][]string{
		{"     "},
		{"     "},
		{"     "},
		{" o-o "},
		{"|  | "},
		{" o-O "},
		{"   | "},
		{"   o "},
	}
	r := [][]string{
		{"    "},
		{"    "},
		{"    "},
		{"o-o "},
		{"|   "},
		{"o   "},
		{"    "},
		{"    "},
	}
	s := [][]string{
		{"    "},
		{"    "},
		{"    "},
		{"o-o "},
		{" \\  "},
		{"o-o "},
		{"    "},
		{"    "},
	}
	t := [][]string{
		{"    "},
		{" o  "},
		{" |  "},
		{"-o- "},
		{" |  "},
		{" o  "},
		{"    "},
		{"    "},
	}
	u := [][]string{
		{"     "},
		{"     "},
		{"     "},
		{"o  o "},
		{"|  | "},
		{"o--o "},
		{"     "},
		{"     "},
	}
	v := [][]string{
		{"      "},
		{"      "},
		{"      "},
		{"o  o "},
		{" \\ /  "},
		{"  o   "},
		{"      "},
		{"      "},
	}
	w := [][]string{
		{"          "},
		{"          "},
		{"          "},
		{"o   o   o "},
		{" \\ / \\ /  "},
		{"  o   o   "},
		{"          "},
		{"          "},
	}
	x := [][]string{
		{"    "},
		{"    "},
		{"    "},
		{"\\ / "},
		{" o  "},
		{"/ \\ "},
		{"    "},
		{"    "},
	}
	y := [][]string{
		{"     "},
		{"     "},
		{"     "},
		{"o  o "},
		{"|  | "},
		{"o--O "},
		{"   | "},
		{"o--o "},
	}
	z := [][]string{
		{"    "},
		{"    "},
		{"    "},
		{"o-o "},
		{" /  "},
		{"o-o "},
		{"    "},
		{"    "},
	}
	rightBrace := [][]string{
		{"      "},
		{"  o-o "},
		{"  |   "},
		{"o-O   "},
		{"  |   "},
		{"  o-o "},
		{"      "},
		{"      "},
	}
	verticalBar := [][]string{
		{"  "},
		{"o "},
		{"| "},
		{"o "},
		{"| "},
		{"o "},
		{"  "},
		{"  "},
	}
	leftBrace := [][]string{
		{"      "},
		{"o-o   "},
		{"  |   "},
		{"  O-o "},
		{"  |   "},
		{"o-o   "},
		{"      "},
		{"      "},
	}
	tilde := [][]string{
		{" o_ / "},
		{"/  o  "},
		{"      "},
		{"      "},
		{"      "},
		{"      "},
		{"      "},
		{"      "},
	}

	switch char {
	case " ":
		toReturn = space
	case "!":
		toReturn = exPoint
	case "\"":
		toReturn = doubleQuote
	case "#":
		toReturn = diez
	case "$":
		toReturn = dollard
	case "%":
		toReturn = pourcentage
	case "&":
		toReturn = and
	case "'":
		toReturn = simpleQuote
	case "(":
		return accolLeft
	case ")":
		return accolRight
	case "*":
		toReturn = asterisque
	case "+":
		toReturn = plus
	case ",":
		toReturn = virgule
	case "-":
		toReturn = tiret
	case ".":
		toReturn = point
	case "/":
		toReturn = slash
	case "0":
		toReturn = zero
	case "1":
		toReturn = un
	case "2":
		toReturn = deux
	case "3":
		toReturn = trois
	case "4":
		toReturn = quatre
	case "5":
		toReturn = cinq
	case "6":
		toReturn = six
	case "7":
		toReturn = sept
	case "8":
		toReturn = huit
	case "9":
		toReturn = neuf
	case ":":
		toReturn = doublePoint
	case ";":
		toReturn = pointVirgule
	case "<":
		toReturn = less
	case "=":
		toReturn = equal
	case ">":
		toReturn = more
	case "?":
		toReturn = questPoint
	case "@":
		toReturn = arobase
	case "A":
		toReturn = A
	case "B":
		toReturn = B
	case "C":
		toReturn = C
	case "D":
		toReturn = D
	case "E":
		toReturn = E
	case "F":
		toReturn = F
	case "G":
		toReturn = G
	case "H":
		toReturn = H
	case "I":
		toReturn = I
	case "J":
		toReturn = J
	case "K":
		toReturn = K
	case "L":
		toReturn = L
	case "M":
		toReturn = M
	case "N":
		toReturn = N
	case "O":
		toReturn = O
	case "P":
		toReturn = P
	case "Q":
		toReturn = Q
	case "R":
		toReturn = R
	case "S":
		toReturn = S
	case "T":
		toReturn = T
	case "U":
		toReturn = U
	case "V":
		toReturn = V
	case "W":
		toReturn = W
	case "X":
		toReturn = X
	case "Y":
		toReturn = Y
	case "Z":
		toReturn = Z
	case "[":
		toReturn = rightBraquet
	case "\\":
		toReturn = counterSlash
	case "]":
		toReturn = leftBraquet
	case "^":
		toReturn = caret
	case "_":
		toReturn = underscore
	case "`":
		toReturn = accentGrave
	case "a":
		toReturn = a
	case "b":
		toReturn = b
	case "c":
		toReturn = c
	case "d":
		toReturn = d
	case "e":
		toReturn = e
	case "f":
		toReturn = f
	case "g":
		toReturn = g
	case "h":
		toReturn = h
	case "i":
		toReturn = i
	case "j":
		toReturn = j
	case "k":
		toReturn = k
	case "l":
		toReturn = l
	case "m":
		toReturn = m
	case "n":
		toReturn = n
	case "o":
		toReturn = o
	case "p":
		toReturn = p
	case "q":
		toReturn = q
	case "r":
		toReturn = r
	case "s":
		toReturn = s
	case "t":
		toReturn = t
	case "u":
		toReturn = u
	case "v":
		toReturn = v
	case "w":
		toReturn = w
	case "x":
		toReturn = x
	case "y":
		toReturn = y
	case "z":
		toReturn = z
	case "{":
		toReturn = rightBrace
	case "|":
		toReturn = verticalBar
	case "}":
		toReturn = leftBrace
	case "~":
		toReturn = tilde
	default:
		toReturn = space
	}
	return toReturn
}
