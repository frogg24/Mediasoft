package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"mediasoft/lesson7/pkg/api"
	"net"
	"os"
)

func Start() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		welcomeMsg := scanner.Text()
		var serverMsg api.MessageOut
		if err := json.Unmarshal([]byte(welcomeMsg), &serverMsg); err == nil {
			fmt.Printf("Server: %s", serverMsg.Body)
		} else {
			fmt.Printf("%s", welcomeMsg)
		}
	}

	inputScanner := bufio.NewScanner(os.Stdin)
	if inputScanner.Scan() {
		username := inputScanner.Text()

		conn.Write([]byte(username + "\n"))

		if scanner.Scan() {
			response := scanner.Text()
			var serverResponse api.MessageOut
			if err := json.Unmarshal([]byte(response), &serverResponse); err == nil {
				if serverResponse.From == "Server" &&
					(len(serverResponse.Body) >= 7 && serverResponse.Body[:7] == "Здорово!") {

					fmt.Printf("Server: %s", serverResponse.Body)
				} else if serverResponse.From == "Server" &&
					len(serverResponse.Body) >= 13 &&
					serverResponse.Body[:13] == "Ошибка: ник '" {

					fmt.Printf("Server: %s\n", serverResponse.Body)
					return
				} else {

					fmt.Printf("Server: %s\n", serverResponse.Body)
				}
			} else {

				fmt.Printf("Server: %s\n", response)
			}
		}
	} else {
		return
	}

	go func() {
		for scanner.Scan() {

			text := scanner.Text()
			var message = new(api.MessageOut)

			err := json.Unmarshal([]byte(text), message)
			if err != nil {
				fmt.Printf("System message: %s\n", text)
				continue
			}
			fmt.Printf("Message from %s: %s\n", message.From, message.Body)
		}
	}()

	fmt.Println("Type message to send, 'exit' to quit:")
	for {
		msg, msgOk := scanMessage(inputScanner)
		if !msgOk {
			break
		}

		msgJson, err := json.Marshal(msg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		conn.Write([]byte(string(msgJson) + "\n"))
	}

	conn.Close()
}
