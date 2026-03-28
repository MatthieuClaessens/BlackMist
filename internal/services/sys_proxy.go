package services

import (
	"os/exec"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func ToggleSystemProxy(enable bool) error {
	switch runtime.GOOS {
	case "windows":
		return setWindowsProxy(enable)
	case "darwin":
		return setMacOSProxy(enable)
	case "linux":
		return setLinuxProxy(enable)
	}
	return nil
}

func setWindowsProxy(enable bool) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	if enable {
		k.SetDWordValue("ProxyEnable", 1)
		k.SetStringValue("ProxyServer", "socks=127.0.0.1:9050")
	} else {
		k.SetDWordValue("ProxyEnable", 0)
	}
	exec.Command("Rundll32.exe", "user32.dll,UpdatePerUserSystemParameters").Run()
	return nil
}

func setMacOSProxy(enable bool) error { return nil }
func setLinuxProxy(enable bool) error { return nil }
