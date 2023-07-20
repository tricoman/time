package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type GetUtcTimeResponse struct {
	CurrentTime string `json:"current_time"`
}

func getTimeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	if rawTimeZoneValue := request.URL.Query().Get("tz"); rawTimeZoneValue != "" {
		produceResponseWithTimezones(writer, rawTimeZoneValue)
		return
	} else {
		json.NewEncoder(writer).Encode(GetUtcTimeResponse{CurrentTime: time.Now().UTC().String()})
	}
}

func produceResponseWithTimezones(writer http.ResponseWriter, rawTimeZoneValue string) {
	timeZones := strings.Split(rawTimeZoneValue, ",")
	currentTimeByTimeZones := produceTimesWithTimeZones(writer, timeZones)
	json.NewEncoder(writer).Encode(currentTimeByTimeZones)
}

func produceTimesWithTimeZones(writer http.ResponseWriter, timeZones []string) map[string]string {
	var currentTimeByTimeZones = make(map[string]string)
	for _, timeZone := range timeZones {
		location, producedError := time.LoadLocation(timeZone)
		if producedError != nil {
			produceErrorResponse(writer)
			return map[string]string{}
		}
		currentTimeByTimeZones[timeZone] = time.Now().In(location).String()
	}
	return currentTimeByTimeZones
}

func produceErrorResponse(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprint(writer, "invalid timezone")
}
