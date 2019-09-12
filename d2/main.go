package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("in.txt")
	if err != nil {
		panic(err)
	}
	p1(string(buf))
	p2(string(buf))
}

func p1(buf string) {
	sum := 0
	lines := strings.Split(buf, "\n")
	for i := 0; i < len(lines); i++ {
		// fmt.Print(strings.Fields(lines[i]))
		max := arrCompare(strings.Fields(lines[i]), getMax)
		min := arrCompare(strings.Fields(lines[i]), getMin)
		diff := max - min
		// fmt.Printf(" max=%d min=%d diff=%d\n", max, min, diff)
		sum += diff
	}
	fmt.Println("p1 sum", sum)
}

func p2(buf string) {
	sum := 0
	lines := strings.Split(buf, "\n")
	for i := 0; i < len(lines); i++ {
		lineVals := strings.Fields(lines[i])
		// fmt.Printf("Line %d  \t", i)
		for j := 0; j < len(lineVals); j++ {
			dividesEvenly, result := anyDividesEvenly(lineVals, lineVals[j], j)
			if dividesEvenly {
				sum += result
				break
			}
		}
	}
	fmt.Println("p2 sum", sum)
}

func anyDividesEvenly(arr []string, divisor string, divisorIndex int) (dividesEvenly bool, result int) {
	for i := 0; i < len(arr); i++ {
		if i == divisorIndex {
			continue
		}
		num, _ := strconv.Atoi(arr[i])
		divisor, _ := strconv.Atoi(divisor)
		dividesEvenly = num%divisor == 0
		if dividesEvenly {
			result = num / divisor
			// fmt.Printf("%d / %d = %d\n", num, divisor, result)
			return
		}
	}
	return
}

func arrCompare(arr []string, f func(int, int) int) int {
	val, _ := strconv.Atoi(arr[0])
	for i := 0; i < len(arr); i++ {
		otherVal, _ := strconv.Atoi(arr[i])
		val = f(otherVal, val)
	}
	return val
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
