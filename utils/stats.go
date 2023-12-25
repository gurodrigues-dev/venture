package utils

import (
	"time"

	"github.com/shirou/gopsutil/host"
)

func GetUptime() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}

	uptime := time.Duration(info.Uptime) * time.Second
	return uptime.String(), nil
}
