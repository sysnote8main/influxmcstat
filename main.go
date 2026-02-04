package main

import (
	"log/slog"

	"github.com/sysnote8main/influxmcstat/internal/config"
	"github.com/sysnote8main/influxmcstat/internal/influx"
	"github.com/sysnote8main/influxmcstat/internal/mcping"
)

func main() {
	config := config.LoadConfig("config.toml")
	influxConfig := config.Influx

	client := influx.NewClient(influxConfig)
	defer client.Close()

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
}
