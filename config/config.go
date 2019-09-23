package config

type Config struct {
	Database database
}

type database struct {
	Server string
	Port   string
}
