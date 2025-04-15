package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

type Words struct {
	Word         string         `json:"word"`
	Translations []Translations `json:"translations"`
	Phrases      []Phrases      `json:"phrases"`
}

type Translations struct {
	Translation string `json:"translation"`
	Type        string `json:"type"`
}

type Phrases struct {
	Phrase      string `json:"phrase"`
	Translation string `json:"translation"`
}

func main() {
	start := time.Now()

	//读取json
	wordsCET4 := ReadJSON("3-CET4-顺序.json")
	wordsCET6 := ReadJSON("4-CET6-顺序.json")

	//合并数据
	words := append(wordsCET4, wordsCET6...)

	//删除旧的db
	os.Remove("./words.db")

	db, err := sql.Open("sqlite", "./words.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	SaveWords(db, words)

	fmt.Printf("数据插入完毕，共插入%d个单词，耗时：%.2f 秒\n", len(words), time.Since(start).Seconds())
}

func ReadJSON(filename string) []Words {
	file, err := os.Open(filename)
	if err != nil {
		panic("读取失败：" + filename)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("读取数据失败 %s: %v", filename, err)
	}

	var words []Words
	if err := json.Unmarshal(data, &words); err != nil {
		panic("解析失败：" + filename)
	}

	return words
}

func createTable(db *sql.DB) {
	sqlStr := `
	CREATE TABLE IF NOT EXISTS words (
		word TEXT PRIMARY KEY UNIQUE,
		translation TEXT,
		type TEXT,
		phrase TEXT
	);`
	_, err := db.Exec(sqlStr)
	if err != nil {
		log.Fatal("建表失败:", err)
	}
}

func SaveWords(db *sql.DB, words []Words) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT OR REPLACE INTO words(word, translation, type, phrase) VALUES (?, ?, ?, ?)")
	defer stmt.Close()

	for _, w := range words {
		var transList, typeList, phraseList []string

		for _, t := range w.Translations {
			transList = append(transList, t.Translation)
			typeList = append(typeList, t.Type)
		}
		for _, p := range w.Phrases {
			phraseList = append(phraseList, fmt.Sprintf("%s(%s)", p.Phrase, p.Translation))
		}

		stmt.Exec(w.Word, strings.Join(transList, "|"), strings.Join(typeList, "|"), strings.Join(phraseList, "|"))
	}
	tx.Commit()
}

// func SaveWords(db *sql.DB, words []Word) {
// 	for _, w := range words {
// 		var transList []string
// 		var typeList []string
// 		var phraseList []string

// 		for _, t := range w.Translations {
// 			transList = append(transList, t.Translation)
// 			typeList = append(typeList, t.Type)
// 		}

// 		for _, p := range w.Phrases {
// 			phraseList = append(phraseList, fmt.Sprintf("%s(%s)", p.Phrase, p.Translation))
// 		}

// 		_, err := db.Exec("INSERT OR REPLACE INTO words(word, translation, type, phrase) VALUES (?, ?, ?, ?)",
// 			w.Word, strings.Join(transList, "|"), strings.Join(typeList, "|"), strings.Join(phraseList, "|"))
// 		if err != nil {
// 			log.Printf("插入失败 %s: %v\n", w.Word, err)
// 		}
// 	}
// }
