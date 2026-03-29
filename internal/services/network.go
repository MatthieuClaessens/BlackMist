package services

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

func GetTorIP() (string, error) {
	proxyURL, err := url.Parse("socks5://127.0.0.1:9050")
	if err != nil {
		return "", fmt.Errorf("invalid proxy URL: %w", err)
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
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

func GetTorLatency() (int64, error) {
	proxyURL, err := url.Parse("socks5://127.0.0.1:9050")
	if err != nil {
		return 0, fmt.Errorf("invalid proxy URL: %w", err)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	start := time.Now()

	resp, err := client.Head("https://check.torproject.org")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return time.Since(start).Milliseconds(), nil
}

func WaitForBootstrap(timeoutSec int) error {
	deadline := time.Now().Add(time.Duration(timeoutSec) * time.Second)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", "127.0.0.1:9050", 1*time.Second)
		if err == nil {
			conn.Close()
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return fmt.Errorf("Tor did not start within %ds", timeoutSec)
}
