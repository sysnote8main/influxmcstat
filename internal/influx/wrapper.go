package influx

import (
	"log/slog"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/sysnote8main/influxmcstat/internal/config"
	"github.com/sysnote8main/influxmcstat/internal/mcping"
)

func NewClient(config config.InfluxConfig) *Client {
	// Create a client and async write api
	influxClient := influxdb2.NewClient(config.ServerUrl, config.Token)
	writeApi := influxClient.WriteAPI(config.Org, config.Bucket)

	// Error handling
	go func() {
		for err := range writeApi.Errors() {
			slog.Error("Failed to write", slog.Any("error", err))
		}
	}()

	// Create wrapped client
	client := Client{
		InfluxClient: influxClient,
		WriteApi:     writeApi,
	}

	// return pointer of client
	return &client
}

type Client struct {
	InfluxClient influxdb2.Client
	WriteApi     api.WriteAPI
}

func (c *Client) Close() {
	c.WriteApi.Flush()
	c.InfluxClient.Close()
}

func (c *Client) WriteMCStat(name string, players int, ping time.Duration) {
	// Create point from minecraft ping result
	p := influxdb2.NewPointWithMeasurement("mcsrv_stats").
		AddTag("name", name).
		AddField("players", players).
		AddField("ping", ping).
		SetTime(time.Now())

	// Write to influxdb
	c.WriteApi.WritePoint(p)
}

func (c *Client) WriteMCStatFromStatus(name string, stat mcping.Status) {
	c.WriteMCStat(name, stat.Players.Online, stat.Delay)
}
