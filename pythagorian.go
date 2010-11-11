/*
 *Project Euler project 9
 *
 *Scott P. Powers
 *5-10-2010
 */

package main

import( 
	"fmt"
	"math"
)

func main(){
	const target = 1000
	cha:=make(chan int)
	chb:=make(chan int)
	go generate(cha, chb, target)
	go checkSum(cha, chb, target)	

}

func checkSum(cha, chb chan int, target float64){
	a:=<-cha
	b:=<-chb
	c:=math.Sqrt(float64(a*a+b*b))
	if float64(a+b)+c==target{
		fmt.Printf("ans = %f", float64(a*b)*c)
		return 
	}
}

func generate(cha chan int, chb chan int, target float64){
	for a:=3; ;a++{
		for b:=4;b<(int(target)-a)/2; b++{
			cha<-a
			chb<-b	
			fmt.Println("a: ", a, " b: ", b)	
		}
	}
}
	
