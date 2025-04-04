package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/asticode/go-astisub"
)

type Sentence struct {
	Content string
	StartAt time.Duration
	EndAt   time.Duration
}

func (s Sentence) String() string {
	return s.Content
}

func main() {
	s, err := astisub.OpenFile("/Users/yanbo/Movies/AlgoExpert/Two Number Sum [229483089].en.vtt")
	if err != nil {
		log.Println(err)
		return
	}

	var sentences []Sentence
	var draftSentence *Sentence
	for _, item := range s.Items {
		if draftSentence == nil {
			draftSentence = &Sentence{StartAt: item.StartAt}
		}

		draftSentence.Content += item.String()
		if strings.HasSuffix(draftSentence.Content, ".") {
			draftSentence.EndAt = item.EndAt
			sentences = append(sentences, *draftSentence)
			draftSentence = nil
		}

		// fmt.Println(item.StartAt, item.EndAt, item)
	}

	// return
	for _, item := range sentences {
		fmt.Println(item.StartAt, item.EndAt, item)
	}
}
