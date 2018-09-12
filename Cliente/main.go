package main

import "fmt"

type Words struct {
	wd [10]string
	boll [10]string
}

//Checar se a palavra digitada existe dentro da palavra inicial

func check(){
	var palavra string

	for palavra != "q"{
		fmt.Scanf("%s\n", &palavra)
		fmt.Println("\x1b[37;1m")
		fmt.Println("\u001b[2J")
		fmt.Println("print", palavra)
		
	}
}

func main(){
	//fmt.Println("ola")
	check()
}