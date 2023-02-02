//Домашнее задание
//Распаковка строки Создать Go функцию, осуществляющую примитивную
//распаковку строки, содержащую повторяющиеся символы / руны, например:
//* "a4bc2d5e" => "aaaabccddddde"
//* "abcd" => "abcd"
//* "45" => "" (некорректная строка)
//Дополнительное задание: поддержка escape - последовательности
//* qwe\4\5 => qwe45 (*)
//* qwe\45 => qwe44444 (*)
//* qwe\\5 => qwe\\\\\ (*)

package hw_2

import (
	"strings"
	"unicode"
)

func StringParse(str string) string {
	toRune := []rune(str)
	var sb strings.Builder
	l := 1
	var currentChar string
	for i := 0; i < len(toRune); i++ {
		switch {
		case toRune[i] == 92 && toRune[i+1] == 92:
			sb.WriteString(strings.Repeat(string(toRune[i+1]), int(toRune[i+2]-'0')))
			i += 2
		case toRune[i] == 92:
			currentChar = string(toRune[i+1])
			i++
			l = 1
			sb.WriteString(strings.Repeat(string(currentChar), l))
		case unicode.IsLetter(toRune[i]):
			currentChar = string(toRune[i])
			l = 1
			sb.WriteString(strings.Repeat(string(currentChar), l))
		case unicode.IsDigit(toRune[i]):
			l = int(toRune[i]-'0') - 1
			sb.WriteString(strings.Repeat(string(currentChar), l))
		}

	}
	return sb.String()
}
