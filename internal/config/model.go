package config

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type Redis struct {
	Address  string
	Password string
}
