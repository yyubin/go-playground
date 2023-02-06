package main

import (
	"fmt"
	"goweb2/adapter/adpater"
	"goweb2/adapter/usba"
	"goweb2/adapter/usbc"
)

type Client struct{}

func (c *Client) InsertUsb(usb Usb) {
	fmt.Println("Clients insert usb to port")
	usb.InsertUsbToPort()
}

type Usb interface {
	InsertUsbToPort()
}

func main() {
	client := &Client{}
	usbc := &usbc.Usbc{}

	client.InsertUsb(usbc)

	usba := &usba.Usba{}
	adapter := &adpater{
		usba: usba,
	}

	client.InsertUsb(adapter)
}
