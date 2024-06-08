package config

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
	"guayavita.dev/gvc/internal/logging"
	"os"
)

var GVConfig GVCConfig

func init() {
	viper.SetConfigName("gvc")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config/.gvc/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	logging.Log.Debug().Msg("Reading configuration file")
	if err != nil {
		logging.Log.Warn().Msg("No configuration file found Using default configuration")
		GVConfig = buildDefaultConfig()
	} else {
		err = viper.Unmarshal(&GVConfig)
		if err != nil {
			logging.Log.Error().Msgf("Unable to decode into struct, %v", err)
		}
	}
	logging.Log.Debug().Msgf("Configuration loaded: %+v", GVConfig)
}

func DumpDefaultConfig() {
	file, err := os.Create("gvc.toml")
	if err != nil {
		logging.Log.Panic().Msgf("Unable to create configuration file: %v", err)
	}
	defer file.Close()
	encoder := toml.NewEncoder(file)
	err = encoder.Encode(GVConfig)
	if err != nil {
		logging.Log.Panic().Msgf("Unable to write configuration file: %v", err)
	}
	logging.Log.Info().Msg("Default configuration file created")
}

func buildDefaultConfig() GVCConfig {
	return GVCConfig{
		Core: Core{
			Verbose: false,
		},
	}
}
