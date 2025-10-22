package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Pokedex > ")
		reader.Scan()

		wordList := cleanInput(reader.Text())
		if len(wordList) == 0 {
			continue
		}

		msg := fmt.Sprintf("Your command was: %s", wordList[0])
		fmt.Println(msg)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
