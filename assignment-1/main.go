package main

import "fmt"

func main() {
	integers := []int{1,2,3,4,5,6,7,8,9,10}
	for _, val := range integers {
		if(val % 2 == 0) {
			fmt.Printf("%d is Even\n", val);
		} else {
			fmt.Printf("%d is Odd\n", val);
		}
	}
}