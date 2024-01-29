package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	var fileName, fileExt string
	last := strings.LastIndex(filePth, "/")
	fileName = filePth[last+1:]
	last = strings.LastIndex(fileName, ".")
	if last != -1 {
		fileExt = fileName[last:]
	}

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
