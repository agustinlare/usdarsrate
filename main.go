package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotizacion struct {
	Fecha  string
	Compra string
	Venta  string
}

type Newmessage struct {
	Text string `json:"text"`
}

func main() {
	webhookUrl := os.Getenv("WEBHOOK_URL")
	res, err := http.Get("https://api-dolar-argentina.herokuapp.com" + os.Getenv("ENDPOINT_URL"))

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data Cotizacion
	_ = decoder.Decode(&data)

	msg := fmt.Sprintf("%s, Compra: %s, Venta: %s", data.Fecha, data.Compra, data.Venta)

	chatBody, _ := json.Marshal(Newmessage{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(chatBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if buf.String() != "ok" {
		log.Fatal("exploto")
	}
}
