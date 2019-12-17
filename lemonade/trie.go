package lemonade

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/killtw/lemonade/trie"
	"os"
)

var dbHost = getenv("DB_HOST", "localhost")

type Word struct {
	Keyword string
}

func InitTrie() (err error) {
	t := trie.New()

	words, err := getWords()

	if err != nil {
		return
	}

	for _, word := range words {
		t.Add(word.Keyword)
	}

	Trie = t

	fmt.Println("Words imported")

	return
}

func getWords() (words []Word, err error) {
	db, err := gorm.Open("mysql", dbHost)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	db.Find(&words)

	return
}

func getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
