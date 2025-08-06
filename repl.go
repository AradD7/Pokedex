package main

import (
	"slices"
	"strings"
)

func cleanInput(text string) []string {
	seperatedText := strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
	seperatedText = slices.DeleteFunc(seperatedText, func(s string) bool {
		return s == ""
	})
	return seperatedText
}



