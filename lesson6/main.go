package main

import (
	"fmt"
	"time"
)

var commands = `
1 - Nil channel - close(c)
2 - Nil channel - c <- val
3 - Nil channel - val := <-c
4 - Closed channel - close(c)
5 - Closed channel - c <- val
6 - Closed channel - val := <-c
7 - Not-nil not-closed channel - close(c)
8 - Not-nil not-closed channel - c <- val
9 - Not-nil not-closed channel - val := <-c
10 - Выйти из программы
`

func main() {
	isEnd := false

	for !isEnd {
		cmd := 0
		fmt.Print(commands)
		fmt.Scanf("%d\n", &cmd)

		switch cmd {
		case 1:
			var c chan int
			close(c) // panic: close of nil channel
		case 2:
			var c chan int
			go func() {
				c <- 42 // Блокируется навсегда
			}()
			fmt.Println("Эта строка выведется, но горутина зависнет")
			select {}
		case 3:
			var c chan int
			go func() {
				val := <-c // Блокируется навсегда
				fmt.Println(val)
			}()
			fmt.Println("Горутина заблокирована")
			select {} // чтобы программа не завершилась
		case 4:
			c := make(chan int)
			close(c)
			close(c) // panic: close of closed channel
		case 5:
			c := make(chan int)
			close(c)
			c <- 42 // panic: send on closed channel
		case 6:
			c := make(chan int)
			close(c)
			val, ok := <-c // не блокируется, возвращает 0, false
			fmt.Printf("Value: %d, OK: %v\n", val, ok)
		case 7:
			c := make(chan int)
			close(c) // успешно
			fmt.Println("Channel closed successfully")
		case 8:
			c := make(chan int)
			go func() {
				time.Sleep(1 * time.Second)
				<-c // чтобы избежать блокировки на отправке
			}()
			c <- 42 // блокируется до тех пор, пока не будет чтение
			fmt.Println("Value sent after unblocking")
		case 9:
			c := make(chan int)
			go func() {
				c <- 42 // отправка значения
			}()
			time.Sleep(100 * time.Millisecond) // даем время на запуск горутины
			val := <-c                         // читаем
			fmt.Printf("Received: %d\n", val)
		case 10:
			isEnd = true
			fmt.Println("Выход из программы")
		default:
			fmt.Println("Неизвестная команда")
		}
		if isEnd {
			break
		}
	}
}
