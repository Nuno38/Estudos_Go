package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"strings"
)

func limpaTela(){
	var cmd *exec.Cmd
	cmd2 := exec.Command("nircmd.exe", "monitor", "off")
	switch runtime.GOOS{
	case "windows":
		cmd = exec.Command("cmd","/c","cls")
		cmd2 = exec.Command("nircmd.exe", "monitor", "off")

	default:
		cmd =exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd2.Stdout = os.Stdout
	cmd.Run()
	cmd2.Run()
}

func telaInicial(){
	fmt.Printf("\nBem vindo ao supermercado Virtual escolha um produto❗")
	time.Sleep(3*time.Second)
	limpaTela()
}

func apresentacaoProdutos(){
	fmt.Printf("\nLista de produtos")
	fmt.Printf("\n1. Café R$ 26,50")
	fmt.Printf("\n2. Ovo R$ 28,50")
	fmt.Printf("\n3. Filé mignon R$ 58,00/Kg")
	fmt.Printf("\n4. Carne moída R$ 34,98/kg")
	fmt.Printf("\n5. Peito de frango R$ 15,60/KG\n")

	fmt.Printf("\n\nRealize sua escolha:")
}

func escolha_Realizada(){

	for{
	apresentacaoProdutos()

	var escolha string
	
    fmt.Scan(&escolha)
	escolha = strings.TrimSpace(escolha)
	escolha = strings.ToLower(escolha) 
    limpaTela()

	switch escolha {
	case "1":
		fmt.Printf("\nVocê escolheu um 1kg de café. ")
		return
	case "2":
		fmt.Printf("\nVocê escolheu uma bandeja com 30 ovos.")
		return
	case "3":
		fmt.Printf("\nVocê escolheu 1kg de Filé mignon.")
		return
	case "4":
		fmt.Printf("\nVocê escolheu 1kg de Carne Moída.")
		return
	case "5":
		fmt.Printf("\nVocê escolheu 1kg de Peito de frango.")
		return
    case "tão","porque","puta","jesus","vix","ai","ta","loco","nem","fudendo","está","muito","mano","que","caro":{
		fmt.Printf("\nFaz o L ai parceiro!\n")
		return
    }

	default:
		fmt.Printf("Escolha incorreta!\n\n escolha novamente um número na lista apresentada!")
		time.Sleep(3*time.Second)
		
	}
 }
}
func main(){
	limpaTela()
	telaInicial()
	escolha_Realizada()
}