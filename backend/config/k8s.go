//go:build k8s

package config

var Config = config{
	DB: DBConfig{
		DSN: "root:root@tcp(mbook-mysql:3308)/mbook",
	},
	Redis: RedisConfig{
		Addr: "mbook-redis:6380",
	},
}
