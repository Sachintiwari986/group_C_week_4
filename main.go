package main

import "fmt"

func main() {

	var num_squareroot float64
	fmt.Print("Enter a number to find out the square root: ")
	fmt.Scan(&num_squareroot)
	fmt.Printf("The square root of given number is %f\n", sqrt_n(num_squareroot))

	var num int
	fmt.Print("Enter a number to calculate factorial: ")
	fmt.Scan(&num)
	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))

	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&num)
	fibSeries := fibonacci(num)
	fmt.Println("Fibonacci series:", fibSeries)

	fmt.Print("Enter the number to check it is even or odd:")
	fmt.Scan(&num)
	if isEven(num) {
		fmt.Println("The given number is even")
	} else {
		fmt.Println("The given number is odd")

	}

}
