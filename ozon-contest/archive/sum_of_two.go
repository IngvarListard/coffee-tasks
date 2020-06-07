package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const chunkSize = 2

var (
	f                *os.File
	chunk            []byte
	err              error
	count            int
	lastPiece        string
	prevStrEndsSpace bool
)

func main() {
	f, err = os.Open("input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	counts := make(map[int32]uint16)

	expectedSumStr, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal("Error Reading: ", err)
	}

	expectedSum, err := strconv.ParseInt(string(expectedSumStr), 10, 64)
	if err != nil {
		log.Fatal("error parse expected sum", err)
	}

	chunk = make([]byte, chunkSize)

	prevStrEndsSpace = true
	EOF := false
	for {
		count, err = reader.Read(chunk)
		if err == io.EOF {
			EOF = true
		} else if err != nil {
			log.Fatal("error reading file", err)
		}

		numbersStr := strings.TrimSuffix(string(chunk[:count]), "\n")
		spaceLast := strings.HasSuffix(numbersStr, " ")
		spaceFirst := strings.HasPrefix(numbersStr, " ")
		numbersStr = strings.Trim(numbersStr, " ")

		numbersList := strings.Split(numbersStr, " ")

		if !prevStrEndsSpace && !spaceFirst {
			numbersList[0] = lastPiece + numbersList[0]
		}
		if !prevStrEndsSpace && spaceFirst {
			numbersList = append(numbersList, "")
			copy(numbersList[1:], numbersList)
			numbersList[0] = lastPiece
			//numbersList = append(numbersList, lastPiece)
		}
		lastPiece = numbersList[len(numbersList)-1]
		if !spaceLast && !EOF {
			numbersList = numbersList[:len(numbersList)-1]
		}

		prevStrEndsSpace = spaceLast

		for _, v := range numbersList {
			number, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				continue
			}
			if number >= expectedSum {
				continue
			}

			counts[int32(number)]++
		}
		if EOF {
			break
		}
	}
	output, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0660)
	defer output.Close()
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
