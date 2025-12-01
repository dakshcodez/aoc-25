package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0
	pos := 50

	for scanner.Scan() {
		text := scanner.Text()

		if strings.TrimSpace(text) == "" {
			continue
		}

		rot := text[0]
		numStr := text[1:]

		val, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("error converting number")
			return
		}
		
		res += val / 100
		val = val % 100
		
		if strings.ToUpper(string(rot)) == "L" {
			newPos := pos - val
			if newPos >= 0 {
				pos = newPos
				if pos == 0 {
					res++				
				}
			} else {
				if pos != 0 {
					res++						// figuring this drained the life out of me
				}
				pos = 100 + newPos % 100
			}
		} else {
			newPos := pos + val

			if newPos < 100 {
				pos = newPos
			} else {
				pos = newPos % 100
				res ++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
