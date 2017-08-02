package main

import (
	"fmt"
)

func main() {
	fmt.Println("print 10 natural numbers")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Print 10 even natural numbers")
	j := 1
	for i := 1; true; i++ {

		if i%2 == 0 {
			fmt.Println(i)
			j++
			if j > 10 {
				break
			}
		}
	}

	fmt.Println("Print first 20 fibonacci numbers")
	j = 1
	a := 1
	fmt.Println(a)

	for b := 1; j < 21; j++ {

		fmt.Println(b)
		b = a + b
		a = b - a
	}

	fmt.Println("Print first  100 natural numbers, with fizz buuz and fizzbuzz")

	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println(i, " : fizzbuzz")
			} else {
				fmt.Println(i, " :  fizz")
			}
		} else {
			if i%5 == 0 {
				fmt.Println(i, " :  buzz")
			} else {
				fmt.Println(i)
			}
		}

	}

}
