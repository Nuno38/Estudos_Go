package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

// Função para limpar a tela
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


func cotacaoDoDia() (float64, error) {
	site := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	resp, err := http.Get(site)
	if err != nil {
		return 0, fmt.Errorf("erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("erro ao ler a resposta: %v", err)
	}

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		return 0, fmt.Errorf("erro ao decodificar JSON: %v", err)
	}

	
	valorDolar, err := strconv.ParseFloat(cotacao.USDBRL.Bid, 64)
	if err != nil {
		return 0, fmt.Errorf("erro ao converter a cotação: %v", err)
	}

	return valorDolar, nil
}

// Função para verificar o investimento e calcular o rendimento
func verificaInvestimento() {
	cotacao, err := cotacaoDoDia()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Cotação do dólar hoje: R$", cotacao)

	var valorInvestido float64

	fmt.Print("Informe o valor investido em reais: ")
	_, err = fmt.Scanln(&valorInvestido)
	if err != nil {
		fmt.Println("Erro ao ler o valor investido:", err)
		return
	}

	// Calcula o valor em dólares e o rendimento
	valorEmDolares := valorInvestido / cotacao

	fmt.Printf("\nCom um investimento de R$ %.2f e a cotação de R$ %.2f\n", valorInvestido, cotacao)
	fmt.Printf("Você teria aproximadamente USD %.2f\n", valorEmDolares)
}

func main() {
	mensagemInicio()
	verificaInvestimento()
}
