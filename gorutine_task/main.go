package main

import (
	"fmt"
	"strconv"
)

func main11() {
	sliceInt := []int{0, 1, 1, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range sliceInt {
		if val%2 == 0 {
			fmt.Printf("%v is even\n", val)
		} else {
			fmt.Printf("%v is odd\n", val)
		}
	}
}

func main22() {
	greeting := "Hi There!"

	go func() {
		fmt.Println(greeting)
	}()
}

func main() {
	c := make(chan string)

	for i := 0; i < 4; i++ {
		go printString("Hello there! ", i, c)
	}

	fmt.Println(len(c))
	for i := 0; i < 4; i++ {
		fmt.Println(<-c)
	}
}

func printString(s string, i int, c chan string) {
	// fmt.Println(s, i)
	str := "Done printing." + strconv.Itoa(i)
	c <- str
}
