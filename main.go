package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"gopkg.in/toast.v1"
)

// Function to show a message on different platforms
func showMessage(notification toast.Notification) {
	switch runtime.GOOS {
	case "windows":
		err := notification.Push()
		if err != nil {
			println("Error showing notification:", err.Error())
		}
	case "darwin":
		cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s"`, notification.Message, notification.Title))
		cmd.Run()
	case "linux":
		cmd := exec.Command("notify-send", notification.Title, notification.Message)
		cmd.Run()
		fmt.Println("Platform not supported")
	default:
		fmt.Println("Platform not supported")
	}
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// duration between reminders
	duration := 60 * time.Minute

	// definition of notifications
	startnotification := toast.Notification{
		AppID:    "Movement Reminder",
		Title:    "Start Movement Reminder",
		Message:  "You will be reminded every " + duration.String() + " to move",
		Icon:     cwd + "/movementremindericon.png", // must exist!
		Duration: "long",
		Audio:    toast.SMS,
	}
	movenotification := toast.Notification{
		AppID:    "Movement Reminder",
		Title:    "Move",
		Message:  "You will feel better when you move a little bit and use your standing desk",
		Icon:     cwd + "/movementremindericon.png", // must exist!
		Duration: "long",
		Audio:    toast.SMS,
	}

	// Start ticker
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	// run
	showMessage(startnotification)
	for {
		select {
		case <-ticker.C:
			showMessage(movenotification)
		}
	}
}
