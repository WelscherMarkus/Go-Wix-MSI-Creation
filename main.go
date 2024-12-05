package main

import (
	"github.com/go-toast/toast"
	"log"
)

func main() {
	notification := toast.Notification{
		AppID:   "Test",
		Title:   "Test",
		Message: "Test",
		Audio:   toast.Reminder,
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
