// Пакет клавиатура считывает ввод пользователя с клавиатуры.
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat считывает номер с плавающей точкой с клавиатуры.
// он возвращает читать номер и любую встречу с ошибкой.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}