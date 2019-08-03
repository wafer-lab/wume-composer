package config

import (
	"path/filepath"
)

var dirs = []string{
	".",
	"..",
	".." + string(filepath.Separator) + "..",
}

var (
	Core CoreConfig
	Db   DatabaseConfig
	Auth AuthorizationConfig
)

type CoreConfig struct {
	Port string `json:"port"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

type AuthorizationConfig struct {
	CookieName string `json:"cookie_name"`
	Secret     string `json:"secret"`
}

type File struct {
	Core CoreConfig          `json:"core"`
	Db   DatabaseConfig      `json:"db"`
	Auth AuthorizationConfig `json:"auth"`
}

func save(config File) {
	Core = config.Core
	Auth = config.Auth
	Db = config.Db
}
