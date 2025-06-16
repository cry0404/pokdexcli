package main

import (
	"strings"
)

func cleanInput(text string) []string{
	input := text
	cleaned := strings.ToLower(strings.TrimSpace(input))
	words := strings.Fields(cleaned)

	return words

}