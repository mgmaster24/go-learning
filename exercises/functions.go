package gol_exercises

import (
	"fmt"
	"time"
)

type transformFunc func(int) int

func RunFunctionsExercise() {
	numbers := []int{1, 2, 3}
	fmt.Printf("Transforming Values by providing function: %v\n", transformNumbers(&numbers, createTransformer(2)))
	fmt.Println()

	var factorialVal int
	fmt.Print("Please enter the value to factorial: ")
	fmt.Scan(&factorialVal)
	fmt.Println(factorial(factorialVal))
	fmt.Println()

	start := time.Now()
	fmt.Println(fibonacci(10))
	fmt.Printf("fibonacci - standalone function: %v\n", time.Since(start).Nanoseconds())
	fmt.Println()

	start = time.Now()
	fmt.Println(fibonacci2(10))
	fmt.Printf("fibonacci - closure: %v\n", time.Since(start).Nanoseconds())
	fmt.Println()

	start = time.Now()
	f := func(levels int) []int {
		fib := make([]int, levels)
		fib1, fib2 := 0, 1
		for i := 0; i < levels; i++ {
			fib[i] = fib1
			fib1, fib2 = fib2, fib1+fib2
		}

		return fib
	}
	fmt.Println(f(10))
	fmt.Printf("fibonacci - anonymous function: %v\n", time.Since(start).Nanoseconds())
	fmt.Println()

	nums := []int{1, 10, 15}
	fmt.Println(sumup(nums))
	fmt.Println(sumupVariadic(1, 2, 3, 4))
	fmt.Println(sumupVariadic(nums...))
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) transformFunc {
	return func(val int) int {
		return val * factor
	}
}

func fibonacci(levels int) []int {
	fib := make([]int, levels)
	fib1, fib2 := 0, 1
	for i := 0; i < levels; i++ {
		fib[i] = fib1
		fib1, fib2 = fib2, fib1+fib2
	}

	return fib
}

func fibonacci2(levels int) []int {
	f := fibonacciClosure()
	fib := make([]int, levels)
	for i := 0; i < levels; i++ {
		fib[i] = f()
	}

	return fib
}

func fibonacciClosure() func() int {
	fib2, fib1 := 0, 1

	return func() int {
		f := fib2
		fib2, fib1 = fib1, fib1+fib2
		return f
	}
}

func factorial(val int) int {
	if val == 1 {
		return val
	}

	return val * factorial(val-1)
}

func sumup(nums []int) int {
	sum := 0

	for _, val := range nums {
		sum += val
	}

	return sum
}

func sumupVariadic(nums ...int) int {
	sum := 0

	for _, val := range nums {
		sum += val
	}

	return sum
}
