package client

import (
	"bufio"
	"fmt"
	"mediasoft/lesson7/pkg/api"
)

func scanMessage(inputScanner *bufio.Scanner) (*api.MessageIn, bool) {

	fmt.Print("Receiver nick: ")
	if !inputScanner.Scan() {
		return nil, false
	}
	toNick := inputScanner.Text()
	if toNick == "exit" {
		return nil, false
	}

	fmt.Print("Message body: ")
	if !inputScanner.Scan() {
		return nil, false
	}
	body := inputScanner.Text()
	if body == "exit" {
		return nil, false
	}

	result := &api.MessageIn{
		To:   toNick,
		Body: body,
	}
	return result, true
}
