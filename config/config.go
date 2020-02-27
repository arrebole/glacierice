package config

import "flag"

// ...
var (
	ListenHost string
	ListenPort string
	DbPassword string
)

func init() {
	flag.StringVar(&ListenHost, "host", "127.0.0.1", "http listen host")
	flag.StringVar(&ListenPort, "port", "3000", "http listen port")
	flag.StringVar(&DbPassword, "password", "", "database password")
	flag.Parse()
}
