package lemonade

import (
	"github.com/killtw/lemonade/trie"
	"strings"
	"unicode/utf8"
)

func init() {
	t := trie.New()
	t.Add("測試")
	t.Add("test")
	t.Add("123")

	Trie = t
}

func Replace(message string) (string, []string) {
	matches := Trie.Search(message)
	r := make([]string, 2*len(matches))

	for i, word := range matches {
		r[i*2] = word
		r[i*2+1] = strings.Repeat("*", utf8.RuneCountInString(word))
	}

	return strings.NewReplacer(r...).Replace(message), matches
}

func Add(word string) {
	Trie.Add(word)
}
