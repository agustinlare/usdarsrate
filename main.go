package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type Image struct {
	URL string `json:"url"`
}

type Embed struct {
	Image Image `json:"image"`
}

type Message struct {
	Username string  `json:"username"`
	Embeds   []Embed `json:"embeds"`
}

func main() {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

	takeScreenshot()

	imageFile, err := ioutil.ReadFile("screenshot.png")
	if err != nil {
		log.Fatal(err)
	}

	message := Message{
		Username: "To the moon",
		Embeds: []Embed{
			{
				Image: Image{
					URL: "attachment://screenshot.png",
				},
			},
		},
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormField("payload_json")
	if err != nil {
		log.Fatal(err)
	}
	part.Write(jsonData)

	part, err = writer.CreateFormFile("file", "screenshot.png")
	if err != nil {
		log.Fatal(err)
	}
	part.Write(imageFile)

	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}
	contentType := writer.FormDataContentType()

	req, err := http.NewRequest("POST", webhookURL, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(responseBody))
}

func takeScreenshot() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://dolarhoy.com/i/cotizaciones/dolar-blue"),
		chromedp.Sleep(2*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, err = page.CaptureScreenshot().WithQuality(90).WithClip(&page.Viewport{
				X:      0,
				Y:      0,
				Width:  330,
				Height: 260,
				Scale:  1,
			}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("screenshot.png", buf, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
