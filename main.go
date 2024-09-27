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

	var num2 int
	fmt.Print("Enter the number to check it is even or odd:")
	fmt.Scan(&num2)
	if isEven(num2) {
		fmt.Println("The given number is even")
	} else {
		fmt.Println("The given number is odd")

	}

}
