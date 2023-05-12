package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatalf("Variable de entorno no declarada")
	}

	endpointURL := os.Getenv("ENDPOINT_DOLAR_URL")
	if webhookURL == "" {
		log.Fatalf("Variable de entorno no declarada")
	}

	list := strings.Split(endpointURL, ",")
	for _, item := range list {
		response, err := http.Get(item)
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		decoder := json.NewDecoder(response.Body)
		var content map[string]interface{}
		_ = decoder.Decode(&content)
		name := strings.Title(strings.Split(item, "/")[4])
		message := fmt.Sprintf("%s: Buy %s - Sell %s, Variacion: %s %s - Cierre Ant: %s | Fecha: %s",
			name,
			content["compra"],
			content["venta"],
			content["class-variacion"],
			content["variacion"],
			content["valor_cierre_ant"],
			content["fecha"],
		)

		sendNotification(webhookURL, message)
		// influxdb := os.Getenv("INFLUX_DB_URL")
		// if influxdb != "" {
		// 	err := writeToInfluxDB(content, name, influxdb)
		// 	if err != nil {
		// 		fmt.Println("Error writing to InfluxDB:", err)
		// 	}
		// }
	}

	os.Exit(0)
}

// func writeToInfluxDB(content map[string]interface{}, name string, influxdb string) error {
// 	c, err := client.NewHTTPClient(client.HTTPConfig{
// 		Addr: influxdb,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	defer c.Close()

// 	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
// 		Database:  "dolar",
// 		Precision: "s",
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	tags := map[string]string{"currency": "USD"}
// 	fields := map[string]interface{}{
// 		"name": name,
// 		"sell": content["venta"],
// 		"buy":  content["compra"],
// 	}
// 	p, err := client.NewPoint("exchange_rate", tags, fields, time.Now())
// 	if err != nil {
// 		return err
// 	}

// 	bp.AddPoint(p)

// 	if err := c.Write(bp); err != nil {
// 		return err
// 	}

// 	return nil
// }

func sendNotification(webhookUrl, message string) {
	var field string = "text"

	if strings.Contains(webhookUrl, "discord") {
		field = "content"
	}

	formData := url.Values{
		field: {message},
	}
	resp, err := http.PostForm(webhookUrl, formData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
