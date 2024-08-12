package main

import (
	"html"
	"math/rand"
	"regexp"
	"sort"
	"strings"
)

// 1)
func Concaten(s1, s2 string) string {
	conc := s1 + " " + s2
	return conc
}

// 2)
func returnLen(s string) int {
	return len(s)
}

// 3
func containsSubstring(str, substr string) bool {
	return strings.Contains(str, substr)
}

// 4)
func findSubstringIndex(str, substr string) int {
	return strings.Index(str, substr)
}

// 5)
func replaceAllSupStr(str, find, new string) string {
	return strings.ReplaceAll(str, find, new)
}

// 6)
func trimStr(str string) string {
	return strings.TrimSpace(str)
}

// 7)
func convertToTo(str string) (string, string) {
	upper := strings.ToUpper(str)
	lower := strings.ToLower(str)
	return upper, lower
}

// 8)
func repeatStr(str string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += "\n" + str
	}
	return result
}

// 9)
func splitStr(str, separator string) []string {
	return strings.Split(str, separator)
}

// 10)
func compareEqualFold(str1, str2 string) bool {
	return strings.EqualFold(str1, str2)

}

// 11)
func replaceFirst(str, find, new string) string {
	return strings.Replace(str, find, new, 1)
}

// 12)
func reverseStrRune(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}

// 13)
func countChars(str string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range str {
		counts[r]++
	}
	return counts
}

// 14)
func removeChar(s string, char rune) string {
	return strings.ReplaceAll(s, string(char), "")
}

// 15)
func countWorks(str string) int {
	words := strings.Fields(str)
	return len(words)
}

// 16)
func hasPrefix(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}
func hasSuffix(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

// 17)
func removerDupl(str string) string {
	seen := make(map[rune]bool)
	var result []rune

	for _, r := range str {
		if !seen[r] {
			seen[r] = true
			result = append(result, r)
		}
	}

	return string(result)
}

// 18)
func formatHTMJl(str string) string {
	return html.EscapeString(str)
}

// 19)
func createSlug(str string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	slug := reg.ReplaceAllString(str, "_")
	return slug
}

// 21)
func reverseWords(str string) string {
	words := strings.Fields(str)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	result := strings.Join(words, " ")
	return result
}

// 24)
func shuffleWords(str string) string {
	words := strings.Fields(str)

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	return strings.Join(words, " ")
}

// 25)
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func sordWordsLength(str string) string {
	words := strings.Fields(str)

	sort.Sort(byLength(words))

	return strings.Join(words, " ")
}

func countWordsCharactes(str string) (int, int) {
	words := strings.Fields(str)
	charactes := len(str)

	return len(words), charactes

}
