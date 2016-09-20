package main

import (
	"math/rand"
	"strings"
)

func lipsumGenerateParagraphs(numParagraphs uint8, getWord func() string) []string {
	var paragraphs []string
	n := int(numParagraphs)
	for i := 0; i < n; i++ {
		paragraph := ""
		numSentences := rand.Intn(6) + 4
		for k := 1; k <= numSentences; k++ {
			numWords := rand.Intn(6) + 2
			sentence := ""
			commaPlacement := 0
			if numWords > 3 {
				if rand.Intn(7) >= 4 {
					commaPlacement = rand.Intn(numWords-2) + 1
				}
			}

			for q := 1; q <= numWords; q++ {
				sentence += getWord()
				if q == 1 {
					//capitalize first word
					newSentence := strings.ToUpper(string(sentence[0]))
					newSentence += sentence[1:]
					sentence = newSentence
				}
				if q == commaPlacement {
					//insert random comma
					sentence += ","
				}
				if q < numWords {
					//add space after all words except last one
					sentence += " "
				}
			}
			sentence += "."
			paragraph += sentence
			if k < numSentences {
				paragraph += " "
			}
		}
		paragraphs = append(paragraphs, paragraph)
	}
	return paragraphs
}
