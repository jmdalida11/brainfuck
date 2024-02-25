package main

import (
	"fmt"
	"os"
)

const MAX_CELL = 65536 // 2 ** 6

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Please Provide Brain fuck source file")
		os.Exit(1)
	}

	filename := args[1]
	code, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("File not found:", filename)
	}

	var cells = [MAX_CELL]byte{}
	var ptr = 0
	var idx = 0

	openBracket := []int{}

	for idx < len(code) {
		switch rune(code[idx]) {
		case '+':
			cells[ptr]++
		case '-':
			cells[ptr]--
		case '>':
			ptr++
			if ptr == MAX_CELL {
				ptr = 0
			}
		case '<':
			ptr--
			if ptr < 0 {
				ptr = MAX_CELL - 1
			}
		case '[':
			if cells[ptr] != 0 {
				n := len(openBracket)
				if n == 0 || idx != openBracket[n-1] {
					openBracket = append(openBracket, idx)
				}
			} else {
				cnt := 1
				idx++
				for idx < len(code) {
					if rune(code[idx]) == '[' {
						cnt++
					} else if rune(code[idx]) == ']' {
						cnt--
					}
					if cnt == 0 {
						break
					}
					idx++
				}
			}
		case ']':
			if cells[ptr] != 0 {
				idx = openBracket[len(openBracket)-1]
				continue
			} else {
				openBracket = openBracket[:len(openBracket)-1]
			}
		case '.':
			fmt.Print(string(cells[ptr]))
		case ',':
			var input string
			fmt.Scanln(&input)
			cells[ptr] = input[0]
		}
		idx++
	}
}
