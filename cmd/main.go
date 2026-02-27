package main

import (
	"fmt"
	"os"
	a "shortify/internal/app"
)

func main() {
	err := a.App()
	if err != nil {
		fmt.Printf("Ошибка, закрытие приложения: %v", err)
		os.Exit(1) 
	}
}