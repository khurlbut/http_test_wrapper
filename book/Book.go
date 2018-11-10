package book

import (
	"encoding/json"
	"strings"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

const NOVEL = "NOVEL"
const SHORT_STORY = "SHORT STORY"

func NewBookFromJSON(js string) (Book, error) {
	var book Book
	err := json.Unmarshal([]byte(js), &book)
	return book, err
}

func (b Book) CategoryByLength() (cat string) {

	if b.Pages > 300 {
		cat = NOVEL
	} else {
		cat = SHORT_STORY
	}

	return cat
}

func (b Book) AuthorLastName() (lastname string) {
	arr := strings.Split(b.Author, " ")
	return arr[len(arr)-1]
}
