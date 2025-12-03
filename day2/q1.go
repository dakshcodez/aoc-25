package main

import (
	"fmt"
	"strings"
	"os"
	"log"
	"strconv"
)

type NumberPair struct {
	Start int
	End int
}

func createGrid(filename string) ([]NumberPair, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	rawPairs := strings.Split(string(content), ",")

	var grid []NumberPair
	
	for _, item := range rawPairs {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		parts := strings.Split(item, "-")
		if len(parts) != 2 {
			continue
		}

		val1, err1 := strconv.Atoi(parts[0])
		val2, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			continue
		}

		grid = append(grid, NumberPair{Start: val1, End: val2})
	}
	return grid, nil
}

func main() {
	data, err := createGrid("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	// fmt.Println("Successfully created grid:")
	// for _, p := range data {
	// 	fmt.Printf("Range: %d -> %d\n", p.Start, p.End)
	// }
}
