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
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

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
		message := fmt.Sprintf("%s %s, Compra: %s, Venta: %s", dolar["name"], dolar["date"], dolar["buy"], dolar["sell"])

		sendNotification(os.Getenv("WEBHOOK_URL"), message)

		influxdb, exists := os.LookupEnv("INFLUXDB_URL")

		if exists && strings.Contains(dolar["name"].(string), "blue") {

			err := writeToInfluxDB(dolar["sell"].(string), dolar["buy"].(string), influxdb)
			if err != nil {
				fmt.Println("Error writing to InfluxDB:", err)
			}
		}
	}

	os.Exit(0)
}

func writeToInfluxDB(sell string, buy string, influxdb string) error {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: influxdb,
	})
	if err != nil {
		return err
	}
	defer c.Close()

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "dolar",
		Precision: "s",
	})
	if err != nil {
		return err
	}

	tags := map[string]string{"currency": "USD"}
	fields := map[string]interface{}{
		"sell": sell,
		"buy":  buy,
	}
	p, err := client.NewPoint("exchange_rate", tags, fields, time.Now())
	if err != nil {
		return err
	}

	bp.AddPoint(p)

	if err := c.Write(bp); err != nil {
		return err
	}

	return nil
}

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
