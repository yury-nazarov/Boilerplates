package main

import (
	"fmt"
	"os"
)

func main()  {
	filePrinter()
}

func filePrinter() {
	file, err := os.OpenFile("output.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Не удалосьсоздать файл", err)
	}

	words := map[int]string{
		1: "hello",
		2: "world",
	}

	for k, v := range words {
		_, err := fmt.Fprintf(file, "%d %s\n", k, v)
		if err != nil {
			fmt.Println("Не удалось записать строку в файл")
		}
	}
}
