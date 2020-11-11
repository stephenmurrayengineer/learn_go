package iteration

import "strings"

const repeatCount = 5

func Repeat(character string, n int) (retstr string) {
	var repeated string
	for i := 0; i<n; i++ {
		repeated += character
	}
	return repeated
}

func MakeTitle(title string) (cleaned string) {
	title = strings.TrimSpace(title)
	title = strings.ReplaceAll(title, ",", " ")
	cleaned = strings.Title(title)
	return cleaned
}
