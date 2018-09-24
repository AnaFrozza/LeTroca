package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
)

// Word contem a palavra e o estado se ela foi encontrada ou não
type Word struct {
	Palavra string
	Achado  bool 
}


var score int

// Pegar palavra do json
func getWord() []Word{
	
	web, err := http.Get("http://127.0.0.1:8080/new")

	if(err != nil){
		fmt.Print(err.Error())
		os.Exit(1)
	}

	webData, err := ioutil.ReadAll(web.Body)
	if(err != nil){
		log.Fatal(err)
	}

	var	listaWord []Word
	json.Unmarshal(webData, &listaWord)
	fmt.Println(listaWord)
	return listaWord
}

// printa as palavras na tela
func printa(listaPalavra []Word) {
	palavraMisturada := randomizeWord(listaPalavra[0].Palavra)
	fmt.Println("\n", palavraMisturada)
	fmt.Println("\u001b[40;1m") //fundo preto
	fmt.Printf("\u001b[31;1m")  // letra vermelha
	for i := 0; i < len(listaPalavra); i++ {
		if listaPalavra[i].Achado == true {
			fmt.Printf("\x1b[37;1m") //negrito
			fmt.Printf("\u001b[42m") //fundo verde
			fmt.Println(listaPalavra[i].Palavra)
			fmt.Printf("\u001b[0m") // reset
		} else {
			fmt.Printf("\x1b[37;1m")   // negrito
			fmt.Printf("\u001b[31;1m") // letra vermelha
			fmt.Printf("\u001b[40;1m") // fundo preto
			for j := len(listaPalavra[i].Palavra); j > 0; j-- {
				fmt.Printf("#")
			}
			fmt.Println()
		}
	}
	fmt.Printf("\u001b[0m") // reset
}

// Checar se a palavra digitada existe dentro da palavra inicial
func check(listaPalavra []Word, sugestao string) string {
	if sugestao == "/q" {
		return "sair"
	}

	for i := 0; i < len(listaPalavra); i++ {
		if sugestao == listaPalavra[i].Palavra {
			if listaPalavra[i].Achado == true {
				return "duplicado"
			}
			listaPalavra[i].Achado = true
			return "achei"
		}
	}
	return "naoEncontrado"
}

// change recebe uma palavra e altera a ordem das letras
func change(palavraOrig string) string {
	palavra := []byte(palavraOrig)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for k := r1.Intn(600); k > 0; k-- {
		for i := 0; i < len(palavra); i++ {
			for j := 0; j < len(palavra); j++ {
				if i != j {
					//fmt.Println(i, j, k)
					rand.Shuffle(len(palavra), func(i, j int) {
						palavra[i], palavra[j] = palavra[j], palavra[i]
					})
				}
			}
		}
	}
	str := string(palavra[:])
	//fmt.Println(str)
	return str

}

// randomizeWord recebe uma palavra e garante que a funcao change retorne uma palavra diferente da original
func randomizeWord(palavra string) string {
	str := palavra
	for str == palavra {
		str = change(str)
	}
	return str
}

func play() {

	var listaPalavra []Word = getWord()
	var sugestao string
	var retorno string
	var word string = randomizeWord(listaPalavra[0].Palavra) // Checar se vai printar aqui ou na print!!!

	printa(listaPalavra)

	for sugestao != "/q" {
		// fmt.Printf("\u001b#3 \u001b#4") // letra grande
		fmt.Println(word)
		// fmt.Printf("\u001b[0m") // reset
		fmt.Println("Digite uma palavra")
		fmt.Scanf("%s\n", &sugestao)

		fmt.Printf("\u001b[2J") // limpa a tela

		retorno = check(listaPalavra, sugestao)

		switch retorno {
		case "naoEncontrado":
			fmt.Printf("\u001b[43m")   // fundo amarelo
			fmt.Printf("\u001b[30;1m") // letra petra
			fmt.Println("-Palavra não encontrada-")
			fmt.Printf("\u001b[0m") // reset
		case "duplicado":
			fmt.Printf("\u001b[43m")   // fundo amarelo
			fmt.Printf("\u001b[30;1m") // letra preta
			fmt.Println("-Você já encontrou essa palavra-")
			fmt.Printf("\u001b[0m") // reset
		}

		printa(listaPalavra)

	}

	fmt.Printf("\u001b[2J") // limpa a tela

	for i := 0; i < len(listaPalavra); i++ {
		fmt.Printf("\x1b[37;1m") //negrito
		fmt.Printf("\u001b[42m") //fundo verde
		fmt.Println(listaPalavra[i].Palavra)
		fmt.Printf("\u001b[0m") // reset
	}
}

func main() {
	var menu string
	fmt.Printf("\u001b[0m") // reset
	fmt.Println("Menu:\n Para começar um jogo digite /new \n Para sair /q")
	fmt.Scanf("%s\n", &menu)

	switch menu {
	case "/new":
		play()
	case "/score":
		fmt.Println("nao ta feito ainda...")
	}
}
