/* 
 * project euler problem 1
 * Scott P. Powers
 * 4-26-2010
 *
 */
package main

import "fmt"

func main(){
     sum := 0
     for i:=1000; i>0;i--{
     	 if i%3==0 || i%5==0{
	    sum+=i
	    }
     }
     fmt.Printf("\n%d\n", sum)	   
}
