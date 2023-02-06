package usbc

import "fmt"

type Usbc struct{}

func (u *Usbc) InsertUsbToPort() {
	fmt.Println("Usb converter is plugged into Usbc.")
}
