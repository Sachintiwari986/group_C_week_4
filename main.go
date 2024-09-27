package main

import "fmt"

func main() {
	fmt.Printf("The square root of given number is %f", sqrt_n(10))
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}
