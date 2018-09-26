package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"time"
	"github.com/gorilla/mux"
)

// Pontos Bla Bla Bla
type Pontos struct{
	Score int
	Nome string
}

// Word Bla Bla bla
type Word struct {
	Palavra string
	Achado  bool
}

// getWord le do banco de dados uma lista de Palavras
func getWord() []Word {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var listaWord = [][]Word{
		{{Palavra: "leite",},{Palavra: "elite",},{Palavra: "tele",},{Palavra: "til",},{Palavra: "lei",},{Palavra: "ele",}},
		{{Palavra: "manter",},{Palavra:"trena",},{Palavra:"menta",},{Palavra:"marte",},{Palavra:"trem",},{Palavra:"reta"},{Palavra:"meta",},{Palavra:"mar",},{Palavra:"ema",}},
		{{Palavra:"tropa",},{Palavra:"trapo",},{Palavra:"prato",},{Palavra:"tora",},{Palavra:"topa",},{Palavra:"rato",},{Palavra:"pra",},{Palavra:"por",},{Palavra:"aro",}},
		{{Palavra:"remoto",},{Palavra:"termo",},{Palavra:"temor",},{Palavra:"metro",},{Palavra:"toro",},{Palavra:"reto",},{Palavra:"remo",},{Palavra:"ter",},{Palavra:"mor",}},
		{{Palavra:"calhar",},{Palavra:"racha",},{Palavra:"achar",},{Palavra:"cara",},{Palavra:"acha",},{Palavra:"lar",},{Palavra:"cha",},{Palavra:"ala",}},
	}
	return listaWord[r1.Intn(len(listaWord))]
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
