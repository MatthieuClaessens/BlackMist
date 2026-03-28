package services

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetTorIP() (string, error) {
	proxyURL, err := url.Parse("socks5://127.0.0.1:9050")
	if err != nil {
		return "", fmt.Errorf("Invalid proxy URL: %w", err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	resp, err := client.Get("https://api.ipify.org")
	if err != nil {
		return "Tor connection in progress...", err
	}
	defer resp.Body.Close()

	ipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Reading error: %w", err)
	}

	return string(ipBytes), nil
}
