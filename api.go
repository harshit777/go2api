package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Details struct {
	Name   interface{}
	Adress interface{}
	Image  string
}

func GetGeocodeLocation(s string) []float64 {
	google_api_key := "AIzaSyAcKRO63SCy87HvebSXO0v6Gjp444PwNG8"
	locationString := strings.ReplaceAll(s, " ", "+")
	url := ("https://maps.googleapis.com/maps/api/geocode/json?address=" + locationString + "&key=" + google_api_key)
	response, err := http.Get(url)
	if err != nil {
		panic("Error retreiving response")
	}

	var Coordinates []float64

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
	Coordinates = append(Coordinates, latitude, longitude)
	return Coordinates
}

func findARestaurant(mealType string, location string) []Details {

	var foursquare_client_id = "PQSPJYQ4ODOORBDB51EQTURIFOQT0PACPFQI2UN0G0P00DAF"
	var foursquare_client_secret = "3ZUE0PGPDY2KV4UFPWEQGLZ4GNDWC2PZHFTX40CKZTCIA3LP"

	Coordinates := GetGeocodeLocation(location)
	lat := fmt.Sprintf("%f", Coordinates[0])
	long := fmt.Sprintf("%f", Coordinates[1])
	var AppendData []Details
	var restaurantInfo Details

	url := "https://api.foursquare.com/v2/venues/search?client_id=" + foursquare_client_id + "&client_secret=" + foursquare_client_secret + "&v=20130815&ll=" + lat + "," + long + "&query=" + mealType

	resp, err := http.Get(url)
	if err != nil {
		panic("Error")
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	if err2 != nil {
		panic("Error")
	}
	for _, v := range data["response"].(map[string]interface{}) {

		for _, v2 := range v.([]interface{}) {
			id := v2.(map[string]interface{})["id"].(string)
			rest_name := v2.(map[string]interface{})["name"]
			rest_address := v2.(map[string]interface{})["location"].(map[string]interface{})["formattedAddress"]
			Imgurl := "https://api.foursquare.com/v2/venues/" + id + "/photos?client_id=" + foursquare_client_id + "&v=20150603&client_secret=" + foursquare_client_secret
			imageResp, err3 := http.Get(Imgurl)
			RespBody, err4 := ioutil.ReadAll(imageResp.Body)
			if err3 != nil || err4 != nil {
				panic(err3)
			}
			var ImgData map[string]interface{}
			json.Unmarshal(RespBody, &ImgData)

			var prefix string
			var suffix string
			var imageURL string

			for _, v := range ImgData["response"].(map[string]interface{}) {
				items := v.(map[string]interface{})["items"]
				for _, v := range items.([]interface{}) {
					prefix = v.(map[string]interface{})["prefix"].(string)
					suffix = v.(map[string]interface{})["suffix"].(string)

					if prefix != "" && suffix != "" {
						imageURL = prefix + "300x300" + suffix
					}

				}
			}

			if imageURL == "" {
				imageURL = "https://cdn.pixabay.com/photo/2014/12/21/23/36/hamburgers-575655_1280.png"
			}

			restaurantInfo = Details{rest_name, rest_address, imageURL}

			AppendData = append(AppendData, restaurantInfo)

		}

	}

	return AppendData
}

func FindARestaurant(mealType string, location string) []Details {
	return findARestaurant(mealType, location)
}
