package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoring = 3
const delay = 5

func main() {
	intro()
	menu()

	command := seeCommand()

	switch command {
	case 1:
		initMonitoring()
	case 2:
		fmt.Println("Displaying logs")
	case 3:
		fmt.Println("In development..")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid command")
		os.Exit(-1)
	}
}

func intro() {
	fmt.Println("Hello. What's your name?")
	var name string
	fmt.Scan(&name)

	fmt.Println("Hello,", name, "!")
}

func menu() {
	fmt.Println("CHOOSE AN OPTION :D")
	fmt.Println("1 - Init monitoring")
	fmt.Println("2 - See logs")
	fmt.Println("3 - Init InfoMoney")
	fmt.Println("0 - Exit")
}

func seeCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("Your choice:", command)

	return command
}

func initMonitoring() {
	fmt.Println("Start monitoring...")
	site := "https://economia.awesomeapi.com.br/last/USD-BRL"
	resp, erro := http.Get(site)

	if erro != nil {
		panic(erro)
	}

	responseJson, erro := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		fmt.Println("Name:", string(responseJson))
		testSite(site)

		for i := 0; i < monitoring; i++ {
			time.Sleep(time.Second * delay)
			testSite(site)
		}

		time.Sleep(delay * time.Second)

	} else {
		fmt.Println("Not good! :(")
	}
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}
