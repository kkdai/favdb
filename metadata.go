package favdb

import (
	"encoding/json"
	"log"
)

type Model struct {
	Db  UserFavData
	Log *log.Logger
}

type MessageCount struct {
	All     int `json:"all"`
	Boo     int `json:"boo"`
	Count   int `json:"count"`
	Neutral int `json:"neutral"`
	Push    int `json:"push"`
}

type ArticleDocument struct {
	ArticleID    string        `json:"article_id"`
	ArticleTitle string        `json:"article_title"`
	Author       string        `json:"author"`
	Board        string        `json:"board"`
	Content      string        `json:"content"`
	Date         string        `json:"date"`
	IP           string        `json:"ip"`
	MessageCount MessageCount  `bson:"message_count"`
	Messages     []interface{} `json:"messages"`
	Timestamp    int           `json:"timestamp"`
	URL          string        `json:"url"`
	ImageLinks   []string      `json:"image_links"`
}

func (d *ArticleDocument) ToString() (info string) {
	b, err := json.Marshal(d)
	if err != nil {
		//fmt.Println(err)
		return
	}
	return string(b)
}

// ArticleDocument for sorting.
type AllArticles []ArticleDocument

func (a AllArticles) Len() int           { return len(a) }
func (a AllArticles) Less(i, j int) bool { return a[i].MessageCount.Count > a[j].MessageCount.Count }
func (a AllArticles) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
