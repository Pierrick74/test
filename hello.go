package main

import "fmt"

func main() {
    var name string
    fmt.Print("Entrez votre nom : ")
    fmt.Scanln(&name)
    fmt.Printf("Bonjour, %s ! Bienvenue dans Go.\n", name)
}