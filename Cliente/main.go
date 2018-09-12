package main

import (
	"fmt"
)

// Word contem a palavra e o estado se ela foi encontrada ou não
type Word struct {
	palavra string
	achado  bool
}

var score int

// Pegar palavra do json
func getWord() []Word {
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

	return listaWord
}

func printa(listaPalavra []Word) {

	fmt.Printf("\u001b[41m")
	for i := 0; i < len(listaPalavra); i++ {
		if listaPalavra[i].achado == true {
			fmt.Printf("\x1b[37;1m")
			fmt.Printf("\u001b[42m")
			fmt.Println(listaPalavra[i].palavra)
			fmt.Printf("\u001b[0m")
		} else {
			fmt.Printf("\x1b[37;1m")
			fmt.Printf("\u001b[41m")
			for j := len(listaPalavra[i].palavra); j > 0; j-- {
				fmt.Printf("#")
			}
			fmt.Println()
		}
	}
	fmt.Println("\u001b[0m")
}

//Checar se a palavra digitada existe dentro da palavra inicial
func check(listaPalavra []Word, sugestao string) {
	for i := 0; i < len(listaPalavra); i++ {
		if sugestao == listaPalavra[i].palavra {
			listaPalavra[i].achado = true
			break
		}
	}
	fmt.Println("Palavra nao encontrada")
}

func play() {

	var listaPalavra []Word = getWord()
	var sugestao string

	printa(listaPalavra)

	for sugestao != "/q" {
		fmt.Println("Digite uma palavra")
		fmt.Scanf("%s\n", &sugestao)
		check(listaPalavra, sugestao)

		fmt.Println("\u001b[2J")

		printa(listaPalavra)

	}
}

// /q sai do jogo
// /new uma palavra nova
// /score retorna o score

func main() {
	var menu string

	fmt.Println("Menu:\n Para começar um jogo digite /new \n Para sair /q")
	fmt.Scanf("%s\n", &menu)

	switch menu {
	case "/new":
		play()
	case "/score":
		fmt.Println("nao ta feito ainda...")
	}

}
