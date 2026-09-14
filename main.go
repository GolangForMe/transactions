package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	transactions := make([]int, 0, 20)
	sum := 0

	for {
		var input string
		var choise string

		fmt.Print("Введите Вашу транзакцию: ")
		fmt.Scan(&input)

		// Обрезка пустых симолов слева и справа от ввода пользователя
		input = strings.TrimSpace(input)

		// Пытаемся преобразовать строку к целому числу
		transaction, err := strconv.Atoi(input)

		// Проверяем, удалось ли преобразовать строку в целое число
		if err != nil {
			fmt.Println("Ошибка: пожалуйста, введите корректное целое число.")
			continue
		}

		// Запрещаем пользователю вводить нулевые транзакции
		if transaction == 0 {
			fmt.Println("Транзакция не может быть равна 0. Введите другое число.")
			continue
		}

		// Все прошло успешно - добавляем в слайс новую транзакцию
		transactions = append(transactions, transaction)

		fmt.Println("Хотите продолжить (yes/no)? ")
		fmt.Scan(&choise)

		// Обрезка пустых симолов слева и справа от ввода пользователя
		choise = strings.ToLower(strings.TrimSpace(choise))

		// Пользователь решил прекратить ввод данных
		if choise == "no" || choise == "n" {
			break
		}
	}

	// Считаем общую сумму
	for _, transaction := range transactions {
		sum += transaction
	}

	fmt.Printf("Колличество Ваших транзакций: %d\n", len(transactions))
	fmt.Printf("Общая сумма на Вашем счете: %d\n", sum)
}
