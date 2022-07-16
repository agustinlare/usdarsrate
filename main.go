package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Cotizacion struct {
	Fecha  string
	Compra string
	Venta  string
}

type Message struct {
	Text string `json:"text"`
}

func main() {
	webhookUrl := os.Getenv("WEBHOOK_URL")

	for _, e := range strings.Split(os.Getenv("ENDPOINT_URL"), ",") {
		c, _ := getCotizacion(e)
		req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(c))
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

		if resp.StatusCode != 200 {
			log.Fatal(resp.Status)
		}

	}

	os.Exit(0)
}

func getCotizacion(e string) ([]byte, error) {
	res, err := http.Get("https://api-dolar-argentina.herokuapp.com" + e)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data Cotizacion
	_ = decoder.Decode(&data)
	name := strings.Split(e, "/")[2]
	msg := fmt.Sprintf("%s %s, Compra: %s, Venta: %s", name, data.Fecha, data.Compra, data.Venta)

	return json.Marshal(Message{Text: msg})
}
