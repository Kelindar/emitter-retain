package main

import (
	"fmt"
	"math/rand"
	"time"

	emitter "github.com/emitter-io/go/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sensor()
	dashboard()

	time.Sleep(5 * time.Second)
}

func dashboard() {
	c, err := emitter.Connect("tcp://127.0.0.1:8080", func(_ *emitter.Client, msg emitter.Message) {})
	if err != nil {
		panic(err)
	}

	if err := c.Subscribe("hAL-QD6ON8Yb-PYvn2LsQfToSZkz-J2T", "retain-demo/", func(_ *emitter.Client, msg emitter.Message) {
		println("received from sensor:", string(msg.Payload()))
	}); err != nil {
		println(err.Error())
	}
}

func sensor() {
	c, err := emitter.Connect("tcp://127.0.0.1:8080", func(_ *emitter.Client, msg emitter.Message) {})
	if err != nil {
		panic(err)
	}

	message := fmt.Sprintf("%d", rand.Int31n(100))
	if err := c.PublishWithRetain("DwdB6Ih6tiQo-t6BdOjmznW3L89yWVjP", "retain-demo/", message); err != nil {
		println(err.Error())
	}
}
