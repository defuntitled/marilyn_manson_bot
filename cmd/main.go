package main

import (
	"log"
	server "marilyn_manson_bot/internal/server"
)

func main() {
	server.Bot.Start()
	log.Default().Print("starting bot")
}
