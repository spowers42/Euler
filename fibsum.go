/* 
 * project euler problem 2
 * Scott P. Powers
 * 4-26-2010
 *
 */
package main

import "fmt"

func main(){
	sum:=0
	nums:=[]int{0,1,0}

	for ;nums[2]<4000000;{
		nums[2]=nums[0]+nums[1]
		nums[0]= nums[1]
		nums[1]= nums[2]
		if nums[2]%2==0{
			sum+=nums[2]
		}
	}

	fmt.Printf("\nSum: %d\n", sum)
}
	
