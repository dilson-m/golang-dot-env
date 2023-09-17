package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	envErr := godotenv.Load(".env")

	if envErr != nil {
		fmt.Println("Erro ao carregar arquivo de configuracao.")
		return
	}

	envMap, err := godotenv.Read(".env")

	if err != nil {
		fmt.Println("Erro ao ler variaveis de ambiente de .ENV")
		return
	}

	fmt.Printf("%s", envMap["MYMESSAGE"])

	// carregar exportacao do git
	// go get github.com/joho/godotenv

}
