package config

import(
    "fmt"
    "github.com/utrack/envconfig"
)

type PSConfig struct {
    Host          string `envconfig:"DB_HOST" default:"postgres" json:"host"`
	DbName        string `envconfig:"DB_NAME" required:"true" json:"dbName"`
	User          string `envconfig:"DB_USER" required:"true" json:"user"`
	Password      string `envconfig:"DB_PASSWORD" required:"true" json:"password"`
}

func GetPSConfig() PSConfig {
    result := PSConfig{}
    err := envconfig.Process("ps", &result)
    if err != nil {
        fmt.Println(err.Error())
    }
    return result
}

type HmacConfig struct {
    PrivateKey    string `envconfig:"PRIVATE_KEY" required:"true"`
}