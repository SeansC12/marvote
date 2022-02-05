package config

type AppConfiguration struct {
	ServerCfg ServerConfiguration `mapstructure:"server"`
	MongoCfg  MongoConfiguration  `mapstructure:"mongo,omitempty"`
}

type ServerConfiguration struct {
	Port int `mapstructure:"port"`
}

type MongoConfiguration struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
