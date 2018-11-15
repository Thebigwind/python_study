package main

import (
	"log"

	"github.com/coreos/go-etcd/etcd"
)

func main() {
	client := etcd.NewClient(
		[]string{
			"http://127.0.0.1:2379",
		},
	)
	for {
		resp, err := client.Get("message", false, false)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Current creds: %s: %s\n", resp.Node.Key, resp.Node.Value)
		receiver := make(chan *etcd.Response)
		go client.Watch("/message", 0, false, receiver, nil)
		r := <-receiver
		log.Printf("Got updated creds: %s: %s\n", r.Node.Key, r.Node.Value)
	}
}
