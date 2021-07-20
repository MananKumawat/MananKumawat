package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addnew(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	// Functions
	fmt.Println(add(42, 13))

	// Functions continued
	fmt.Println(addnew(42, 13))

	// Multiple results
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// Named return values
	fmt.Println(split(17))

	// Variadic Functions
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// Closures
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())

	// Recursion
	fmt.Println(fact(8))

}