package main

import (
	"fmt"
	"example/hello/util"
)

func main() {
	var n int
	fmt.Print("Combien de chiffres ? ")
	fmt.Scanln(&n)

	tab := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Chiffre %d : ", i+1)
		fmt.Scanln(&tab[i])
	}

	utils.TriBubble(tab)
	fmt.Println("Tableau trié :", tab)
}