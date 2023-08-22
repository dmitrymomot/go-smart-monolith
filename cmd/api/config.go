package main

import "github.com/dmitrymomot/go-env"

// Application configuration.
var (
	httpPort = env.GetInt("HTTP_PORT", 8080)
)
