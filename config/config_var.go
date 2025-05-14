package config

import (
	"os"

	"github.com/spf13/cast"
)

var AUTO_UPGRADE = false
var EVENT_EXPIRE_DAY = 3

func init() {
	if os.Getenv("AUTO_UPGRADE") == "true" {
		AUTO_UPGRADE = true
	}
	if os.Getenv("EVENT_EXPIRE_DAY") != "" {
		EVENT_EXPIRE_DAY = cast.ToInt(os.Getenv("EVENT_EXPIRE_DAY"))
	}
}
