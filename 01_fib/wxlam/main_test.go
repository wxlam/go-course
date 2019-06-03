package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutputFib7(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("0\n1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main() | \n%v\n%v", actual, expected)
	}
}

func TestFib0(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(0)

	expected := strconv.Quote("0\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(0) | \n%v\n%v", actual, expected)
	}
}

func TestFib1(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(1)

	expected := strconv.Quote("0\n1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(1) | \n%v\n%v", actual, expected)
	}
}

func TestMainOutputFibNeg7(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(-7)

	expected := strconv.Quote("0\n1\n-1\n2\n-3\n5\n-8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in fib(-7) | \n%v\n%v", actual, expected)
	}
}
