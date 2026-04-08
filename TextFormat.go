package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hextodec(s string) (int64, error) {
	return strconv.ParseInt(s, 16, 64)
}

func bintodec(s string) (int64, error) {
	return strconv.ParseInt(s, 2, 64)
}

func capitalizeFirstLetter(s string) string {
	// return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
	w := strings.ToLower(s)
	return strings.Title(w)
}

func uppercaseLastWord(s []string, n int) []string {
	for i := 1; i <= n; i++ {
		w := len(s) - i
		s[w] = strings.ToUpper(s[w])
	}
	return s
}

func punctuation(s string) bool {
	return strings.ContainsAny(s, ".,;!'")
}

func fixPunctuation(str []string) string {
	s := strings.Join(str, " ")
	s = strings.ReplaceAll(s, " ,", ",")
	s = strings.ReplaceAll(s, " !", "!")
	return s
}

func article(s string) string {
	first := s[:1]
	article := "aeiouhAEIOUH"
	if strings.ContainsAny(first, article) {
		return "An " + s
	}
	return "A " + s
}

func fixArticle(s string) string {
	s = strings.ReplaceAll(s, "a a", "an a")
	s = strings.ReplaceAll(s, "a e", "an e")
	s = strings.ReplaceAll(s, "a h", "an h")
	return s
}

func fixQuote(s string) string {
	s = strings.Trim(s, "'")
	s = strings.TrimSpace(s)
	return "'" + s + "'"
}

func toUpper(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToUpper(s[1:])
}

func toLower(s string) string {
	return strings.ToLower(s[:1]) + strings.ToLower(s[1:])
}

func isCommand(s string) bool {
	return strings.HasPrefix(s, "(")
	//return s == "up" || s == "cap" || s == "hex" || s == "bin" || s == "low"
}

func removeCommand(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func fixPunctuation2(str []string) string {
	s := strings.Join(str, " ")
	s = strings.ReplaceAll(s, " ,", ",")
	s = strings.ReplaceAll(s, " !", "!")
	return s
}

func isGroupedPunctuation(s string) bool {
	if len(s) <= 1 {
		return false
	}

	return true
}

func uppercaseN(s []string, n int) []string {
	for i := 1; i <= n; i++ {
		w := len(s) - i
		s[w] = strings.ToUpper(s[w])
	}
	return s
}

func lowercaseN(s []string, n int) []string {
	for i := 1; i <= n; i++ {
		w := len(s) - 1
		s[w] = strings.ToLower(s[w])
	}
	return s
}

func main() {
	fmt.Println(hextodec("1E"))
	fmt.Println(hextodec("FF"))
	fmt.Println(bintodec("10"))
	fmt.Println(bintodec("1010"))
	fmt.Println(bintodec("11111111"))
	fmt.Println(capitalizeFirstLetter("hELLO"))
	fmt.Println(capitalizeFirstLetter("WORLD"))
	fmt.Println(uppercaseLastWord([]string{"I", "love", "my", "life"}, 2))
	fmt.Println(punctuation("."))
	fmt.Println(punctuation("!"))
	fmt.Println(punctuation("x"))
	fmt.Println(fixPunctuation([]string{"Hello", ",", "World", "!"}))
	fmt.Println(article("Boy"))
	fmt.Println(article("horse"))
	fmt.Println(article("apple"))
	fmt.Println(article("honest"))
	fmt.Println(fixArticle("there it was. a honest man. a amazing man"))
	fmt.Println(fixQuote("' Hello '"))
	fmt.Println(fixQuote("' Awesome '"))
	fmt.Println(toUpper("hello"))
	fmt.Println(toLower("HELLO"))
	fmt.Println(isCommand("(up)"))
	fmt.Println(isCommand("(cap)"))
	fmt.Println(isCommand("hello"))
	fmt.Println(removeCommand([]string{"hello", "(up)"}))
	fmt.Println(fixPunctuation2([]string{"Hello", ",", "World"}))
	fmt.Println(isGroupedPunctuation("..."))
	fmt.Println(isGroupedPunctuation("!?"))
	fmt.Println(isGroupedPunctuation("."))
	fmt.Println(uppercaseN([]string{"Hello", "my", "love"}, 2))
	fmt.Println(lowercaseN([]string{"HELLO", "MY", "LOVE"}, 2))

}
