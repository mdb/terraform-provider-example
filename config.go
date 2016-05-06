package example

import (
	"crypto/tls"
	"log"
	"net/http"
)

type Config struct {
	AccessKey string
	SecretKey string
	Host      string
}

type ExampleClient struct {
	AccessKey  string
	SecretKey  string
	Host       string
	HttpClient *http.Client
}

func (c *Config) Client() (ExampleClient, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := ExampleClient{
		AccessKey:  c.AccessKey,
		SecretKey:  c.SecretKey,
		Host:       c.Host,
		HttpClient: &http.Client{Transport: tr},
	}

	log.Printf("[INFO] Example Client configured for use")

	return client, nil
}
