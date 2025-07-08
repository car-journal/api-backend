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
	HTTPPort      = "HTTP_PORT"
	IReadTimeout  = "I_READ_TIMEOUT"
	IWriteTimeout = "I_WRITE_TIMEOUT"
	IIdleTimeout  = "I_IDLE_TIMEOUT"
	IWaitShutdown = "I_WAIT_SHUTDOWN"
)
