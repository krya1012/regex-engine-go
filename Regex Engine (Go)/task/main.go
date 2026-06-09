package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func matchOne(regex, text string) bool {
	if len(regex) == 0 {
		return true
	}
	if len(text) == 0 {
		return false
	}
	return regex == "." || regex == text
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.SplitN(scanner.Text(), "|", 2)
	fmt.Println(matchOne(parts[0], parts[1]))
}
