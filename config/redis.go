package config

// Redis configures redis options.
type Redis struct {
	// Addr sets host(or ip) and optional port for redis connection.
	Addr string `env:"REDIS_ADDR" env-required:"true"`
	// Password sets redis password, optional.
	Password string `env:"REDIS_PASSWORD"`
	// Db sets DB number, optional.
	Db int `env:"REDIS_DB"`
}
