package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"mediasoft/lesson7/pkg/api"
	"net"
	"sync"
)

var (
	mutex   sync.Mutex
	clients = make(map[string]net.Conn)
)

func Handler(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	//Получить ник
	welcomeMsg := api.MessageOut{
		From: "Server",
		Body: "Введи свой ник: ",
	}
	welcomeJSON, err := json.Marshal(welcomeMsg)
	if err != nil {
		return
	}

	fmt.Fprintln(conn, string(welcomeJSON))
	if !scanner.Scan() {
		return
	}
	username := scanner.Text()

	//Проверить уникальность и зарегистрировать
	if !TryAddUser(username, conn) {
		errorMsg := api.MessageOut{
			From: "Server",
			Body: fmt.Sprintf("Ошибка: ник '%s' занят", username),
		}
		errorJSON, _ := json.Marshal(errorMsg)
		fmt.Fprintln(conn, string(errorJSON))
		return
	}
	greetingMsg := api.MessageOut{
		From: "Server",
		Body: fmt.Sprintf("Здарова, %s!", username),
	}
	greetingJSON, _ := json.Marshal(greetingMsg)
	fmt.Fprintln(conn, string(greetingJSON))

	// Обработка сообщений
	for scanner.Scan() {
		line := scanner.Text()
		var msgIn api.MessageIn

		if err := json.Unmarshal([]byte(line), &msgIn); err != nil {
			continue
		}

		mutex.Lock()
		recipientConn, recipientExists := clients[msgIn.To]
		mutex.Unlock()

		if recipientExists && recipientConn != conn {
			msgOut := api.MessageOut{
				From: username,
				Body: msgIn.Body,
			}

			jsonMsgOut, err := json.Marshal(msgOut)
			if err != nil {
				continue
			}

			fmt.Fprintln(recipientConn, string(jsonMsgOut))
		} else {
			errorMsg := api.MessageOut{
				From: "Server",
				Body: fmt.Sprintf("Ошибка: пользователь '%s' не найден и тд.", msgIn.To),
			}
			errorJSON, _ := json.Marshal(errorMsg)
			fmt.Fprintln(conn, string(errorJSON))
		}
	}
	removeUser(username)
}

func TryAddUser(username string, conn net.Conn) bool {
	mutex.Lock()
	defer mutex.Unlock()

	if _, userExists := clients[username]; userExists {
		return false
	}

	clients[username] = conn
	return true
}

func removeUser(username string) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(clients, username)
}
