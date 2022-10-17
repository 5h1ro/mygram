package main

import (
	"mygram/handler/rest"
)

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	rest.StartApp()
}
