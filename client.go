package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Client struct {
	send chan Message
}

func (client *Client) write() {
	for msg := range client.send {
		//TODO: socket.sendJSON(msg)
		fmt.Printf("%#v\n", msg)
	}

}

func (client *Client) subscribeChannels() {
	//TODO: changefeed Query RethinkDB
	for {
		time.Sleep(r())
		client.send <- Message{"channel add", ""}
	}
}

func (client *Client) subscribeMessages() {
	//TODO: changefeed Query RethinkDB
	for {
		time.Sleep(r())
		client.send <- Message{"message add", ""}
	}
}

func NewClient() *Client {
	return &Client{
		send: make(chan Message),
	}
}

func r() time.Duration {
	return time.Millisecond * time.Duration(rand.Intn(1000))
}

func main() {
	client := NewClient()
	go client.subscribeChannels()
	go client.subscribeMessages()
	client.write()
}
