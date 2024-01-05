package main

// Чтение из файла

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

func main() {
	ParseFilePerString()
}


type item struct {
	ID 			int
	Category 	int
	Key 		string
	Value 		string
}

// ParseFilePerString простейший способ распарсить файл
func ParseFilePerString() {
	// Открываем файл
	// file - содержит ссылку на *File. Пример:  &{0xc0000ae120}
	file, err := os.OpenFile("input.txt", os.O_RDWR, 0755)
	if err != nil {
		fmt.Println("can not open file", err)
	}

	// Создаем новый сканер, принимающий io.Reader, которому соответствует *File
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Объект item в который будем записиывать каждую строчку
		var i item

		// line - содержит строку из файла
		line := scanner.Text()

		// Парсим строку, извлекая значение каждого элемента разделенного табом
		fmt.Sscanf(
			line,
			"%d\t%d\t%s\t%s",
			&i.ID, &i.Category, &i.Key, &i.Value)

		// Если элемента не существует, вернет пустое значение
		// 			Дополнительно добоально не плохой пример: https://stackoverflow.com/questions/20921619/is-there-any-example-and-usage-of-url-queryescape-for-golang
		uv, err := url.QueryUnescape(i.Value)
		if err != nil {
			fmt.Println("Ошибка парсинга строки файла", err)
		}
		i.Value = uv

		fmt.Println(i)
	}

}
