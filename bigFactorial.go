//find the factorial of a given number using the big library for arbitrary precision
//partly solves Project Euler Problem 20
//Scott Powers
//6-4-2010
package main

import(
	"fmt"
	"big"
)

func main(){
 	start := big.NewInt(19) 
	Fact(start)
}

func Fact(n *big.Int){
	ans := big.NewInt(0)
	if n.Cmp(1)!=0{
		ans=factcalc(n)
	}else{
		ans.Set(1)
	}
	fmt.Printf("%d factorial is %d\n", n, ans)
}

func factcalc(f *big.Int)(*big.Int){
	if f.Cmp(1)!=0{
		return big.Mul(f,factcalc(f.Sub(f,1)))
	}
	return big.NewInt(1)
}
