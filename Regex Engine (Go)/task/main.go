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

func matchEqual(regex, text string) bool {
	if len(regex) == 0 {
		return true
	}
	if regex == "$" && len(text) == 0 {
		return true
	}
	if len(regex) >= 2 && regex[0] == '\\' {
		if len(text) == 0 {
			return false
		}
		return regex[1:2] == text[:1] && matchEqual(regex[2:], text[1:])
	}
	if len(regex) >= 2 {
		switch regex[1] {
		case '?':
			return matchEqual(regex[2:], text) ||
				(len(text) > 0 && matchOne(regex[:1], text[:1]) && matchEqual(regex[2:], text[1:]))
		case '*':
			return matchEqual(regex[2:], text) ||
				(len(text) > 0 && matchOne(regex[:1], text[:1]) && matchEqual(regex, text[1:]))
		case '+':
			return len(text) > 0 && matchOne(regex[:1], text[:1]) &&
				(matchEqual(regex[2:], text[1:]) || matchEqual(regex, text[1:]))
		}
	}
	if len(text) == 0 {
		return false
	}
	return matchOne(regex[:1], text[:1]) && matchEqual(regex[1:], text[1:])
}

func matchAnywhere(regex, text string) bool {
	if len(regex) > 0 && regex[0] == '^' {
		return matchEqual(regex[1:], text)
	}
	for i := 0; i <= len(text); i++ {
		if matchEqual(regex, text[i:]) {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.SplitN(scanner.Text(), "|", 2)
	fmt.Println(matchAnywhere(parts[0], parts[1]))
}
