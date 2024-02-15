package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Item struct {
	Name string
	Date time.Time
	Tags string
	Link string
}

func main() {
	var item []Item

	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода и приложения нажмите Esc")

OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильные аргументы в формате url описание теги")
				continue OuterLoop
			}
			newItem := Item{
				Link: args[0],
				Name: args[1],
				Tags: strings.Join(args[2:], ", "),
				Date: time.Now(),
			}
			if validateUrl(newItem.Link) {
				item = append(item, newItem)
				fmt.Println("Ссылка успешно добавлена!")
			}

		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>
			fmt.Printf("Всего записей: %d\n", len(item))
			if len(item) > 1 {
				fmt.Println("Список ссылок:")
				for i, link := range item {
					fmt.Printf("%d. "+
						"Имя: %s\n   "+
						"URL: %s\n   "+
						"Теги: %s\n   "+
						"Дата: %s\n",
						i+1, link.Name, link.Link, link.Tags, link.Date.Format("2024-02-15 15:04:05"))
				}
			}
		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("Введите имя ссылки, которое нужно удалить")
			//
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			_ = text
			text = strings.TrimSpace(text)

			for i, link := range item {
				if link.Name == text {
					item = append(item[:i], item[i+1:]...)
					fmt.Println("Ссылка успешно удалена!")
					break
				}
			}
		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}

func validateUrl(urlAddr string) bool {
	// Парсим URL
	u, err := url.ParseRequestURI(urlAddr)
	if err != nil {
		fmt.Println("Ошибка при парсинге URL:", err)
		return false
	}

	// Проверяем, является ли URL валидным
	if u.Scheme == "" || u.Host == "" {
		return false
	} else {
		return true
	}
}
