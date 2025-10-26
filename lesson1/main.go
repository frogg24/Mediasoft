package main

import "fmt"

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Сортировка пузырьком
5 - Выйти из программы
`

func main() {
	arr := [100]int{542, -565, 531, -294, -56, 14, 270, -51, -914, 605, -117, -768, 331, 708, -603, 84, -548, 579, 434, 751, 592, -349, 408, -602, 721, 909, 170, -432, -970, -171, -972, 316, 405, -676, -929, -795, -682, -646, 46, -609, -84, 180, -158, -662, -384, 854, -721, 39, 180, -197, -818, -946, -529, -555, -36, -853, -322, 540, -936, -919, 473, 978, 782, 586, 869, 333, -977, -548, -789, 988, -393, 807, -609, 997, 824, -480, -205, -576, 856, 494, 131, 40, -601, 467, 221, -640, 34, -220, 482, 948, 523, -27, -771, -914, 438, 957, 205, -411, -749, -723}

	const size = 512
	empls := [size]*Employee{}

	isEnd := false

	for !isEnd {
		cmd := 0
		fmt.Print(commands)

		//чтобы \n не оставался в буфере и не дублировал команды
		fmt.Scanf("%d\n", &cmd)

		switch cmd {
		case 1:
			empl := new(Employee)
			fmt.Println("\nИмя:")
			fmt.Scanf("%s\n", &empl.Name)
			fmt.Println("Возраст:")
			fmt.Scanf("%d\n", &empl.Age)
			fmt.Println("Позиция:")
			fmt.Scanf("%s\n", &empl.Position)
			fmt.Println("Зарплата:")
			fmt.Scanf("%d\n", &empl.Salary)
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					fmt.Println("Был добавлен сотрудник", empl)
					break
				}
			}
		case 2:
			fmt.Println("Удаляем сотрудника. Введите номер(из списка) сотрудника для удаления:")
			id := -1
			fmt.Scanf("%d\n", &id)
			if id < 0 {
				fmt.Println("Неверный формат")
			} else if empls[id-1] == nil {
				fmt.Println("Нет такого сотрудника")
			} else {
				empls[id-1] = nil
				fmt.Println("Сотрудник удален")
			}
		case 3:
			fmt.Println("Вывод сотрудников")
			for i := 0; i < size; i++ {
				if empls[i] != nil {
					fmt.Println(i+1, ".", empls[i])
				}
			}
		case 4:
			len := len(arr) - 1
			for i := 0; i < len; i++ {
				for j := 0; j < len-i; j++ {

					if arr[j] > arr[j+1] {
						temp := arr[j]
						arr[j] = arr[j+1]
						arr[j+1] = temp
					}
				}
			}

			for i := 0; i < 100; i++ {
				fmt.Println(arr[i])
			}
		case 5:
			isEnd = true
		}
		if isEnd {
			break
		}
	}

}
