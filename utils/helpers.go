package utils

import "fmt"

func InputInt(label string) int {
	var v int
	fmt.Print(label)
	fmt.Scanln(&v)
	return v
}
