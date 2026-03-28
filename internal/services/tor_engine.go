package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type TorEngine struct {
	BinaryPath string
	Cmd        *exec.Cmd
}

func (e *TorEngine) Init() {
	osName := runtime.GOOS
	binName := "tor"
	if osName == "windows" {
		binName = "tor.exe"
	}
	path, _ := filepath.Abs(filepath.Join("bin", osName, binName))
	e.BinaryPath = path

	if osName != "windows" {
		os.Chmod(e.BinaryPath, 0755)
	}
}

func (e *TorEngine) WriteConfig() (string, error) {
	tempDir := os.TempDir()
	dataDir := filepath.Join(tempDir, "blackmist_tor_data")
	configPath := filepath.Join(tempDir, "torrc_blackmist")
	os.MkdirAll(dataDir, 0700)

	content := fmt.Sprintf(
		"SocksPort 9050\nControlPort 9051\nDataDirectory \"%s\"\n",
		filepath.ToSlash(dataDir),
	)
	return configPath, os.WriteFile(configPath, []byte(content), 0644)
}
