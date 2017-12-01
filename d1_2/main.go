package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1]
	sum := 0
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+(len(input)/2))%len(input)] {
			sum += (int(input[i]) - 48)
		}
	}
	fmt.Println(sum)
}
