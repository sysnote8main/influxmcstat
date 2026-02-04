package mcping

import (
	"encoding/json"

	"github.com/Tnze/go-mc/bot"
)

func Ping(addr string) (*Status, error) {
	resp, delay, err := bot.PingAndList(addr)
	if err != nil {
		return nil, err
	}

	var s Status
	err = json.Unmarshal(resp, &s)
	if err != nil {
		return nil, err
	}
	s.Delay = delay

	return &s, nil
}
