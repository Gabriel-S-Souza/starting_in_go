package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

var filePath string

func main() {
	initialise()
	for {
		showMenu()
		command := getCommand()
		if command == 1 {
			monitor()
		} else if command == 2 {
			exit()
		} else {
			handleUnknownCommand()
		}
		time.Sleep(1 * time.Second)
	}

}

func initialise() {
	fmt.Println("Iniciando...")

	if len(os.Args) > 1 {
		filePath = os.Args[1]
	} else {
		fmt.Println("Informe o caminho do arquivo:")
		fmt.Scan(&filePath)
	}
}

func showMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Sair do programa")
	fmt.Println("")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func monitor() {
	sites := getWebsites(filePath)

	for _, site := range sites {
		fmt.Println("Monitorando...", site)
		resp, err := http.Get(site)
		isSuccess := resp.StatusCode >= 200 && resp.StatusCode <= 299

		if isSuccess {
			fmt.Println("Site:", site, "está online. Status Code:", resp.StatusCode)
		} else {
			fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		}

		if err != nil {
			fmt.Println("Ocorreu um erro:", err)
		}

		fmt.Println("")
	}
}

func exit() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}

func handleUnknownCommand() {
	fmt.Println("Não conheço este comando")
	os.Exit(-1)
}

func getWebsites(filePath string) []string {
	fileBytes, err := os.ReadFile(filePath)

	if err != nil {
		var errorText string
		if os.IsNotExist(err) {
			errorText = "Ocorreu um erro: o arquivo " + filePath + " não existe"
		} else {
			errorText = "Ocorreu um erro: " + err.Error()
		}
		fmt.Println("Ocorreu um erro:", errorText)
	}

	fileString := strings.TrimSpace(string(fileBytes))
	sites := strings.Split(fileString, "\n")

	return sites
}
