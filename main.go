package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Pokedex > ")

		scanner.Scan()

		line := scanner.Text()
		line_list := strings.Fields(line)
		msg := fmt.Sprintf("Your command was: %s", line_list[0])
		fmt.Println(msg)
	}

}
