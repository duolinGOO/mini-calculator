package main

import "fmt"

func main() {
	pot1 := make(chan int)
	pot2 := make(chan int)
	pot3 := make(chan int)

	fmt.Print(<-sum(5, 9, pot1))

	fmt.Print(" ")
	fmt.Print(<-multipl(5, 23, pot2))
	fmt.Print(" ")
	fmt.Print(<-divis(9, 3, pot3))

}
func sum(x int, y int, ch1 chan int) chan int {
	go func() {
		a := x + y
		ch1 <- a
	}()
	return ch1
}
func multipl(x int, y int, ch2 chan int) chan int {

	go func() {
		a := x * y
		ch2 <- a
	}()
	return ch2
}
func divis(x int, y int, ch3 chan int) chan int {
	go func() {
		a := x / y
		ch3 <- a
	}()
	return ch3
}
