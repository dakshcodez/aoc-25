package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func q1(rot byte, val int, pos *int, res *int) {
	if strings.ToUpper(string(rot)) == "L" {
			r := (*pos - val) % 100
			if r < 0 {
				*pos = 100 + r
			} else {
				*pos = r
			}
		} else {
			*pos = (*pos + val) % 100
		}
	if *pos == 0 {
		*res ++
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	pos := 50
	res := 0

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
		
		q1(rot, val, &pos, &res)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
