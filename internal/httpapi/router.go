package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/andreyvla/weather-infra-demo/internal/weather"
)

type API struct {
	weather *weather.Service
}

func NewRouter(weatherService *weather.Service) http.Handler {
	api := &API{
		weather: weatherService,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", api.healthHandler)
	mux.HandleFunc("/weather", api.weatherHandler)

	return LoggingMiddleware(mux)
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (a *API) weatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wth, err := a.weather.Get(r.Context())
	if err != nil {
		http.Error(w, "failed to fetch weather", http.StatusBadGateway)
		return
	}

	resp := map[string]interface{}{
		"temperature_c": wth.TemperatureC,
		"updated_at":    wth.UpdatedAt,
	}

	json.NewEncoder(w).Encode(resp)
}
