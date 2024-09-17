package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
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

func main() {
	http.HandleFunc("/clima/",
		func(w http.ResponseWriter, r *http.Request) {
			cidade := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := consulta(cidade)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		})

	http.ListenAndServe(":8080", nil)
}

func consulta(cidade string) (weatherData, error) {
	apiConfig, err := configApi(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?appid={API key}" + apiConfig.ApiKey + "&q=" + cidade + "&units=metric")
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}
