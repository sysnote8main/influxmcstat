package config

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
	Influx  InfluxConfig
	Servers map[string]ServerConfig
}

func NewConfig() Config {
	return Config{
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
