package location

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetGeocodeLocation(s string) (float64, float64) {

	locationString := strings.ReplaceAll(s, " ", "+")
	url := ("https://maps.googleapis.com/maps/api/geocode/json?address=" + locationString + "&key=" + google_api_key)
	response, err := http.Get(url)
	if err != nil {
		panic("Error retreiving response")
	}

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		panic("Error retreiving response")
	}

	var longitude float64
	var latitude float64
	var values map[string]interface{}

	json.Unmarshal(body, &values)
	for _, v := range values["results"].([]interface{}) {
		for i2, v2 := range v.(map[string]interface{}) {
			if i2 == "geometry" {
				latitude = v2.(map[string]interface{})["location"].(map[string]interface{})["lat"].(float64)
				longitude = v2.(map[string]interface{})["location"].(map[string]interface{})["lng"].(float64)
				break
			}
		}
	}
	fmt.Println("Longituude and latitude", latitude, longitude)
	return latitude, longitude
}
