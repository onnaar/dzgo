package main

import (
	"errors"
	"fmt"
	"lesson5/student/practice"
)

func output(notifiers []practice.Notifier) {
	for i, n := range notifiers {
		fmt.Printf("%d %v\n", i, n)
	}
	fmt.Println()
}

func main() {
	// TODO: домашняя работа. Покажите успешную рассылку, пустой текст,
	// ошибку настройки и продолжение после BrokenNotifier.
	// Импорт учебного пакета: "lesson5/student/practice".
	sliceRight := []practice.Notifier{
		&practice.EmailNotifier{
			BaseNotifier: practice.BaseNotifier{
				Channel: "c1",
			},
			Address:  "pushkin's street",
			Messages: []string{"kolotushkin's home", "AWdawd"},
		},
		&practice.PushNotifier{
			BaseNotifier: practice.BaseNotifier{
				Channel: "c2",
			},
			DeviceID: "phone9913",
			Messages: []string{"a", "hey"},
		},
	}
	sliceEmptyMsg := []practice.Notifier{
		&practice.EmailNotifier{
			BaseNotifier: practice.BaseNotifier{
				Channel: "c3",
			},
			Address:  "pupunya@gmail.com",
			Messages: []string{"kwdolotushkn's home", "AWdawd"},
		},
		&practice.PushNotifier{
			BaseNotifier: practice.BaseNotifier{
				Channel: "c4",
			},
			DeviceID: "laptop666",
			Messages: []string{"a", "hey"},
		},
	}
	sliceWrongFields := []practice.Notifier{
		&practice.EmailNotifier{
			BaseNotifier: practice.BaseNotifier{
				Channel: "c3",
			},
			Address:  "",
			Messages: []string{"kWdawd"},
		},
		&practice.PushNotifier{
			BaseNotifier: practice.BaseNotifier{
				Channel: "c4",
			},
			DeviceID: "",
			Messages: []string{"a", "hey"},
		},
		&practice.EmailNotifier{
			BaseNotifier: practice.BaseNotifier{
				Channel: "",
			},
			Address:  "pauawdawdshkin's street",
			Messages: []string{"kawd"},
		},
		&practice.PushNotifier{
			BaseNotifier: practice.BaseNotifier{
				Channel: "",
			},
			DeviceID: "phone1213",
			Messages: []string{"a", "hey"},
		},
	}
	sliceBroken := []practice.Notifier{
		&practice.BrokenNotifier{
			BaseNotifier: practice.BaseNotifier{
				Channel: "c5",
			},
		},
		&practice.PushNotifier{
			BaseNotifier: practice.BaseNotifier{
				Channel: "c6",
			},
			DeviceID: "12324",
			Messages: []string{"ah"},
		},
	}

	//1 w/o errors
	msg := "Заказ готов"
	if err := practice.Dispatch(sliceRight, msg); err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("all was good")
	}
	output(sliceRight)

	empty := ""
	if err := practice.Dispatch(sliceEmptyMsg, empty); errors.Is(err, practice.ErrEmptyMessage) {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("good bruh")
	}
	output(sliceEmptyMsg)

	if errs := practice.SafeDispatch(sliceWrongFields, msg); errs != nil {
		for i, err := range errs {
			var sendErr *practice.SendError
			if errors.As(err, &sendErr) {
				fmt.Printf("%d %v\n", i, sendErr)
			} else {
				fmt.Printf("another error\n")
			}
		}
	}
	output(sliceWrongFields)

	if errs := practice.SafeDispatch(sliceBroken, msg); errs != nil {
		for i, err := range errs {
			if errors.Is(err, practice.ErrNotifierPanic) {
				fmt.Printf("%d %v\n", i, err)
			} else {
				fmt.Printf("another error\n")
			}
		}
	}
	output(sliceBroken)
}
