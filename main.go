package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zeroidentidad/gopherbot/config"
	"github.com/zeroidentidad/gopherbot/messages"
	"github.com/zeroidentidad/gopherbot/status"
)

func main() {
	dg, err := discordgo.New("Bot " + config.TOKEN)
	if err != nil {
		log.Println("Session error:", err)
		return
	}
	dg.Debug = true

	// Register handlers.
	dg.AddHandler(messages.MessageCreate)
	dg.AddHandler(status.SetStatus)

	err = dg.Open()
	if err != nil {
		log.Println("Websocket error,", err)
		return
	}

	go func() {
		http.HandleFunc("/", web)
		log.Fatal("Err http serv", http.ListenAndServe(":"+config.Port(), nil))
	}()

	// Until CTRL-C or other term signal.
	log.Println("Bot running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	// Close session.
	dg.Close()
}
