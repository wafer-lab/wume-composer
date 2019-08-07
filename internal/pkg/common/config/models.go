package config

var paths = []string{
	"./config/config.json",
	"../config/config.json",
	"../../config/config.json",
}

var (
	Core    CoreConfig
	Db      DatabaseConfig
	Auth    AuthorizationConfig
	Logger  LoggerConfig
	Storage StorageConfig
)

type CoreConfig struct {
	Port   string `json:"port"`
	Prefix string `json:"prefix"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

type AuthorizationConfig struct {
	CookieName     string `json:"cookie_name"`
	CookieLifetime uint64 `json:"cookie_lifetime"`
	Secret         string `json:"secret"`
}

type LoggerOut struct {
	Mode     string `json:"mode"`
	Filename string `json:"filename, omitempty"`
}

type LoggerConfig struct {
	Debug   LoggerOut `json:"debug"`
	Info    LoggerOut `json:"info"`
	Warning LoggerOut `json:"warning"`
	Error   LoggerOut `json:"error"`
	Fatal   LoggerOut `json:"fatal"`
}

type StoragePaths struct {
	Dir string `json:"dir"`
	Url string `json:"url"`
}

type StorageConfig struct {
	Avatar StoragePaths `json:"avatar"`
	Trash  StoragePaths `json:"trash"`
}

type File struct {
	Core    CoreConfig          `json:"core"`
	Db      DatabaseConfig      `json:"db"`
	Auth    AuthorizationConfig `json:"auth"`
	Logger  LoggerConfig        `json:"logger"`
	Storage StorageConfig       `json:"storage"`
}

func save(config File) {
	Core = config.Core
	Auth = config.Auth
	Db = config.Db
	Logger = config.Logger
	Storage = config.Storage
}
