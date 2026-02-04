package mcping

import (
	"time"

	"github.com/Tnze/go-mc/chat"
	"github.com/google/uuid"
)

type Icon string

type Status struct {
	Description chat.Message
	Players     struct {
		Max    int
		Online int
		Sample []struct {
			ID   uuid.UUID
			Name string
		}
	}
	Version struct {
		Name     string
		Protocol int
	}
	Favicon Icon
	Delay   time.Duration
}
