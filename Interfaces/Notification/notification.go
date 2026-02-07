package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	message string
}

func (e *EmailNotifier) Send(message string) error {
	fmt.Println("Email sent with message: ", message)
	return nil
}

type SMSNotifier struct {
	message string
}

func (s *SMSNotifier) Send(message string) error {
	fmt.Println("SMS sent with message: ", message)
	return nil
}

func main() {
	notifiers := []Notifier{
		&EmailNotifier{},
		&SMSNotifier{},
	}

	for _, v := range notifiers {
		v.Send("Hello!")
	}

}
