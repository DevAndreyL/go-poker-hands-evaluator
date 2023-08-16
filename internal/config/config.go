package config

import (
	"os"
	"strings"

	"github.com/FZambia/viper-lite"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
)

type httpConfig struct {
	Listen string `mapstructure:"listen"`
}

type Config struct {
	HTTP httpConfig `mapstructure:",squash"`
}

func ReadConfig(interspersed bool) (*Config, error) {
	_ = godotenv.Load() // nolint

	var commandLine = pflag.NewFlagSet("config", pflag.ContinueOnError)
	commandLine.SetInterspersed(interspersed)

	_ = commandLine.StringP("listen", "l", ":80", "HTTP binding address")

	if err := commandLine.Parse(os.Args[1:]); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	if err := viper.BindPFlags(commandLine); err != nil {
		return nil, err
	}

	var (
		config Config
		err    = viper.Unmarshal(&config)
	)

	return &config, err
}
