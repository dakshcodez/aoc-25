package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

func trimRange(start, end string) (int, int) {
	var newStart int
	var newEnd int

	if len(start) % 2 != 0 {
		newStart = int(math.Pow(10, float64(len(start))))
	} else {
		newStart, _ = strconv.Atoi(start)
	}

	if len(end) % 2 != 0 {
		newEnd = int(math.Pow(10, float64(len(end) - 1))) - 1
	} else {
		newEnd, _ = strconv.Atoi(end)
	}

	return newStart, newEnd
}

func sumOfInvalid(start, end int) int {
	//var sum int

	str1 := strconv.Itoa(start)
	str2 := strconv.Itoa(end)

	if len(str1) % 2 != 0 && len(str2) % 2 != 0 && len(str1) == len(str2) {
		return 0
	}

	start, end = trimRange(str1, str2)
	
	fmt.Println(start, end)
	return 0

}

func main() {
	data, err := createGrid("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	var sum int

	fmt.Println("Successfully created grid:")
	for _, p := range data {
		//fmt.Printf("Range: %d -> %d\n", p.Start, p.End)
		sum += sumOfInvalid(p.Start, p.End)
	}
}
