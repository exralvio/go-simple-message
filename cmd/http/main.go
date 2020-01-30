package main

import "github.com/exralvio/go-simple-message/cmd/http/container"

import "github.com/exralvio/go-simple-message/cmd/http/handlers"

func main() {
	App := container.NewAppContainer()

	handlers.HTTPHandler(App)

	App.Run()
}
