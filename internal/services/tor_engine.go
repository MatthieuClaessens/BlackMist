package services

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type TorEngine struct {
	BinaryPath string
	Cmd        *exec.Cmd
}

func (e *TorEngine) Init() error {
	osName := runtime.GOOS
	binName := "tor"
	if osName == "windows" {
		binName = "tor.exe"
	}

	ex, err := os.Executable()
	if err != nil {
		e.BinaryPath = filepath.Join("bin", osName, binName)
	} else {
		executableDir := filepath.Dir(ex)
		projectRoot := filepath.Join(executableDir, "..", "..")
		e.BinaryPath = filepath.Clean(filepath.Join(projectRoot, "bin", osName, binName))
	}

	info, err := os.Stat(e.BinaryPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("binaire tor introuvable : %s", e.BinaryPath)
		}
		return fmt.Errorf("erreur lors de l'accès au binaire : %v", err)
	}

	if info.IsDir() {
		return fmt.Errorf("le chemin spécifié est un dossier : %s", e.BinaryPath)
	}

	if osName != "windows" {
		err := os.Chmod(e.BinaryPath, 0755)
		if err != nil {
			return fmt.Errorf("impossible de rendre le binaire exécutable : %v", err)
		}
	}

	return nil
}

func (e *TorEngine) WriteConfig() (string, error) {
	tempDir := os.TempDir()
	dataDir := filepath.Join(tempDir, "blackmist_tor_data")
	configPath := filepath.Join(tempDir, "torrc_blackmist")

	err := os.MkdirAll(dataDir, 0700)
	if err != nil {
		return "", err
	}

	content := fmt.Sprintf(
		"SocksPort 9050\nControlPort 9051\nDataDirectory \"%s\"\nCookieAuthentication 1\n",
		filepath.ToSlash(dataDir),
	)

	return configPath, os.WriteFile(configPath, []byte(content), 0644)

}

func (e *TorEngine) NewIdentity() error {
	conn, err := net.Dial("tcp", "127.0.0.1:9051")
	if err != nil {
		return fmt.Errorf("connexion au ControlPort impossible : %v", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "AUTHENTICATE \"\"\r\n")
	fmt.Fprintf(conn, "SIGNAL NEWNYM\r\n")
	fmt.Fprintf(conn, "QUIT\r\n")

	return nil
}
