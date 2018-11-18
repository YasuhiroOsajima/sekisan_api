package config

import "os"

var (
	Port   = getEnv("APP_PORT", "5000")
	//dbhost = getEnv("DB_HOST", "127.0.0.1")
	//dbport = getEnv("DB_PORT", "3306")
	//dbuser = getEnv("DB_USER", "root")
	//dbpass = getEnv("DB_PASSWORD", "")
	//dbname = getEnv("DB_NAME", "testdb")
	//public = getEnv("PUBLIC_DIR", "public")
)

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}
