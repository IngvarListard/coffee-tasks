package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[int32]uint16)

	lines := strings.Split(string(content), "\n")
	expectedSumStr, numbersStr := lines[0], lines[1]

	splitted := strings.Split(numbersStr, " ")

	expectedSum, err := strconv.ParseInt(expectedSumStr, 10, 64)

	output, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0660)

	for _, v := range splitted {
		number, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		if number >= expectedSum {
			continue
		}

		counts[int32(number)]++
	}

	expectedSumInt := int32(expectedSum)
	for k := range counts {
		desiredTerm := expectedSumInt - k
		if entries, ok := counts[desiredTerm]; ok {
			if desiredTerm == k && entries >= 2 && desiredTerm+k == expectedSumInt {
				output.WriteString("1")
				return
			} else if desiredTerm+k == expectedSumInt && desiredTerm != k {
				output.WriteString("1")
				return
			}

		}
	}

	output.WriteString("0")
}
