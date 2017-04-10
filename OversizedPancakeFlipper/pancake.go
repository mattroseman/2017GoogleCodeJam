package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Pancakes is a string consisting of only - and + characters
type Pancakes string

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	numTests, _ := strconv.Atoi(text)

	for i := 0; i < numTests; i++ {
		testCase, _ := reader.ReadString('\n')
		testCase = testCase[:len(testCase)-1]
		pancakes, k := parseTest(testCase)
		fmt.Printf("Case #%d: %s\n", i+1, solveTest(pancakes, k))
	}
}

func parseTest(testCase string) (Pancakes, int) {
	test := strings.Split(testCase, " ")
	pancakes := Pancakes(test[0])
	k, _ := strconv.Atoi(test[1])
	return pancakes, k
}

// solveTest returns the number of tries needed to solve the test, or the string IMPOSSIBLE
func solveTest(pancakes Pancakes, k int) string {
	// if pancakes is alread all pluses
	if match, _ := regexp.MatchString("^\\++$", string(pancakes)); match {
		return "0"
	}
	// if size is less than k no flips can happen
	if len(pancakes) < k {
		return "IMPOSSIBLE"
	}
	// find the first minus and flip there
	// if there isn't enough room return IMPOSSIBLE
	// find the first minus index
	re, _ := regexp.Compile("-")
	// the maximum number of flips to make
	for j := 0; j < len(pancakes)-k+1; j++ {
		// if there is a - present in pancakes
		if i := re.FindStringIndex(string(pancakes)); i != nil {
			if i[0] > len(pancakes)-k {
				return "IMPOSSIBLE"
			}
			pancakes.flip(k, i[0])
		} else {
			return strconv.Itoa(j)
		}
	}

	// if there are still minuses after max flips are made
	if i := re.FindStringIndex(string(pancakes)); i != nil {
		return "IMPOSSIBLE"
	}

	// else there are no more minuses
	return strconv.Itoa(len(pancakes) - k + 1)
}

// flip takes in the current pancake state, the flipper size, and the left hand index of the flipper
// i must be between [0, pancakes.length - k]
func (pancakes *Pancakes) flip(k int, i int) {
	if i < 0 || i > len(*pancakes)-k {
		return
	}

	for j := i; j < i+k; j++ {
		// TODO figure out how to properly check values and assign to string pointer
		if (*pancakes)[j] == '+' {
			*pancakes = replaceAtIndex(*pancakes, '-', j)
		} else {
			*pancakes = replaceAtIndex(*pancakes, '+', j)
		}
	}
	return
}

func replaceAtIndex(in Pancakes, r rune, i int) Pancakes {
	out := []rune(string(in))
	out[i] = r
	return Pancakes(out)
}
