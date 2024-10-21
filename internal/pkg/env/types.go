package env

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

type ENV struct {
	Server Server `yaml:"server"`
	Redis  Redis  `yaml:"redis"`
	Crypto Crypto `yaml:"crypto"`
}
