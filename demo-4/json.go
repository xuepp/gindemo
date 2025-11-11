package main

import (
	"encoding/json"
	"fmt"
)

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

type RelWords struct {
	Root string       `json:"root"`
	Rels []RelWordPos `json:"rels"`
}

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
	word := Word{
		ID:        1,
		Word:      "develop",
		Phonetic0: "dɪˈveləp",
		Phonetic1: "dəˈvɛləp",
		Trans: []Tran{
			{Pos: "v.", Cn: "发展；开发；培养"},
		},
		Sentences: []Sentence{
			{C: "He wants to develop his communication skills.", Cn: "他想要提高自己的沟通能力。"},
			{C: "The company plans to develop a new product line.", Cn: "公司计划开发一条新的产品线。"},
		},
		Phrases: []Phrase{
			{C: "develop into", Cn: "发展成；变成"},
			{C: "develop from", Cn: "由……发展而来"},
		},
		Synos: []Syno{
			{Pos: "v.", Cn: "发展；成长", Ws: []string{"expand", "grow", "advance", "improve"}},
		},
		RelWords: RelWords{
			Root: "velop",
			Rels: []RelWordPos{
				{
					Pos: "n.",
					Words: []RelWordSub{
						{C: "development", Cn: "发展；开发"},
						{C: "developer", Cn: "开发者；开发商"},
					},
				},
				{
					Pos: "adj.",
					Words: []RelWordSub{
						{C: "developed", Cn: "发达的；成熟的"},
						{C: "developing", Cn: "发展中的"},
					},
				},
			},
		},
		Etymology: []Etymology{
			{T: "from Old French ‘desveloper’", D: "意为“揭开，展开”，词根 velop 表示“包裹”"},
		},
	}

	jsonData, err := json.MarshalIndent(word, "", "  ")
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}

	fmt.Println(string(jsonData))
}
