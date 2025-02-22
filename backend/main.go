package main

import (
	"backend/authentication"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Message struct {
	Type    string `json:"type"`
	Payload string `json: "payload"`
}

var clients = make(map[string]*websocket.Conn)
var mutex = sync.Mutex{}

var secretKey = "your_secret_key"

func main() {
	// API Authentication
	app := fiber.New()

	// Public route
	app.Post("/login", authentication.LoginHandler)

	// Protected route group
	api := app.Group("/api")
	api.Get("/protected", authentication.ProtectedHandler)

	// Database Connection

	// Data Processing

	// WebSocket Connection
	app.Get("/ws", websocket.New(HandleWebSocket))

	go monitorConnections()
	log.PrintLn("Server is running on port 3000")
	app.Listen(":3000")
}

// WebSocket Request Handling
