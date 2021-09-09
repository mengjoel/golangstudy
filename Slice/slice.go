package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	//append
	numbers = append(numbers, 1, 2, 3, 4, 5)

	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers) * 2))
	//copy
	copy(numbers1, numbers)
	printSlice(numbers1)
}
func printSlice(x []int) {
	fmt.Printf("%d %d %v", len(x), cap(x), x)
}
