package config

import "time"

type InfluxConfig struct {
	Url    string `toml:"url"`
	Token  string `toml:"token"`
	Org    string `toml:"org"`
	Bucket string `toml:"bucket"`
}

type Config struct {
	Duration time.Duration     `toml:"duration"`
	Influx   InfluxConfig      `toml:"influx"`
	Servers  map[string]string `toml:"servers"`
}

func NewConfig() Config {
	return Config{
		Duration: time.Second * 10,
		Influx: InfluxConfig{
			Url:    "http://127.0.0.1:8428",
			Token:  "auth-token",
			Org:    "test-org",
			Bucket: "test-bucket",
		},
		Servers: map[string]string{
			"hypixel": "mc.hypixel.net",
		},
	}
}
