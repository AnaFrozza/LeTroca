package main

import "fmt"

type Word struct {
	palavra string
	achado bool
}

// Adicionando palavras na lista de palavras para teste
// 5 palavras
var listaWord = []Word{
	Word {
		palavra: "leite",
	},
	Word {
		palavra: "elite",
	},
	Word {
		palavra: "tele",
	},
	Word {
		palavra: "til",
	},
	Word {
		palavra: "lei",
	},
	Word {
		palavra: "ele",
	},
	
}

// Compara se a palavra digitada se encontra na lista de palavra
func comparaWord(sugestao string){

}

//Checar se a palavra digitada existe dentro da palavra inicial
func check(){
	var sugestao string

	for sugestao != "q"{
		fmt.Scanf("%s\n", &sugestao)
		fmt.Println("\x1b[37;1m")
		fmt.Println("\u001b[2J")
		fmt.Println("print", sugestao)
		
	}
}

func main(){
	fmt.Println(listaWord)
	
	check()
}