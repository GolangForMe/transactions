package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Если ожидаем ~20 транзакций — выделяем память заранее
	transactions := make([]int, 0, 20)
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)

	for {
		choice := ""

		fmt.Print("Введите Вашу транзакцию: ")

		// 1. Сначала сканируем ввод
		if !scanner.Scan() {
			break
		}

		// 2. Потом получаем текст, обрезаем пробелы
		input := strings.TrimSpace(scanner.Text())

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

		fmt.Print("Хотите продолжить (yes/no)? ")

		// 1. Сначала сканируем ввод ответа
		if !scanner.Scan() {
			break
		}

		// 2. Только потом получаем текст, обрезаем пробелы и приводим к нижнему регистру
		choice = strings.ToLower(strings.TrimSpace(scanner.Text()))

		// Пользователь решил прекратить ввод данных
		if choice == "no" || choice == "n" {
			break
		}
	}

	// Проверяем, что команда scanner.Scan() завершилась по ошибке ввода (например, переполнение буфера)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения ввода: %v\n", err)
		// Завершаем программу с кодом ошибки
		os.Exit(1)
	}

	// Считаем общую сумму
	for _, transaction := range transactions {
		sum += transaction
	}

	fmt.Printf("\nКоличество Ваших транзакций: %d\n", len(transactions))
	fmt.Printf("\nОбщая сумма на Вашем счете: %d\n", sum)
}
