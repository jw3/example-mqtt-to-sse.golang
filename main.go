package main

import (
	"fmt"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"net/http"
	"os"
	"time"
)

func handleHealth(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

func main() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := mqtt.NewClientOptions().AddBroker("tcp://mqtt.teserakt.io:1883").SetClientID("go")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// channel for shuttling mqtt events
	events := make(chan [2]string)

	if token := c.Subscribe("#", 0, func(client mqtt.Client, msg mqtt.Message) {
		events <- [2]string{msg.Topic(), string(msg.Payload())}
		println(msg.Topic())
	}); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	s := sse.NewServer(nil)
	defer s.Shutdown()

	http.Handle("/events", s)
	http.HandleFunc("/health", handleHealth)

	go func() {
		for {
			e := <-events
			s.SendMessage("/events", sse.NewMessage("", e[1], e[0]))
		}
	}()

	http.ListenAndServe(":9000", nil)
}
