package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Ponto guarda a pontuação do jogador e o nome para ser adicionado a uma array com o score
type Ponto struct {
	Score int
	Nome  string
}

// Word contem a palavra e o estado se ela foi encontrada ou não
type Word struct {
	Palavra string
	Achado  bool
}

// getWord le do banco de dados uma lista de Palavras
func getWord() []Word {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var listaWord = [][]Word{
		{{Palavra: "leite"}, {Palavra: "elite"}, {Palavra: "tele"}, {Palavra: "til"}, {Palavra: "lei"}, {Palavra: "ele"}},
		{{Palavra: "manter"}, {Palavra: "trena"}, {Palavra: "menta"}, {Palavra: "marte"}, {Palavra: "trem"}, {Palavra: "reta"}, {Palavra: "meta"}, {Palavra: "mar"}, {Palavra: "ema"}},
		{{Palavra: "tropa"}, {Palavra: "trapo"}, {Palavra: "prato"}, {Palavra: "tora"}, {Palavra: "topa"}, {Palavra: "rato"}, {Palavra: "pra"}, {Palavra: "por"}, {Palavra: "aro"}},
		{{Palavra: "remoto"}, {Palavra: "termo"}, {Palavra: "temor"}, {Palavra: "metro"}, {Palavra: "toro"}, {Palavra: "reto"}, {Palavra: "remo"}, {Palavra: "ter"}, {Palavra: "mor"}},
		{{Palavra: "calhar"}, {Palavra: "racha"}, {Palavra: "achar"}, {Palavra: "cara"}, {Palavra: "acha"}, {Palavra: "lar"}, {Palavra: "cha"}, {Palavra: "ala"}},
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

//Index rota para a pagina inicial do site
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}

// New retorna um array do tipo Word com as palavras para o jogo
func New(w http.ResponseWriter, r *http.Request) {
	var listaPalavra = getWord()
	json.NewEncoder(w).Encode(listaPalavra)

	//fmt.Fprintln(w, listaPalavra)
}

// Score retorna uma pagina com a array das pontuações registradas
func Score(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "score!")
}

// Save retorna uma pagina para receber a pontuação do cliente
func Save(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "save")
}
