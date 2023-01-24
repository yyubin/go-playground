package cj

import "fmt"

type PostSender struct{}

func (c *PostSender) Send(parcel string) {
	fmt.Printf("CJ에서 택배 %s를 보냅니다 \n", parcel)
}
