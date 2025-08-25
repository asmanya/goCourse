package main

import "fmt"

func main() {

	// Simple iteration over a range
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// iterate over collections
	// numbers := []int{1,2,3,4,5,6,7}
	// for idx, val := range numbers {
	// 	fmt.Println("Index:", idx, "Value:", val)
	// }

	// loop with break condition
	// for i := 1; i <= 10; i++ {
	// 	if i % 2 == 0 {
	// 	 	continue
	// 	}
	// 	fmt.Println("Odd number:", i)
	// 	if i == 5 {
	// 		break // breaks out of the loop
	// 	}
	// }

	// rows := 10

	// // Outer loop
	// for i := 1; i <= rows ; i++ {
	// 	// Inner loop for spaces before stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	// Inner loop for stars
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // move to next line
	// }

	for i := range 10 {
		i++
		fmt.Println(i)
	}
	fmt.Println("we have a lift off!")
}