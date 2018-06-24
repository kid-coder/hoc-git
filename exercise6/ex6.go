package main

import (
	"fmt"
)

func sum(a []int, ch chan int) {
	sum := 0
	for _, num := range a {
		sum += num
	}
	ch <- sum
}

func main() {
	a := []int{5, 6, 2, 1, 5, 6, 4, 2, 3}
	ch := make(chan int)
	go sum(a[:len(a)/2], ch)
	go sum(a[len(a)/2:len(a)], ch)
	sum1, sum2 := <-ch, <-ch
	fmt.Printf("Sum: %d\n", sum1+sum2)
}
