package main

import (
	"blackmist/internal/services"
	"context"
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
	a.engine.Init()
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
	}
	a.active = false
	return "Stopped Tor"
}

func (a *App) CheckIP() (string, error) {
	return services.GetTorIP()
}
