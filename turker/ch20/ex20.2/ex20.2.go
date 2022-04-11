package main

import (
	"golang-test/test01/example01/turker/ch20/fedex"
	"golang-test/test01/example01/turker/ch20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var sender Sender = &fedex.FedexSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)

	var koreaPostSender Sender = &koreaPost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)
}
