package main

import (
	"goweb2/interface/cj"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	sender := &cj.PostSender{}
	SendBook("안나 카레리나", sender)
	SendBook("죄와 벌", sender)

}
