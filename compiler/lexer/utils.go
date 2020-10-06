package lexer

import "regexp"

func isSpace(str string) bool {
	if str == "\n" || str == " " || str == "\t" {
		return true
	}
	return false
}

func isDigit(str string) bool {
	ok, _ := regexp.MatchString(`[0-9]+`, str)
	return ok
}

func isWord(str string) bool {
	ok, _ := regexp.MatchString(`[a-zA-Z]+`, str)
	return ok
}
