package main

import "fmt"

func main() {
	// create a NotificationBuilder and use it to set properties
	builder := newNotificationBuilder()

	// use the builder to set some properties
	builder.SetTitle("New Notification")
	builder.SetIcon("icon.png")
	builder.SetSubTitle("This is a subtitle")
	builder.SetImage("image.jpg")
	builder.SetPriority(5)
	builder.SetMessage("This is a basic notification with some text in it")
	builder.SetType("alert")

	// use the Build function to create a finished object
	notification, err := builder.Build()
	if err != nil {
		fmt.Println("error building notification:", err)
		return
	}
	fmt.Printf("notification: %+v\n", notification)
}
