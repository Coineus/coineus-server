package config

type DatabaseConfiguration struct {
	Name     string `default:"coineus"`
	Username string `default:"user"`
	Password string `default:"password"`
	Host     string `default:"localhost"`
	Port     string `default:"5432"`
	LogMode  bool   `default:"false"`
}
