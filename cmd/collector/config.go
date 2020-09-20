package main

// Config describes configuration for news collector.
type Config struct {
	Redis Redis
}

// Redis configures redis options.
type Redis struct {
	// Addr sets host(or ip) and optional port for redis connection.
	Addr string `env:"REDIS_ADDR,required"`
	// Password sets redis password, optional.
	Password string `env:"REDIS_PASSWORD"`
	// Db sets DB number, optional.
	Db int `env:"REDIS_DB"`
}
