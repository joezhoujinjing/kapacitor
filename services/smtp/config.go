package smtp

import (
	"time"

	"github.com/influxdata/influxdb/toml"
)

type Config struct {
	Enabled  bool   `toml:"enabled"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	// Whether to skip TLS verify.
	NoVerify bool `toml:"no-verify"`
	// Whether all alerts should trigger an email.
	Global bool `toml:"global"`
	// Whether all alerts should automatically use stateChangesOnly mode.
	// Only applies if global is also set.
	StateChangesOnly bool `toml:"state-changes-only"`
	// From address
	From string `toml:"from"`
	// Default To addresses
	To []string `toml:"to"`
	// Close connection to SMTP server after idle timeout has elapsed
	IdleTimeout toml.Duration `toml:"idle-timeout"`
}

func NewConfig() Config {
	return Config{
		Host:        "localhost",
		Port:        25,
		Username:    "",
		Password:    "",
		IdleTimeout: toml.Duration(time.Second * 30),
	}
}
