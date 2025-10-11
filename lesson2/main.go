package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands = `
1 - Добавить значение в стек
2 - Получить значение из стека (с удалением)
3 - Просмотр значения на вершине стека
4 - Добавление в односвязный список
5 - Удаление из односвязного списка по индексу
6 - Получение значения из односвязного списка по индексу
7 - Получение слайса значений из односвязного списка
8 - Добавления в бинарное дерево
9 - Удаление из бинарного дерева по значению
10 - Получение отсортированного слайса значений из бинарного дерева
11 - Конвертирование из римских цифр в арабские
12 - Генерация двумерного массива со случайными уникальными числами
13 - Закончить работу
`

func main() {
	stack := newStack(100)
	linkedList := newSinglyLinkedList()
	binaryTree := newTree()
	isEnd := false
	scanner := bufio.NewScanner(os.Stdin)

	for !isEnd {
		fmt.Print(commands)
		fmt.Print("Выберите команду: ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		cmd := 0

		if _, err := fmt.Sscanf(input, "%d", &cmd); err != nil {
			fmt.Println("Ошибка: введите число от 1 до 13")
			continue
		}

		switch cmd {
		case 1:
			fmt.Print("Введите значение: ")
			if !scanner.Scan() {
				break
			}
			value := scanner.Text()
			push(stack, value)
			fmt.Printf("Добавлено в стек: %v\n", value)

		case 2:
			temp := pop(stack)
			if temp != nil {
				fmt.Printf("Получено из стека: %v\n", temp)
			} else {
				fmt.Println("Стек пуст")
			}

		case 3:
			temp := peek(stack)
			if temp != nil {
				fmt.Printf("Вершина стека: %v\n", temp)
			} else {
				fmt.Println("Стек пуст")
			}

		case 4:
			fmt.Print("Введите значение: ")
			if !scanner.Scan() {
				break
			}
			value := scanner.Text()
			addToList(linkedList, value)
			fmt.Printf("Добавлено в список: %v\n", value)

		case 5:
			fmt.Print("Индекс: ")
			if !scanner.Scan() {
				break
			}
			idx := 0
			if _, err := fmt.Sscanf(scanner.Text(), "%d", &idx); err != nil {
				fmt.Println("Ошибка: введите целое число")
				break
			}
			removeFromList(linkedList, idx)

		case 6:
			fmt.Print("Индекс: ")
			if !scanner.Scan() {
				break
			}
			idx := 0
			if _, err := fmt.Sscanf(scanner.Text(), "%d", &idx); err != nil {
				fmt.Println("Ошибка: введите целое число")
				break
			}
			fmt.Println(getFromList(linkedList, idx))

		case 7:
			values := valuesOfList(linkedList)
			fmt.Println("Значения в списке:")
			for i := 0; i < len(values); i++ {
				fmt.Println(values[i])
			}

		case 8:
			fmt.Print("Значение: ")
			if !scanner.Scan() {
				break
			}
			number := 0
			if _, err := fmt.Sscanf(scanner.Text(), "%d", &number); err != nil {
				fmt.Println("Ошибка: введите целое число")
				break
			}
			addToTree(binaryTree, number)

		case 9:
			fmt.Print("Значение: ")
			if !scanner.Scan() {
				break
			}
			number := 0
			if _, err := fmt.Sscanf(scanner.Text(), "%d", &number); err != nil {
				fmt.Println("Ошибка: введите целое число")
				break
			}
			removeFromTree(binaryTree, number)

		case 10:
			values := values(binaryTree)
			fmt.Println(values)

		case 11:
			fmt.Print("Римское число: ")
			if !scanner.Scan() {
				break
			}
			number := strings.TrimSpace(scanner.Text())
			result := convertRomanNumerals(number)
			if result != -1 {
				fmt.Println(result)
			} else {
				fmt.Println("Некорректный ввод")
			}

		case 12:
			fmt.Print("Количество строк: ")
			if !scanner.Scan() {
				break
			}
			m := 0
			if _, err := fmt.Sscanf(scanner.Text(), "%d", &m); err != nil {
				fmt.Println("Ошибка: введите целое число")
				break
			}

			fmt.Print("Количество столбцов: ")
			if !scanner.Scan() {
				break
			}
			n := 0
			if _, err := fmt.Sscanf(scanner.Text(), "%d", &n); err != nil {
				fmt.Println("Ошибка: введите целое число")
				break
			}

			if n*m > 1000 {
				fmt.Println("Число значений в матрице не должно превышать тысячу")
				break
			}
			matrix := generateMatrix(m, n)
			for _, row := range matrix {
				for _, value := range row {
					fmt.Print(value, "\t")
				}
				fmt.Println()
			}

		case 13:
			isEnd = true

		default:
			fmt.Println("Неизвестная команда. Введите число от 1 до 13.")
		}
	}
}
