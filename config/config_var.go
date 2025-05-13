package config

import "os"

var AUTO_UPGRADE = false

func init() {
	if os.Getenv("AUTO_UPGRADE") == "true" {
		AUTO_UPGRADE = true
	}
}
