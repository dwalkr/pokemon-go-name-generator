package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var firstPart []string
var lastPart []string

var homeTemplate = template.Must(template.ParseFiles("view/main.tpl", "view/index.tpl"))
var lipsumTemplate = template.Must(template.ParseFiles("view/main.tpl", "view/lipsum.tpl"))

type TextResponse struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

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

func GenerateNameWrapper() string {
	return GenerateName(100)
}

func GenerateLipsum(numParagraphs uint8) []string {
	return lipsumGenerateParagraphs(numParagraphs, GenerateNameWrapper)
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var resp = &TextResponse{}
	numNames, err := strconv.ParseUint(vars["quantity"], 10, 8)
	if err != nil || numNames < 1 {
		numNames = 1
	}
	var names []string
	for i := 0; i < int(numNames); i++ {
		names = append(names, GenerateName(12))
	}
	if r.FormValue("nameFormat") == "capitalize" {
		names = ucwords(names)
	}
	resp.Status = "success"
	resp.Data = names
	sendResponse(w, r, resp)
}

func sendResponse(w http.ResponseWriter, r *http.Request, resp *TextResponse) {
	responseType := r.FormValue("format")
	if responseType == "text" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(strings.Join(resp.Data, " ")))
		return
	}
	js, err := json.Marshal(resp)
	if err != nil {
		w.Header().Set("Status", "500 Internal Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ucwords(words []string) []string {
	for i := 0; i < len(words); i++ {
		words[i] = ucfirst(words[i])
	}
	return words
}

func ucfirst(word string) string {
	newWord := strings.ToUpper(string(word[0]))
	newWord += word[1:]
	return newWord
}

func LipsumHander(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var resp = &TextResponse{}
	paragraphs, err := strconv.ParseUint(vars["paragraphs"], 10, 8)
	if err != nil || paragraphs < 1 {
		paragraphs = 3
	}
	resp.Status = "success"
	resp.Data = GenerateLipsum(uint8(paragraphs))
	sendResponse(w, r, resp)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, nil)
}
func LipsumFrontHandler(w http.ResponseWriter, r *http.Request) {
	lipsumTemplate.Execute(w, nil)
}

func main() {
	// TODO: Lorem Ipsum
	loadData()
	rand.Seed(time.Now().UnixNano())

	router := mux.NewRouter()

	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/lipsum", LipsumFrontHandler)

	router.HandleFunc("/generate", NameHandler)
	router.HandleFunc("/generate/{quantity:[0-9]+}", NameHandler)
	router.HandleFunc("/generate/lipsum", LipsumHander)
	router.HandleFunc("/generate/lipsum/{paragraphs:[0-9]+}", LipsumHander)

	mux := http.NewServeMux()
	mux.Handle("/", router)
	n := negroni.Classic()
	n.UseHandler(mux)

	port := os.Getenv("PORT_NAMEGEN")
	n.Run(":" + port)
}
