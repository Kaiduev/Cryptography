package main

import (
	"fmt"
	"os"
	s "strings"
)

// получение индекса буквы из массива
func index(alphabet []string, symbol string) int64 {
	var index int64
	for i := 0; i < len(alphabet); i++ {
		if symbol == alphabet[i] {
			index = int64(i)
		}
	}
	return index
}

// вычисление обратного по модулю числа
func rev(a int64, m int64) int64 {
	if a == 1 {
		return 1
	}
	return (1-rev(m%a, a)*m)/a + m
}

// Дешифратор
func deCode(alphabet []string, symbol string, k1 int64, k2 int64) string {
	idx := index(alphabet, symbol)
	d := rev(k1, int64(len(alphabet))) * (idx + int64(len(alphabet)) - k2) % int64(len(alphabet))
	return alphabet[d]
}

// Шифратор
func enCode(alphabet []string, symbol string, k1 int64, k2 int64) string {
	idx := index(alphabet, symbol)
	characterIndex := ((k1 * idx) + k2) % int64(len(alphabet))
	return alphabet[characterIndex]
}

// Дешифрует сообщение и возвращает расшифрованное сообщение
func deCodeAffineCipher(alphabet []string, text string, k1 int64, k2 int64) string {
	var deCodedText string
	textTOdeCode := s.Split(text, "")
	for i := 0; i < len(textTOdeCode); i++ {
		elemFromAlphabet := deCode(alphabet, textTOdeCode[i], k1, k2)
		deCodedText = deCodedText + elemFromAlphabet
	}
	return deCodedText
}

// Перебирает все ключи 12*26=312, 312 возможных ключей
func hackAffineCipher(alphabet []string, text string) {
	var bool int64
	for i := 1; i < len(alphabet); i += 2 {
		if i == 13 {
			continue
		}
		for j := 0; j < len(alphabet); j++ {
			fmt.Println(deCodeAffineCipher(alphabet, text, int64(i), int64(j)))
			fmt.Println("Нажмите 0 если нашли результат, нажмиие 1 для продолжения")
			fmt.Fscan(os.Stdin, &bool)
			if bool == 0 {
				break
			}
		}
		if bool == 0 {
			break
		}
	}
}

func main() {
	alphabet := []string{"A", "B", "C",
		"D", "E", "F",
		"G", "H", "I",
		"J", "K", "L",
		"M", "N", "O",
		"P", "Q", "R",
		"S", "T", "U",
		"V", "W", "X",
		"Y", "Z"}

	var a int64
	var b int64
	var menu int64
	fmt.Println("Меню выбора режима")
	fmt.Println("Нажмите 1 для того чтобы зашифровать сообщение\nНажмите 2 для того чтобы расшифровать сообщение\nНажмите 3 если вы не имеете ключей")
	fmt.Fscan(os.Stdin, &menu)
	if menu == 1 {
		fmt.Print("Введите первый ключ: ")
		fmt.Fscan(os.Stdin, &a)
		fmt.Print("Введите второй ключ: ")
		fmt.Fscan(os.Stdin, &b)
		if (a%2) > 0 && a < int64(len(alphabet)) && a > 0 {
			if b < int64(len(alphabet)) && b >= 0 {
				var text string
				fmt.Print("Введите сообщение: ")
				fmt.Fscan(os.Stdin, &text)
				var enCodedText string
				textTOenCode := s.Split(text, "")

				for i := 0; i < len(textTOenCode); i++ {
					elemFromAlphabet := enCode(alphabet, textTOenCode[i], a, b)
					enCodedText = enCodedText + elemFromAlphabet
				}
				fmt.Println(enCodedText)
			} else {
				fmt.Println("Один из ключей неверный")
			}
		} else {
			fmt.Println("Один из ключей неверный")
		}
	} else if menu == 2 {
		fmt.Print("Введите первый ключ: ")
		fmt.Fscan(os.Stdin, &a)
		fmt.Print("Введите второй ключ: ")
		fmt.Fscan(os.Stdin, &b)
		if (a%2) > 0 && a < int64(len(alphabet)) && a > 0 {
			if b < int64(len(alphabet)) && b >= 0 {
				var text string
				fmt.Print("Введите сообщение: ")
				fmt.Fscan(os.Stdin, &text)
				var deCodedText string
				textTOdeCode := s.Split(text, "")
				for i := 0; i < len(textTOdeCode); i++ {
					elemFromAlphabet := deCode(alphabet, textTOdeCode[i], a, b)
					deCodedText = deCodedText + elemFromAlphabet
				}
				fmt.Println(deCodedText)
			} else {
				fmt.Println("Один из ключей неверный")
			}
		} else {
			fmt.Println("Один из ключей неверный")
		}
	} else if menu == 3 {
		var text string
		fmt.Print("Введите сообщение: ")
		fmt.Fscan(os.Stdin, &text)
		hackAffineCipher(alphabet, text)
	}
}
