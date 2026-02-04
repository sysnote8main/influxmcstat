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

	breakableLoop(sigChan, config.Duration, func() {
		for name, address := range config.Servers {
			status, err := mcping.Ping(address)
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

func breakableLoop(stopSignal chan os.Signal, tickDuration time.Duration, onTick func()) {
	ticker := time.NewTicker(tickDuration)
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
