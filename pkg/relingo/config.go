package relingo

type Config struct {
	BaseURL string `toml:"baseUrl"`
	Token   string `toml:"token"`
}

func NewConfig() *Config {
	return &Config{BaseURL: "https://api.relingo.net/api"}
}
