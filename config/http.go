package config

import "time"

var httpConfig = map[string]string{
	HTTPPort: ":8080",
}

var httpInterfaceConfig = map[string]interface{}{
	IReadTimeout:  5 * time.Minute,
	IWriteTimeout: 5 * time.Minute,
	IIdleTimeout:  30 * time.Second,
	IWaitShutdown: 5 * time.Second,
}

const (
	HTTPPort      = "HTTPPort"
	IReadTimeout  = "IReadTimeout"
	IWriteTimeout = "IWriteTimeout"
	IIdleTimeout  = "IIdleTimeout"
	IWaitShutdown = "IWaitShutdown"
)
