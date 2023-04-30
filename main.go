package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	// Replace with your webhook URL
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

	// Create a new Embed object
	embed := struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	}{
		URL:   "https://dolarhoy.com/i/cotizaciones/dolar-blue",
		Title: "Dolar Hoy Blue",
	}

	// Marshal the Embed object to JSON
	payload, err := json.Marshal(embed)
	if err != nil {
		panic(err)
	}

	// Create a new POST request to the webhook URL with the JSON payload
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
