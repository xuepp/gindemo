package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Word 结构体
type Word struct {
	ID        int         `json:"id"`
	Word      string      `json:"word"`
	Phonetic0 string      `json:"phonetic0"`
	Phonetic1 string      `json:"phonetic1"`
	Trans     []Tran      `json:"trans"`
	Sentences []Sentence  `json:"sentences"`
	Phrases   []Phrase    `json:"phrases"`
	Synos     []Syno      `json:"synos"`
	RelWords  RelWords    `json:"relWords"`
	Etymology []Etymology `json:"etymology"`
}

type Tran struct {
	Pos string `json:"pos"`
	Cn  string `json:"cn"`
}

type Phrase struct {
	C  string `json:"c"`
	Cn string `json:"cn"`
}

type Sentence struct {
	C  string `json:"c"`
	Cn string `json:"cn"`
}

type Syno struct {
	Pos string   `json:"pos"`
	Cn  string   `json:"cn"`
	Ws  []string `json:"ws"`
}

// 派生词（relWords）
type RelWords struct {
	Root string       `json:"root"`
	Rels []RelWordPos `json:"rels"`
}

// 派生词按词性分类
type RelWordPos struct {
	Pos   string       `json:"pos"`
	Words []RelWordSub `json:"words"`
}

type RelWordSub struct {
	C  string `json:"c"`
	Cn string `json:"cn"`
}

type Etymology struct {
	T string `json:"t"`
	D string `json:"d"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/dict/words/:id", func(c *gin.Context) {
		dictID := c.Param("id")

		// 通过dict_id查询 dict_to_word 下的所有 word_id
		wordIDs := []int{}
		rows, err := db.Query("SELECT word_id FROM dict_to_word WHERE dict_id = ?", dictID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//spew.Dump("rows", rows)
		defer rows.Close()

		//获取所有 word_id
		for rows.Next() {
			var wid int
			rows.Scan(&wid)
			wordIDs = append(wordIDs, wid)
		}

		if len(wordIDs) == 0 {
			c.JSON(http.StatusOK, []Word{})
			return
		}

		//fmt.Println("查询的 word IDs:", wordIDs)

		// 查询 words 基础信息
		query, args := buildInQuery("SELECT id, word, phonetic0, phonetic1 FROM words WHERE id IN (?)", wordIDs)
		rows, err = db.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("执行的查询语句:", query)
		fmt.Println("查询参数:", args)
		//fmt.Print("rows", rows)
		//spew.Dump("rows-----", rows.Next())
		defer rows.Close()

		words := []Word{}
		for rows.Next() {
			var w Word
			rows.Scan(&w.ID, &w.Word, &w.Phonetic0, &w.Phonetic1)
			words = append(words, w)
		}
		fmt.Print("words", words)

		// 填充关联表
		for i := range words {
			w := &words[i]
			w.Trans = queryTrans(db, w.ID)
			w.Phrases = queryPhrases(db, w.ID)
			w.Sentences = querySentences(db, w.ID)
			w.Synos = querySynos(db, w.ID)
			w.RelWords = queryRelWords(db, w.ID)
			w.Etymology = queryEtymology(db, w.ID)
		}

		c.JSON(http.StatusOK, words)
	})

	r.Run(":8080")
}

// 构建 IN 查询
func buildInQuery(base string, args []int) (string, []interface{}) {
	params := make([]interface{}, len(args))
	placeholders := make([]string, len(args))
	for i, v := range args {
		params[i] = v
		placeholders[i] = "?"
	}
	query := strings.Replace(base, "(?)", "("+strings.Join(placeholders, ",")+")", 1)
	return query, params
}

// 查询翻译
func queryTrans(db *sql.DB, wordID int) []Tran {
	rows, _ := db.Query("SELECT pos, cn FROM word_trans WHERE word_id=?", wordID)
	defer rows.Close()
	list := []Tran{}
	for rows.Next() {
		var t Tran
		rows.Scan(&t.Pos, &t.Cn)
		list = append(list, t)
	}
	return list
}

// 查询短语
func queryPhrases(db *sql.DB, wordID int) []Phrase {
	rows, _ := db.Query("SELECT c, cn FROM word_phrases WHERE word_id=?", wordID)
	defer rows.Close()
	list := []Phrase{}
	for rows.Next() {
		var t Phrase
		rows.Scan(&t.C, &t.Cn)
		list = append(list, t)
	}
	return list
}

// 查询例句
func querySentences(db *sql.DB, wordID int) []Sentence {
	rows, _ := db.Query("SELECT c, cn FROM word_sentences WHERE word_id=?", wordID)
	defer rows.Close()
	list := []Sentence{}
	for rows.Next() {
		var t Sentence
		rows.Scan(&t.C, &t.Cn)
		list = append(list, t)
	}
	return list
}

// 查询词源
func queryEtymology(db *sql.DB, wordID int) []Etymology {
	rows, _ := db.Query("SELECT t,d FROM word_etymology WHERE word_id=?", wordID)
	defer rows.Close()
	list := []Etymology{}
	for rows.Next() {
		var t Etymology
		rows.Scan(&t.T, &t.D)
		list = append(list, t)
	}
	return list
}

// 查询同义词
func querySynos(db *sql.DB, wordID int) []Syno {
	rows, _ := db.Query("SELECT id,pos,cn FROM word_synos WHERE word_id=?", wordID)
	defer rows.Close()
	list := []Syno{}
	for rows.Next() {
		var s Syno
		var synoID int
		rows.Scan(&synoID, &s.Pos, &s.Cn)

		wsRows, _ := db.Query("SELECT ws FROM word_syno_ws WHERE syno_id=?", synoID)
		defer wsRows.Close()
		s.Ws = []string{}
		for wsRows.Next() {
			var w string
			wsRows.Scan(&w)
			s.Ws = append(s.Ws, w)
		}
		list = append(list, s)
	}
	return list
}

// 查询相关词
func queryRelWords(db *sql.DB, wordID int) RelWords {
	var relWords RelWords

	// 查询 root
	err := db.QueryRow("SELECT root FROM word_relwords WHERE word_id=? LIMIT 1", wordID).Scan(&relWords.Root)
	if err != nil {
		// 没查到返回空结构体
		return relWords
	}

	// 查询每个词性分组（pos）
	rows, err := db.Query("SELECT id, pos FROM word_relwords WHERE word_id=?", wordID)
	if err != nil {
		fmt.Println("查询 word_relwords 出错:", err)
		return relWords
	}
	defer rows.Close()

	for rows.Next() {
		var relID int
		var pos string
		if err := rows.Scan(&relID, &pos); err != nil {
			continue
		}

		// 当前词性组
		relPos := RelWordPos{
			Pos:   pos,
			Words: []RelWordSub{},
		}

		// 查询对应的单词子项
		itemRows, err := db.Query("SELECT c, cn FROM word_relwords_items WHERE rel_id=?", relID)
		if err != nil {
			fmt.Println("查询 word_relwords_items 出错:", err)
			continue
		}

		for itemRows.Next() {
			var sub RelWordSub
			itemRows.Scan(&sub.C, &sub.Cn)
			relPos.Words = append(relPos.Words, sub)
		}
		itemRows.Close()

		relWords.Rels = append(relWords.Rels, relPos)
	}

	return relWords
}
