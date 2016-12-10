package main

import "fmt"

func main() {
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %05e \t %0.6X \n", i, i, float32(i), i)
	}
}
