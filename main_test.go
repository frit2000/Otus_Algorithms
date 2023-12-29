package main

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func numFromFile(fileName string) int {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataStr := strings.Trim(string(data), "\r\n")
	result, err := strconv.Atoi(string(dataStr))
	if err != nil {
		panic(err)
	}
	return result
}

func TestLucky(t *testing.T) {
	i := 0
	for {
		fNameIn := "tests/test." + strconv.Itoa(i) + ".in"
		fNameOut := "tests/test." + strconv.Itoa(i) + ".out"
		_, err1 := os.Stat(fNameIn)
		_, err2 := os.Stat(fNameOut)
		if err1 != nil || err2 != nil {
			break
		}
		n := numFromFile(fNameIn)
		want := numFromFile(fNameOut)

		have := LuckyTickets(int(n))

		if int(want) != have {
			t.Errorf("Result was incorrect for n=%d, got: %d, want: %d.", i, have, want)
		}
		i++
	}
}
