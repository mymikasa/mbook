//go:build !k8s

package config

var Config = config{
	DB: DBConfig{
		DSN: "localhost:3309",
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}
