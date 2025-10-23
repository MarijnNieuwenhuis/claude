package app

const (
	Dev     Environment = "dev"
	Stage   Environment = "stage"
	Acc     Environment = "acc"
	Sandbox Environment = "sandbox"
	Prod    Environment = "prod"
)

type Environment string

type Configuration struct {
	Environment Environment
	LogLevel    string
	HTTPPort    string
	SentryDSN   string
	DatabaseDSN string
	Pubsub      pubsubConfig
}

type pubsubConfig struct {
	Emulator string
	Project  string
}
