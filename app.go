package main

import (
	"blackmist/internal/services"
	"context"
	"fmt"
	"os/exec"
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

func (a *App) StartTor() (string, error) {
	if a.active {
		return "Already active", nil
	}

	config, _ := a.engine.WriteConfig()
	a.engine.Cmd = exec.Command(a.engine.BinaryPath, "-f", config)

	if err := a.engine.Cmd.Start(); err != nil {
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
