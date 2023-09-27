package network

import (
	"os/exec"
	"strings"
)

func DiscoverDevices() ([]string, error) {
	cmd := exec.Command("arp", "-a")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	devices := make([]string, 0, len(lines))

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			devices = append(devices, line)
		}
	}

	return devices, nil
}
