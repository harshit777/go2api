package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetGeocodeLocation(s string) {
	google_api_key := "AIzaSyAcKRO63SCy87HvebSXO0v6Gjp444PwNG8"
	locationString := strings.ReplaceAll(s, " ", "+")
	url := ("https://maps.googleapis.com/maps/api/geocode/json?address=" + locationString + "&key=" + google_api_key)
	var values map[string]interface{}
	response, err := http.Get(url)
	if err != nil {
		panic("Error retreiving response")
	}

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		panic("Error retreiving response")
	}

	json.Unmarshal(body, &values)
	for _, v := range values {
		for i2, v2 := range v.(map[string]interface{}) {

			if i2 == "location" {
				fmt.Print(i2, v2.(string))

				fmt.Println()
			}
		}
	}
	fmt.Println("Body Respons %s", values)
}
