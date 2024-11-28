package goutil

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	urlutil "net/url"
	"os"
)

func Network_SetHeader(reqClient *http.Request, headers []map[string]string) *http.Request {
	if len(headers) > 0 {
		for k, v := range headers[0] {
			reqClient.Header.Set(k, v)
		}
	}
	return reqClient
}

func Network_SetFormData(data map[string]string) urlutil.Values {
	kvdata := urlutil.Values{}
	for k, v := range data {
		kvdata.Set(k, v)
	}
	return kvdata
}

func Network_Get(url string, headers ...map[string]string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req = Network_SetHeader(req, headers)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func Network_Post(url string, data any, dataType string, headers ...map[string]string) *http.Response {
	client := &http.Client{}
	if dataType == "text" {
		data := []byte(data.(string))
		req, err := http.NewRequest("POST", url, bytes.NewReader(data))
		if err != nil {
			log.Fatal(err)
		}
		req = Network_SetHeader(req, headers)
		req.Header.Set("Content-Type", "text/plain")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		return resp
	}
	if dataType == "binary" {
		req, err := http.NewRequest("POST", url, bytes.NewReader(data.([]byte)))
		if err != nil {
			log.Fatal(err)
		}
		req = Network_SetHeader(req, headers)
		req.Header.Set("Content-Type", "application/octet-stream")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		return resp
	}
	if dataType == "json" {
		data, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			log.Fatal(err)
		}
		req = Network_SetHeader(req, headers)
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		return resp
	}
	if dataType == "form-kv" {
		data := Network_SetFormData(data.(map[string]string))
		req, err := http.NewRequest("POST", url, bytes.NewBufferString(data.Encode()))
		if err != nil {
			log.Fatal(err)
		}
		req = Network_SetHeader(req, headers)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		return resp
	}
	if dataType == "form-file" {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		data := data.(map[string]string)
		for partname, filepath := range data {
			file, err := os.Open(filepath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			part, err := writer.CreateFormFile(partname, filepath)
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(part, file)
			if err != nil {
				log.Fatal(err)
			}
		}
		err := writer.Close()
		if err != nil {
			log.Fatal(err)
		}
		req, err := http.NewRequest("POST", url, body)
		if err != nil {
			log.Fatal(err)
		}
		req = Network_SetHeader(req, headers)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		return resp
	}
	return nil
}
