package main

import (
	"encoding/json"
	"os"
)

type apiConfigData struct {
	ApiKey string `json:"ApiKeyWeather"`
}

type weatherData struct {
	Nome string  `json:"nome"`
	Temp float64 `json:"temp"`
}

func configApi(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}

	var api apiConfigData
	err = json.Unmarshal(bytes, &api)
	if err != nil {
		return apiConfigData{}, err
	}

	return api, nil
}
