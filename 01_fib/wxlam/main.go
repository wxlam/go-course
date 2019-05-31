package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// determines if fibonacci number is positive or negative
// based on index and whether number inputted is negative or not
func getPrintFibNum(index int, n int, isNeg bool) int {
	if isNeg && index%2 == 0 {
		return n * -1
	}
	return n
}

// print out fibonnaci numbers given n (number of numbers to print)
func fib(n int) {
	var fibArray []int
	var fibNum, counter int
	var isNeg bool

	if n < 0 {
		counter = n * -1
		isNeg = true
	} else {
		counter = n
		isNeg = false
	}
	// loop through to calculate fibonacci number print out number as it's calcuated
	for i := 0; i <= counter; i++ {
		switch i {
		case 0:
			fibNum = 0
			fibArray = append(fibArray, fibNum)
			fmt.Fprintln(out, fibNum)
		case 1:
			fibNum = 1
			fibArray = append(fibArray, fibNum)
			fmt.Fprintln(out, getPrintFibNum(i, fibNum, isNeg))
		default:
			fibNum = fibArray[i-1] + fibArray[i-2]
			fibArray = append(fibArray, fibNum)
			fmt.Fprintln(out, getPrintFibNum(i, fibNum, isNeg))
		}
	}
}

func main() {
	fib(7)
}
