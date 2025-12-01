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

		if strings.ToUpper(string(rot)) == "L" {
			r := (pos - val) % 100
			if r < 0 {
				pos = 100 + r
			} else {
				pos = r
			}
		} else {
			pos = (pos + val) % 100
		}

		if pos == 0 {
			res++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
