package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Word Bla Bla bla
type Word struct {
	Palavra string
	achado  bool
}

// getWord le do banco de dados uma lista de Palavras
func getWord() []Word {
	var listaWord = []Word{
		{
			Palavra: "leite",
		},
		{
			Palavra: "elite",
		},
		{
			Palavra: "tele",
		},
		{
			Palavra: "til",
		},
		{
			Palavra: "lei",
		},
		{
			Palavra: "ele",
		},
	}
	return listaWord
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/new", New)
	router.HandleFunc("/score", Score)
	router.HandleFunc("/scoresave", Save)
	//log em caso de erros
	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index bla bla bla
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}

// New descrever
func New(w http.ResponseWriter, r *http.Request) {
	var listaPalavra = getWord()
	json.NewEncoder(w).Encode(listaPalavra)

	//fmt.Fprintln(w, listaPalavra)
}

// Score descrever
func Score(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "score!")
}

// Save descrever
func Save(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "save")
}
