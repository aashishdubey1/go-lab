package interfaces

import (
	"errors"
	"fmt"
)

// Question -> Notification System
// Design Notifier interface{ Send(msg string) error }
// with EmailNotifier, SMSNotifier, and PushNotifier implementations,
// plus a MultiNotifier that holds []Notifier and fans a single Send call out to all of them,
// collecting any errors with errors.Join.

type Notifier interface {
	Send(msg string) error
}

type EmailNotifier struct{}

func (en *EmailNotifier) Send(msg string) error {
	if msg == "" {
		return errors.New("no msg fucker")
	}
	fmt.Println("msg send Email ")
	return nil
}

type SMSNotifier struct{}

func (sn *SMSNotifier) Send(msg string) error {
	if msg == "" {
		return errors.New("no msg fucker")
	}
	fmt.Println("msg send by Sms ")
	return nil
}

type PushNotifier struct{}

func (pn *PushNotifier) Send(msg string) error {
	if msg == "" {
		return errors.New("no msg fucker")
	}
	fmt.Println("msg send Push  ")
	return nil
}

type MultiNotifier struct {
	Notifier []Notifier
}

func (mn *MultiNotifier) Send(msg string) error {
	var combinedErr error
	for _, v := range mn.Notifier {
		err := v.Send(msg)
		if err != nil {
			combinedErr = errors.Join(combinedErr, err)
		}
	}
	return combinedErr
}

func RunEx3() {

	email := EmailNotifier{}
	// sms := SMSNotifier{}
	// push := PushNotifier{}

	// Notifier := MultiNotifier{Notifier: []Notifier{&email, &sms, &push}}

	email.Send("how are you")
}
