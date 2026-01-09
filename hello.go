package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Как тебя зовут?")

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text) // Убираем лишние пробелы и символы переноса

	fmt.Printf("Привет, %s! Это твое первое приложение на Go.\n", text)

	fmt.Println("Нажми Enter, чтобы выйти...")
	fmt.Scanln()
}
