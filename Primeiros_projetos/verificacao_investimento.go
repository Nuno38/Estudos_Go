package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)


func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Struct para armazenar a cotação da API
type Cotacao struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func mensagemInicio() {
	fmt.Println("\nBem-vindo ao sistema de cotação do dia!")
	time.Sleep(3 * time.Second)
	clearScreen()
}


func insecaoDeValores() (string, float64) {
	var nomeFundo string
	var valorFundo float64

	fmt.Print("Insira o nome do Fundo: ")
	fmt.Scanln(&nomeFundo)

	fmt.Print("Insira o valor do Fundo: ")
	fmt.Scanln(&valorFundo)

	return nomeFundo, valorFundo
}

// Função para buscar a cotação do dólar
func cotacaoDoDia() {
	site := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	// Fazendo a requisição HTTP
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer resp.Body.Close() 


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return
	}


	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}


	fmt.Println("Cotação do dólar hoje: R$", cotacao.USDBRL.Bid)
}

func main() {
	mensagemInicio()
	cotacaoDoDia()
	
}
