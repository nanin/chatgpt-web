package main

import (
	bootstrap "chatgpt-web/bootstarp"
	"chatgpt-web/config"

	"github.com/alecthomas/kong"
)

func main() {
	kong.Parse(&config.CLI)
	bootstrap.StartWebServer()
}
