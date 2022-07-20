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

type Message struct {
	Text string `json:"text"`
}

func main() {
	res, err := http.Get(os.Getenv("ENDPOINT_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var data map[string]interface{}
	_ = decoder.Decode(&data)

	for _, v := range data {
		dolar, _ := v.(map[string]interface{})
		msg, _ := json.Marshal(Message{Text: fmt.Sprintf("%s %s, Compra: %s, Venta: %s", dolar["name"], dolar["date"], dolar["buy"], dolar["sell"])})
		sendMessage(os.Getenv("WEBHOOK_URL"), msg)
	}

	os.Exit(0)
}

func sendMessage(w string, b []byte) {
	req, err := http.NewRequest(http.MethodPost, w, bytes.NewBuffer(b))
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
