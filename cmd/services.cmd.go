package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/mati/latencia/schema"
)

func Classify(signal schema.Task) ([]string, []string) {
	var emergencies []string
	var warnings []string

	switch {
	case signal.Body.Speed > 100:
		emergencies = append(emergencies, "high_speed")
	case signal.Body.Speed > 60:
		warnings = append(warnings, "low_speed")
	}

	switch {
	case signal.Body.Temperature > 50:
		emergencies = append(emergencies, "high_temperature")
	case signal.Body.Temperature > 30:
		warnings = append(warnings, "low_temperature")
	}

	switch {
	case signal.Body.StopsNormal > 20:
		emergencies = append(emergencies, "high_stops_normal")
	case signal.Body.StopsNormal > 10:
		warnings = append(warnings, "low_stops_normal")
	}

	switch {
	case signal.Body.StopsIregular > 20:
		emergencies = append(emergencies, "high_stops_iregular")
	case signal.Body.StopsIregular > 10:
		warnings = append(warnings, "low_stops_iregular")
	}
	return emergencies, warnings
}

func Endpoints(signal schema.Task) []string {
	var endpoints []string
	if signal.Emergency != nil {
		police := "https://44ace8b9-e8a7-4771-8afe-3e9c9620fdc8.mock.pstmn.io/police"
		emergency := "https://44ace8b9-e8a7-4771-8afe-3e9c9620fdc8.mock.pstmn.io/emergency"
		endpoints = append(endpoints, police, emergency)
	}
	if signal.Warning != nil {
		owner := "https://44ace8b9-e8a7-4771-8afe-3e9c9620fdc8.mock.pstmn.io/owner"
		endpoints = append(endpoints, owner)
	}
	return endpoints
}

func SendNotify(endpoint string, signal schema.Task) {
	signal_body, _ := json.Marshal(signal.Body)
	json_body := []string{string(signal_body)}

	http_body := url.Values{"body": json_body}
	resp, err := http.PostForm(endpoint, http_body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(body, err)
	}
}

func Save(signal schema.Task) {
	if signal.Elapsed > 100*time.Millisecond {
		fmt.Println("Save", signal.Elapsed)
	}
	err := InsertItem(schema.TableStruct{
		Id:            uuid.New().String(),
		CreationDate:  time.Now().Format(time.DateTime),
		Longitude:     signal.Body.Latitude,
		Latitude:      signal.Body.Latitude,
		Speed:         signal.Body.Speed,
		Temperature:   signal.Body.Temperature,
		Address:       signal.Body.Address,
		StopsNormal:   signal.Body.StopsNormal,
		StopsIregular: signal.Body.StopsIregular,
		Camera:        signal.Body.Camera,
		State:         signal.Body.State,
		Panic:         signal.Body.Panic,
		Elapsed:       int64(signal.Elapsed),
		Warnings:      signal.Warning,
		Emergencies:   signal.Emergency,
	})
	if err != nil {
		fmt.Println("error:", err)
	}
}
