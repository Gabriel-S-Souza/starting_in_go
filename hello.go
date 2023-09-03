package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	initialise()
	for {
		showMenu()
		command := getCommand()

		if command == 1 {
			path := os.Args[1]
			monitor(path)
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
}

func showMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Sair do programa")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func monitor(filePath string) {
	sites := getWebsites(filePath)

	for i := 0; i < len(sites); i++ {
		fmt.Println("Monitorando...", sites[i])
		resp, _ := http.Get(sites[i])
		isSuccess := resp.StatusCode >= 200 && resp.StatusCode <= 299

		if isSuccess {
			fmt.Println("Site:", sites[i], "está online. Status Code:", resp.StatusCode)
		} else {
			fmt.Println("Site:", sites[i], "está com problemas. Status Code:", resp.StatusCode)
		}
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
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	reader := bufio.NewReader(file)

	var sites []string

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		sites = append(sites, strings.TrimSpace(line))
	}

	file.Close()
	return sites
}
