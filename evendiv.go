/*
 * project euler problem 5
 * Scott P. Powers
 * 4-27-2010
 */
package main

import "fmt"

func main(){
	found :=0
	tmp:=0
	largest:=20
	for i:=2; found==0; i++{
		for n:=1; n<=largest;n++{
			if i%n==0{
				tmp++
				if tmp<n{
					break
				}
			}
		}
		if tmp==largest{
			fmt.Printf("\n%d\n", i)	
			found=1
		}
		tmp=0
	
	}
	
}

