package main

import (
	"fmt"
)

func main(){
	var val int
	var i int
	for i = 0; i <= 1000; i++ { //loop 0-1000
		if i % 3 == 0 || i % 5 == 0 {
			val += i;
		}
	}
	fmt.Printf("Sum of all the multiples of 3 or 5 below 1000 %d", val)
	
}