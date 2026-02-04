package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sysnote8main/influxmcstat/internal/config"
	"github.com/sysnote8main/influxmcstat/internal/influx"
	"github.com/sysnote8main/influxmcstat/internal/mcping"
)

func main() {
	config := config.LoadConfig("config.toml")
	influxConfig := config.Influx

	client := influx.NewClient(influxConfig)
	defer client.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Start sending metrics... (Press Ctrl+C to stop)")

	breakableLoop(sigChan, func() {
		for name, server := range config.Servers {
			if !server.Enabled {
				continue
			}

			status, err := mcping.Ping(server.Address)
			if err != nil {
				slog.Warn("Failed to get server info", slog.String("server_name", name))
				client.WriteMCStat(name, -1, 0)
				continue
			}

			// Write minecraft server status
			client.WriteMCStatFromStatus(name, *status)
		}
	})

	fmt.Println("See you!")
}

func breakableLoop(stopSignal chan os.Signal, onTick func()) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			onTick()
		case <-stopSignal:
			return
		}
	}
}
