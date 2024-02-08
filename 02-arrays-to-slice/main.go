package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [100]int
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	for i := range arr {
		value := rnd.Intn(20)
		arr[i] = value
	}
	var slice [20]int
	for _, v := range arr {
		slice[v] += 1
	}
	// Создайте слайс нужного размера
	// slice :=

	// Заполните слайс количеством встречаемых чисел
	fmt.Println(arr)
	fmt.Println(slice)
}
