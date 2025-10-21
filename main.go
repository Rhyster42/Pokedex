package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	text_list := strings.Fields(text)
	text_list_lower := []string{}
	for i := range text_list {
		text_list_lower = append(text_list_lower, strings.ToLower(text_list[i]))
	}
	return text_list_lower
}
