package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWord(word string) string {
	reversed := ""
	
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i]) 
	}

	return reversed
}

func main() {
	fmt.Println("Welcome to Reverse String!")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Masukkan minimal 3 kata:")

		scanner.Scan()
		input := scanner.Text()
		words := strings.Fields(input)

		if len(words) < 3 {
			fmt.Println("Jumlah kata kurang dari 3. Coba lagi dari awal.")
			continue
		}

		fmt.Print("Output: ")
		for _, word := range words {
			fmt.Print(reverseWord(word) + " ")
		}
		fmt.Println()
		break
	}
}
