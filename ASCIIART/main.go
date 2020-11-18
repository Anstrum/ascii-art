package main

import (
	"os"

	"./Base"
)

func main() {
	args := os.Args[1:]
	if Base.VerifArgs(args) {
		Base.PrintASCII(args)
	}
}
