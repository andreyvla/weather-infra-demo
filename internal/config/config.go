package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port            int
	Latitude        float64
	Longitude       float64
	WeatherCacheTTL time.Duration
	WeatherTimeout  time.Duration
}

func Load() (*Config, error) {
	cfg := &Config{}

	// PORT
	portStr := getEnv("PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.New("invalid PORT")
	}
	cfg.Port = port

	// LATITUDE
	latStr := getEnv("LATITUDE", "6.93")
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return nil, errors.New("invalid LATITUDE")
	}
	cfg.Latitude = lat

	// LONGITUDE
	lonStr := getEnv("LONGITUDE", "79.85")
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return nil, errors.New("invalid LONGITUDE")
	}
	cfg.Longitude = lon

	// WEATHER_CACHE_TTL
	cacheTTLStr := getEnv("WEATHER_CACHE_TTL", "60s")
	cacheTTL, err := time.ParseDuration(cacheTTLStr)
	if err != nil {
		return nil, errors.New("invalid WEATHER_CACHE_TTL")
	}
	cfg.WeatherCacheTTL = cacheTTL

	// WEATHER_TIMEOUT
	timeoutStr := getEnv("WEATHER_TIMEOUT", "2s")
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return nil, errors.New("invalid WEATHER_TIMEOUT")
	}
	cfg.WeatherTimeout = timeout

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
