package main

import "fmt"

func main() {
	var str string
	var max int

	fmt.Scan(&str)

	for i := 0; i < len(str); i++ {
		digit := int(str[i]) - '0'
		if digit > max {
			max = digit
		}
	}

	fmt.Print(max)
}
