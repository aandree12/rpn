package main

import (
	"fmt"
	"strings"
	"unicode"
)

func transformString(s string) string {
	words := strings.Fields(s) // Разделяем строку на слова
	for i, word := range words {
		var transformedWord strings.Builder // Используем StringBuilder для эффективной конкатенации строк
		for j, char := range word {
			if j%2 == 0 {
				transformedWord.WriteRune(unicode.ToUpper(char)) // Четный индекс - переводим в верхний регистр
			} else {
				transformedWord.WriteRune(unicode.ToLower(char)) // Нечетный индекс - переводим в нижний регистр
			}
		}
		words[i] = transformedWord.String() // Обновляем каждое слово в срезе
	}
	return strings.Join(words, " ") // Соединяем слова обратно в строку
}

func main() {
	fmt.Println(transformString("String"))            // Вывод: StRiNg
	fmt.Println(transformString("Weird string case")) // Вывод: WeIrD StRiNg CaSe
}
