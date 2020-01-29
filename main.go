package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	Host = "https://localhost:8080"
	Path = "/api/v1/test"
	TPS  = 5
)

func main() {
	ticker := time.NewTicker(time.Duration((1000 / TPS)) * time.Millisecond)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				go TestThisApi()
				fmt.Println()
			}
		}
	}()

	time.Sleep(240 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

type ApiUnderTestPayload struct {
	SelfieURL string `json:"selfie_url"`
}

func TestThisApi() {

	data := ApiUnderTestPayload{
		// fill struct
		SelfieURL: "selfie_url",
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)
	start := time.Now()

	req, err := http.NewRequest("POST", Host+Path, body)
	if err != nil {
		// handle err
	}
	req.Host = Host
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "")
	req.Header.Set("Accept-Language", "id")
	req.Header.Set("Auth-Token", "")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "218")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	elapsed := time.Since(start).Seconds()
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Response:", string(b), " Status Code", resp.StatusCode, " Response Time", elapsed, " seconds")
}
