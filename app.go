package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type App struct {
	ctx      context.Context
	isActive bool
	torPath  string
	torCmd   *exec.Cmd
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.GetOS()
}

func (a *App) GetOS() string {
	osName := runtime.GOOS
	binaryName := "tor"
	if osName == "windows" {
		binaryName = "tor.exe"
	}
	path, _ := filepath.Abs(filepath.Join("bin", osName, binaryName))
	a.torPath = path

	switch osName {
	case "darwin":
		exec.Command("xattr", "-d", "com.apple.quarantine", a.torPath).Run()
		os.Chmod(a.torPath, 0755)
	case "linux":
		os.Chmod(a.torPath, 0755)
	}
	fmt.Printf("[SYSTEM] Engine target: %s | Path %s\n", osName, a.torPath)
	return osName
}

func (a *App) writeTorConfig() (string, error) {
	tempDir := os.TempDir()
	dataDir := filepath.Join(tempDir, "blackmist_tor_data")
	configPath := filepath.Join(tempDir, "torrc_blackmist")

	os.MkdirAll(dataDir, 0700)

	cleanDataDir := filepath.ToSlash(dataDir)

	configContent := fmt.Sprintf(
		"SocksPort 9050\n"+
			"ControlPort 9051\n"+
			"DataDirectory \"%s\"\n"+
			"CookieAuthentication 1\n",
		cleanDataDir,
	)

	err := os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		return "", err
	}
	return configPath, nil
}

func (a *App) StartTor() (string, error) {
	fmt.Printf("[DEBUG] Trying to launch: %s\n", a.torPath)

	if a.isActive {
		return "Tor is already running", nil
	}

	if a.torPath == "" {
		a.GetOS()
	}

	configPath, err := a.writeTorConfig()
	if err != nil {
		return "", fmt.Errorf("Impossible to generate config : %w", err)
	}

	a.torCmd = exec.Command(a.torPath, "-f", configPath)

	if runtime.GOOS == "linux" {
		binDir := filepath.Dir(a.torPath)
		a.torCmd.Env = append(os.Environ(), "LD_LIBRARY_PATH="+binDir)
	}

	a.torCmd.Dir = filepath.Dir(a.torPath)
	a.torCmd.Stdout = os.Stdout
	a.torCmd.Stderr = os.Stderr
	err = a.torCmd.Start()
	if err != nil {
		return "", fmt.Errorf("Error while launching binary : %w", err)
	}
	a.isActive = true
	fmt.Printf("[SUCCESS] Successfully launched Tor with config: %s\n", configPath)
	return "Succesfully started Tor", nil
}

func (a *App) GetTorIP() (string, error) {
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
		return "", fmt.Errorf("Error reading response: %w", err)
	}
	finalIP := string(ipBytes)
	fmt.Printf("[NETWORK] Tor IP Identified: %s\n", finalIP)
	return finalIP, nil
}
