package config

type dbConfig struct {
	Addr     string
	Password string
}

type globalConfig struct {
	Redis dbConfig
	Mysql dbConfig
	Port  string
}

var Config = globalConfig{
	Redis: dbConfig{
		Addr:     "1.117.141.66:6379",
		Password: "123456abc",
	},
	Mysql: dbConfig{
		Addr:     "127.0.0.1:3306",
		Password: "123456",
	},
	Port: "8000",
}
