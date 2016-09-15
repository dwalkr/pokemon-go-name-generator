package main

import (
	"bufio"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/codegangsta/negroni"
)

var firstPart []string
var lastPart []string

func loadData() {
	first, _ := getNames("data/first")
	last, _ := getNames("data/last")
	either, _ := getNames("data/either")

	firstPart = append(first, either...)
	lastPart = append(last, either...)
}

func getNames(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := strings.Trim(scanner.Text(), " ")
		if utf8.RuneCountInString(lineText) > 0 {
			lines = append(lines, lineText)
		}
	}
	return lines, scanner.Err()
}

func GenerateName(maxLength int) string {

	finalFirst := strings.Trim(firstPart[rand.Intn(len(firstPart))], " ")
	finalLast := strings.Trim(lastPart[rand.Intn(len(lastPart))], " ")

	finalName := finalFirst + finalLast

	if maxLength > 0 && len(finalName) > maxLength {
		return GenerateName(maxLength)
	}

	return finalName
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(GenerateName(12)))
}

func main() {
	loadData()
	rand.Seed(time.Now().UnixNano())
	mux := http.NewServeMux()
	mux.HandleFunc("/generate", NameHandler)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
