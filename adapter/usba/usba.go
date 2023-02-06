package usba

import "fmt"

type Usba struct{}

func (u *Usba) InsertUsbToPort() {
	fmt.Printf("Usb connecting terminal is plugged into usba")
}
