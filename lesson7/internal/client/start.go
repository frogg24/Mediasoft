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

	// First, get the welcome message from server asking for nickname
	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		welcomeMsg := scanner.Text()
		var serverMsg api.MessageOut
		if err := json.Unmarshal([]byte(welcomeMsg), &serverMsg); err == nil {
			fmt.Printf("Server: %s", serverMsg.Body) // Display the prompt without newline
		} else {
			fmt.Printf("%s", welcomeMsg) // Just print the raw message if it's not JSON
		}
	}

	// Get the username from the user
	inputScanner := bufio.NewScanner(os.Stdin)
	if inputScanner.Scan() {
		username := inputScanner.Text()
		
		// Send the username to the server
		conn.Write([]byte(username + "\n"))
		
		// Wait for server response to check if nickname is accepted
		if scanner.Scan() {
			response := scanner.Text()
			var serverResponse api.MessageOut
			if err := json.Unmarshal([]byte(response), &serverResponse); err == nil {
				if serverResponse.From == "Server" && 
				   (len(serverResponse.Body) >= 7 && serverResponse.Body[:7] == "Здарова") {
					// Nickname was accepted, continue with the program
					fmt.Printf("Server: %s", serverResponse.Body) // Show greeting
				} else if serverResponse.From == "Server" && 
					      len(serverResponse.Body) >= 13 && 
						  serverResponse.Body[:13] == "Ошибка: ник '" {
					// Nickname was rejected, display error and exit
					fmt.Printf("Server: %s\n", serverResponse.Body)
					return
				} else {
					// Other server message
					fmt.Printf("Server: %s\n", serverResponse.Body)
				}
			} else {
				// If we can't parse the response, treat as system message
				fmt.Printf("Server: %s\n", response)
			}
		}
	} else {
		return
	}

	// Now start the message receiving goroutine
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
