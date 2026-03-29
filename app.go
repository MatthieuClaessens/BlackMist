package main

import (
	"blackmist/internal/services"
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

type App struct {
	ctx    context.Context
	engine *services.TorEngine
	active bool
}

func NewApp() *App {
	return &App{engine: &services.TorEngine{}}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	err := a.engine.Init()
	if err != nil {
		fmt.Printf("CRITICAL ERROR: %v\n", err)
	}
	services.ToggleSystemProxy(false)
}

func (a *App) shutdown(ctx context.Context) {
	a.StopTor()
}

func waitForTor(port string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", "127.0.0.1:"+port, 1*time.Second)
		if err == nil {
			conn.Close()
			return nil
		}
		time.Sleep(500 * time.Millisecond)
	}
	return fmt.Errorf("Tor did not start within the allotted time.")
}

func (a *App) StartTor() (string, error) {
	if a.active {
		return "Already active", nil
	}

	config, err := a.engine.WriteConfig()
	if err != nil {
		return "", fmt.Errorf("erreur config Tor : %w", err)
	}

	a.engine.Cmd = exec.Command(a.engine.BinaryPath, "-f", config)
	a.engine.Cmd.Stdout = os.Stdout
	a.engine.Cmd.Stderr = os.Stderr

	if err := a.engine.Cmd.Start(); err != nil {
		return "", err
	}

	if err := waitForTor("9050", 60*time.Second); err != nil {
		a.engine.Cmd.Process.Kill()
		return "", err
	}

	services.ToggleSystemProxy(true)
	a.active = true
	return "Started Tor", nil
}

func (a *App) StopTor() string {
	services.ToggleSystemProxy(false)
	if a.engine.Cmd != nil && a.engine.Cmd.Process != nil {
		a.engine.Cmd.Process.Kill()
		a.engine.Cmd = nil
	}
	a.active = false
	return "Stopped Tor"
}

func (a *App) CheckIP() (string, error) {
	return services.GetTorIP()
}

func (a *App) GetPing() int64 {
	latency, err := services.GetTorLatency()
	if err != nil {
		return 0
	}
	return latency
}

func (a *App) ChangeIP() (string, error) {
	if err := a.engine.NewIdentity(); err != nil {
		return "", err
	}
	time.Sleep(10 * time.Second)
	return services.GetTorIP()
}
