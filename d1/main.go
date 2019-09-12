package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	p1()
	p2()
}

func p1() {
	buf, err := ioutil.ReadFile("in.txt")
	in := string(buf)
	if err != nil {
		panic(err)
	}

	// in = in[:6]
	// fmt.Println(in)
	sum := 0

	for i, rune := range in {
		if i == len(in)-1 && byte(rune) == in[0] {
			sum += int(rune) - '0'
		} else if byte(rune) == in[i+1] {
			sum += int(rune) - '0'
		}
	}
	fmt.Println("p1 sum", sum)
}

func p2() {
	buf, err := ioutil.ReadFile("in.txt")
	in := string(buf)
	if err != nil {
		panic(err)
	}
	// in = in[:8]
	// fmt.Println(in)
	sum := 0

	for i, rune := range in {
		// fmt.Printf("%c %c", rune, in[(i+len(in)/2)%len(in)])
		if byte(rune) == in[(i+len(in)/2)%len(in)] {
			// fmt.Print("bam")
			sum += int(rune) - '0'
		}
		// fmt.Println()
	}
	fmt.Println("p2 sum", sum)
}
