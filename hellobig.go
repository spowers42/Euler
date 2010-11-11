//find the factorial of a given number using the big library for arbitrary precision
//partly solves Project Euler Problem 20
//Scott Powers
//6-4-2010

package main

import (
	"fmt"
	"big"
	"flag"
)

func main() {
	input := flag.Int64("num", 100, "num!")
	flag.Parse()
	a := big.NewInt(*input)
	b := big.NewInt(1)
	for ; a.Cmp(big.NewInt(0)) > 0; a.Sub(a, big.NewInt(1)) {
		b.Mul(a, b)
	}
	c := b.String()
	fmt.Println(c)
}
