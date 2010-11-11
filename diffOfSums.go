/* 
 * project euler problem 6
 * Scott P. Powers
 * 4-27-2010
 *
 */
package main

import "fmt"

func main(){
	sums:=[]int{0,0}

	for i:=0; i<=100; i++{
		sums[0]+=i*i
		sums[1]+=i
	}
	sums[1]=sums[1]*sums[1]

	ans:=sums[1]-sums[0]

	fmt.Printf("Difference of sums\n%d\n", ans)
}
 
