package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
func New(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "new!")
}
func Score(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "score!")
}
func Save(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "save")
}
