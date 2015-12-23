package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("len(slice) = %d\n", len(slice))
	fmt.Printf("cap(slice) = %d\n", cap(slice))
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Hi, I'm %d\n", slice[i])
	}
	slice = append(slice, []int{11, 12, 13, 14, 15}...)
	fmt.Printf("len(slice) = %d\n", len(slice))
	fmt.Printf("cap(slice) = %d\n", cap(slice))
}
