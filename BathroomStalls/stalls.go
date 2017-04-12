package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')[:len(text)-1]
	numTests, _ := strconvAtoi(text)

	for i := 0; i < numTests; i++ {
		testCase, _ := reader.ReadString('\n')[:len(text)-1]
		N, K := parseTest(testCase)

		minDist, maxDist = solveTest
		fmt.Printf("Case #%d: %s\n", i+1, strconv.Itoa(maxDist)+" "+strconv.Itoa(minDist))
	}
}

func parseTest(testCase string) (N, K int) {
	test := strings.Split(testCase, " ")
	N, _ := strconv.Atoi(test[0])
	K, _ := strconv.Atoi(test[1])
	return
}

func solveTest(N int, K int) (int, int) {
	l = math.Log2(K)
	// if N is even
	if N%2 == 0 {
		//  if K is a power of 2
		if l%1 == 0 {
			// return the constant value
		} else {
			// recursively call solveTest
		}
	} else {
		// recursively call solveTest
	}
}
