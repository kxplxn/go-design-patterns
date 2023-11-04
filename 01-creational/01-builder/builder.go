package main

import "fmt"

// NotificationBuilder has fields exported
type NotificationBuilder struct {
	Title    string
	SubTitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) {
	nb.SubTitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(pri int) {
	nb.Priority = pri
}

func (nb *NotificationBuilder) SetType(notType string) {
	nb.NotType = notType
}

// The Build method returns a fully finished Notification object
func (nb *NotificationBuilder) Build() (*Notification, error) {
	// error checking can be done at the Build stage
	if nb.Icon != "" && nb.SubTitle == "" {
		return nil, fmt.Errorf("subtitle must be specified when using an icon")
	}

	if nb.Priority < 0 || nb.Priority > 5 {
		return nil, fmt.Errorf("priority must be between 0 and 5")
	}

	// return a newly created Notification object using the current settings
	return &Notification{
		title:    nb.Title,
		subtitle: nb.SubTitle,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
