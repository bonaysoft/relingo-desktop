package system

type Config struct {
	DBPath   string `toml:"dbPath"`
	EngraURL string `toml:"engraUrl"`
}

func NewConfig() *Config {
	return &Config{
		DBPath:   "relingo.db",
		EngraURL: "https://engra.saltbo.cn/query",
	}
}
