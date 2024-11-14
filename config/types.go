package config

type Server struct {
	Port string `yaml:"port"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Crypto struct {
	Key string `yaml:"key"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
	SSLMode  bool   `yaml:"sslMode"`
}

type HTTPAuth struct {
	JWTSecretKey string `yaml:"jwt_secret_key"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Redis    Redis    `yaml:"redis"`
	Crypto   Crypto   `yaml:"crypto"`
	Postgres Postgres `yaml:"postgres"`
	HTTPAuth HTTPAuth `yaml:"http_auth"`
}
