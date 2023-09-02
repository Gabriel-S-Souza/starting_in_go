package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	initialise()
	for {
		showMenu()
		command := getCommand()

		if command == 1 {
			monitor()
		} else if command == 2 {
			showLogs()
		} else if command == 3 {
			exit()
		} else {
			handleUnknownCommand()
		}
		time.Sleep(2 * time.Second)
	}

}

func initialise() {
	fmt.Println("Iniciando...")
}

func showMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Sair do programa")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func monitor() {
	site := "https://alura.com.br/"
	fmt.Println("Monitorando...", site)
	resp, _ := http.Get(site)
	isSuccess := resp.StatusCode >= 200 && resp.StatusCode <= 299

	if isSuccess {
		fmt.Println("Site:", site, "está online. Status Code:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func showLogs() {
	fmt.Println("Exibindo logs...")
}

func exit() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}

func handleUnknownCommand() {
	fmt.Println("Não conheço este comando")
	os.Exit(-1)
}
