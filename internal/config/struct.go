package config

import "time"

type InfluxConfig struct {
	ServerUrl string
	Token     string
	Org       string
	Bucket    string
}

type ServerConfig struct {
	Address string
	Enabled bool
}

type Config struct {
	TickDuration time.Duration
	Influx       InfluxConfig
	Servers      map[string]ServerConfig
}

func NewConfig() Config {
	return Config{
		TickDuration: time.Second * 10,
		Influx: InfluxConfig{
			ServerUrl: "http://127.0.0.1:8428",
			Token:     "auth-token",
			Org:       "test-org",
			Bucket:    "test-bucket",
		},
		Servers: map[string]ServerConfig{
			"hypixel": {
				Address: "mc.hypixel.net",
				Enabled: true,
			},
		},
	}
}
