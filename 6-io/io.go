package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Пакет io просто пакет для чтения файлов с помощью интерфейсов и удобного управления чтением из дескрипторов
	bytes, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(len(bytes))

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	var closer io.ReadCloser = file // потому что file соотвествует интерфейсу ReadCloser

	fileStat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if fileStat.IsDir() {
		fmt.Println("this is directory")
	} else {
		fmt.Println("this is file")
	}

	fmt.Printf("last mod time: %s\n", fileStat.ModTime())

	buffer := make([]byte, 8, 8)

	count, err := closer.Read(buffer)

	fmt.Println(count)
	fmt.Println(err)
	fmt.Println(buffer)
	fmt.Println(string(buffer))

	closer.Close() // закрываем дескриптор
}
