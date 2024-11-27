package goutil

import (
	"bytes"
	"log"
	"net/http"
)

func Network_Get(url string, headers ...map[string]string) *http.Response {
	header := http.Header{}
	for _, h := range headers {
		for k, v := range h {
			header.Add(k, v)
		}
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp
}

func Network_Post(url string, data []byte, headers ...map[string]string) *http.Response {
	header := http.Header{}
	for _, h := range headers {
		for k, v := range h {
			header.Add(k, v)
		}
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp
}
