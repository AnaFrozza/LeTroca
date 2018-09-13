package main

import (
	"fmt"
	// "math/rand"
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


// printa as palavras na tela
func printa(listaPalavra []Word) {

	fmt.Println("\u001b[40;1m") //fundo preto
	fmt.Printf("\u001b[31;1m") // letra vermelha
	for i := 0; i < len(listaPalavra); i++ {
		if listaPalavra[i].achado == true {
			fmt.Printf("\x1b[37;1m") //negrito
			fmt.Printf("\u001b[42m") //fundo verde
			fmt.Println(listaPalavra[i].palavra)
			fmt.Printf("\u001b[0m") // reset
		} else {
			fmt.Printf("\x1b[37;1m") // negrito
			fmt.Printf("\u001b[31;1m") // letra vermelha
			fmt.Printf("\u001b[40;1m") // fundo preto
			for j := len(listaPalavra[i].palavra); j > 0; j-- {
				fmt.Printf("#")
			}
			fmt.Println()
		}
	}
	fmt.Printf("\u001b[0m") // reset
}

//Checar se a palavra digitada existe dentro da palavra inicial
func check(listaPalavra []Word, sugestao string) string{
	if sugestao == "/q" {
		return "sair"
	}

	for i := 0; i < len(listaPalavra); i++ {
		if sugestao == listaPalavra[i].palavra {
			if listaPalavra[i].achado == true {
				return "duplicado"
			}
			listaPalavra[i].achado = true
			return "achei"
		}
	}
	return "naoEncontrado"
}

func randomWord(palavra string) string{
	
 	return palavra
}

func play() {

	var listaPalavra []Word = getWord()
	var sugestao string
	var retorno string
	var word string = randomWord(listaPalavra[0].palavra)

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
			fmt.Printf("\u001b[43m") // fundo amarelo
			fmt.Printf("\u001b[30;1m") // letra petra
			fmt.Println("-Palavra não encontrada-")
			fmt.Printf("\u001b[0m") // reset
		case "duplicado":
			fmt.Printf("\u001b[43m") // fundo amarelo
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
		fmt.Println(listaPalavra[i].palavra)
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
