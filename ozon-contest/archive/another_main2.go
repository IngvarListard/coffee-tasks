package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[int]int)

	lines := strings.Split(string(content), "\n")
	expectedSumStr, numbersStr := lines[0], lines[1]

	splitted := strings.Split(numbersStr, " ")

	numbers := make([]int, 0)

	expectedSum, err := strconv.ParseInt(expectedSumStr, 10, 64)
	for _, v := range splitted {
		number, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatalf("error when converting number: %v", err)
		}
		if number >= expectedSum {
			continue
		}
		numbers = append(numbers, int(number))
		counts[int(number)]++
	}

	sort.Ints(numbers)
	if err != nil {
		log.Fatalf("error when converting number: %v", err)
	}

	first, last := numbers[len(numbers)-1], numbers[0]
	fmt.Println(first, last)
	output, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0660)
	if first+last == int(expectedSum) {
		output.WriteString("1")
		return
	}

	output.WriteString("0")
}
