package main

import (
	"github.com/01-edu/z01"

	"os"
)

func main() {
	arg := os.Args
	for _, char := range arg[0] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
