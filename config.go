package config

// Configuration type
type Configuration struct {
	HTTP     *HTTP
	MYSQL    *MYSQL
	MONGO    *MONGO
	RABBITMQ *RABBITMQ
	REDIS    *REDIS
	TOKEN    *TOKEN
}

// HTTP configuration
type HTTP struct {
	PORT string
}

type REDIS struct {
	HOST string
}

type MYSQL struct {
	HOST     string
	PORT     string
	DB       string
	USERNAME string
	PASSWORD string
}

type MONGO struct {
	HOST     string
	PORT     string
	DB       string
	USERNAME string
	PASSWORD string
}

type RABBITMQ struct {
	HOST     string
	PORT     string
	USERNAME string
	PASSWORD string
}

type TOKEN struct {
	PUBLIC  string
	PRIVATE string
	AUTH    string
}
