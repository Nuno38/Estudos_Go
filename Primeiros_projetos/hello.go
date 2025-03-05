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

func main(){
    clearScreen()
     nome := "Elias"
     versao := 1.1
    fmt.Println("\nOlá sr.", nome,"")
    fmt.Println("Este programa está na versão", versao,"\n")
    
    fmt.Println("1. Iniciar Monitoramento")
    fmt.Println("2. Exibir Logs")
    fmt.Println("3. Sair do Programa")

    var comando int
    fmt.Scan(&comando)
    clearScreen()
    
    fmt.Printf("\nOpção escolhida foi %d",comando)
}
