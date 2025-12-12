package main

import (
	"github.com/andreyvla/weather-infra-demo/internal/config"
	"github.com/andreyvla/weather-infra-demo/internal/observability"
	"github.com/rs/zerolog/log"
)

func main() {
	observability.InitLogger()

	log.Info().Msg("starting weather service")

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	log.Info().
		Int("port", cfg.Port).
		Float64("latitude", cfg.Latitude).
		Float64("longitude", cfg.Longitude).
		Dur("cache_ttl", cfg.WeatherCacheTTL).
		Dur("weather_timeout", cfg.WeatherTimeout).
		Msg("config loaded successfully")
}
