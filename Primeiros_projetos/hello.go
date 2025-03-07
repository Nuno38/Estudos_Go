package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"	
)

func clearScreen(){
    var cmd *exec.Cmd
    switch runtime.GOOS{
    case"windows":
        cmd = exec.Command("cmd","/c","cls")
    default:
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func mensagemInicial(){
    fmt.Println("1. Iniciar Monitoramento")
    fmt.Println("2. Exibir Logs")
    fmt.Println("3. Sair do Programa")
}

func versaoSistema(){
    nome := "Elias"
    versao := 1.1
   fmt.Println("\nOlá sr.", nome,"")
   fmt.Println("Este programa está na versão", versao,"\n")
}

func escolhaInserida(){
    for{    
       
     var comando string
     fmt.Scan(&comando)
    
    fmt.Printf("\nOpção escolhida foi %d",comando)
       
    switch comando{
    case "1":
        clearScreen()
        fmt.Printf("\nIniciando monitoramento....")
        os.Exit(0)
     case"2":
        clearScreen()
        fmt.Printf("\nExibindo logs..")
        os.Exit(0)
     case "3":
        clearScreen()
        fmt.Printf("\nSaindo do sistema")
        os.Exit(0)
     default:
        clearScreen()
        fmt.Printf("\nOpção incorreta, favor escolher uma opção correta...\n")
        mensagemInicial()
        }
    } 
}
func main(){
    clearScreen()
    versaoSistema()
    mensagemInicial()
    escolhaInserida()
}
