package main

import "fmt"

func main() {
	fmt.Printf("The square root of given number is %f\n", sqrt_n(10))
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))

	var num1 int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&num1)
	fibSeries := fibonacci(num1)
	fmt.Println("Fibonacci series:", fibSeries)

}
