package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"net/http"
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

func mensagemInicial() {
	fmt.Println("1. Iniciar Monitoramento")
	fmt.Println("2. Exibir Logs")
	fmt.Println("3. Sair do Programa")
}

func versaoSistema() {
	nome := "Elias"
	versao := 1.1
	fmt.Println("\nOlá sr.", nome, "")
	fmt.Println("Este programa está na versão", versao, "\n\n")
}

func iniciarMonitoramento() {
	fmt.Println("\nIniciando monitoramento....")
	site := "https://httpbin.org/status/404"
	resp, _ := http.Get(site)
	if resp.StatusCode == 404 {
		fmt.Println("Erro ao acessar o site:", resp.Status)
		return
	}
	fmt.Println("Status do site:", resp.StatusCode)
}

func processarEscolha() {
	for {
		var comando string
		fmt.Scan(&comando)

		fmt.Println("\nOpção escolhida foi", comando)

		switch comando {
		case "1":
			clearScreen()
			iniciarMonitoramento()
            os.Exit(0)

		case "2":
			clearScreen()
			fmt.Println("\nExibindo logs...")
            os.Exit(0)
		case "3":
			clearScreen()
			fmt.Println("\nSaindo do sistema")
			os.Exit(0)
		default:
			clearScreen()
			fmt.Println("\nOpção incorreta, favor escolher uma opção correta...\n")
			mensagemInicial()
		}
	}
}

func main() {
	clearScreen()
	versaoSistema()
	mensagemInicial()
	processarEscolha()
}