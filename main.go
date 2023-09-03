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
			showLogs()
		} else if command == 3 {
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
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Sair do programa")
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
		fmt.Println("Testando...", site)
		resp, err := http.Get(site)
		isSuccess := resp.StatusCode >= 200 && resp.StatusCode <= 299

		if isSuccess {
			fmt.Println("Site:", site, "está online. Status Code:", resp.StatusCode)
		} else {
			fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		}

		registerLog(site, isSuccess)

		if err != nil {
			fmt.Println("Ocorreu um erro:", err)
		}

		fmt.Println("")
	}
}

func showLogs() {
	fmt.Println("Exibindo logs...")
	fileBytes, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(fileBytes))
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

func registerLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	var statusText string
	if status {
		statusText = "online"
	} else {
		statusText = "offline"
	}

	now := time.Now().Format("02/01/2006 15:04:05")

	file.WriteString(now + " - " + site + " - " + statusText + "\n")
	file.Close()
}
