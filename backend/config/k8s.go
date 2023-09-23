//go:build k8s

package config

var Config = config{
	DB: DBConfig{
		DSN: "root:root@tcp(mbook-mysql:3309)/mbook",
	},
	Redis: RedisConfig{
		Addr: "mbook-redis:16379",
	},
}
