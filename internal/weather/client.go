package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	latitude   float64
	longitude  float64
}

func NewClient(timeout time.Duration, latitude, longitude float64) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		latitude:  latitude,
		longitude: longitude,
	}
}

func (c *Client) Fetch(ctx context.Context) (*Weather, error) {
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true",
		c.latitude,
		c.longitude,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var apiResp struct {
		CurrentWeather struct {
			Temperature float64 `json:"temperature"`
		} `json:"current_weather"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return &Weather{
		TemperatureC: apiResp.CurrentWeather.Temperature,
		UpdatedAt:    time.Now().UTC(),
	}, nil
}
