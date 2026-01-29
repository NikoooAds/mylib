package mylib

import "fmt"

// Something
func Hello(name string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += fmt.Sprintf("Hello, %s! ", name)
	}
	return result
}

func Add(a, b int) int {
	return a + b
}
