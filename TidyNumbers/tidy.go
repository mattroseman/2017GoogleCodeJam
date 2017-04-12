package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	numTests, _ := strconv.Atoi(text)

	for i := 0; i < numTests; i++ {
		testCase, _ := reader.ReadString('\n')
		testCase = testCase[:len(testCase)-1]
		// start with the second to last digit and go to the 0th character
		for j := len(testCase) - 2; j >= 0; j-- {
			// if this digit is greater than the previous one
			currentDigit, _ := strconv.Atoi(string(testCase[j]))
			nextDigit, _ := strconv.Atoi(string(testCase[j+1]))
			if currentDigit > nextDigit {
				if j == 0 {
					testCase = strconv.Itoa(currentDigit-1) + strings.Repeat("9", len(testCase)-(j+1))
				} else {
					testCase = string(testCase[:j]) + strconv.Itoa(currentDigit-1) + strings.Repeat("9", len(testCase)-(j+1))
				}
			}
		}
		fmt.Printf("Case #%d: %s\n", i+1, strings.TrimLeft(testCase, "0"))
	}
}
