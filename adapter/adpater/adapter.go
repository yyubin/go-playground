package adpater

import (
	"fmt"
	"goweb2/adapter/usba"
)

type UsbaToUsbcConverter struct {
	usba *usba.Usba
}

func (u *UsbaToUsbcConverter) InsertUsbToPort() {
	fmt.Println("Usba is plugged into the converter.")
	usba.InsertUsbToPort()
}
