package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Word contem a palavra e o estado se ela foi encontrada ou não
type Word struct {
	palavra string
	achado  bool
}

// Adicionando palavras na lista de palavras para teste
// 5 palavras
var listaWord = []Word{
	Word{
		palavra: "leite",
	},
	Word{
		palavra: "elite",
	},
	Word{
		palavra: "tele",
	},
	Word{
		palavra: "til",
	},
	Word{
		palavra: "lei",
	},
	Word{
		palavra: "ele",
	},
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

// Index apresenta a função inicial
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// New retorna uma nova palavra do banco de dados para o cliente
func New(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "new!")
}

// Score retorna os valores do score ao cliente
func Score(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "score!")
}

// Save permite adicionar uma entrada aos valores de score
func Save(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "save")
}
