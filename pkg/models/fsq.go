package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"go2api/cmd/api"
)

type BuildResponse struct {
	Name    string   `json:"name"`
	Address []string `json:"address"`
	Image   string   `json:"image"`
}

const (
	GET string = http.MethodGet
)

var categoriesCode map[string]string = map[string]string{
	"Café":       "13032", // Cafe, Coffee, and Tea House
	"Restaurant": "13065", // Restaurant
}

func FindPlaces(placeType string, location string) []BuildResponse {
	FSQ_API_KEY := api.Foursquare_API_KEY

	var responses []BuildResponse

	fsqSearchBaseURL := "https://api.foursquare.com/v3/places/search"
	req, err := http.NewRequest(GET, fsqSearchBaseURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	latitude, longitude := GetGeocodeLocation(location)
	lat := fmt.Sprintf("%f", latitude)
	long := fmt.Sprintf("%f", longitude)

	queryParams := map[string]string{
		"categories": categoriesCode[placeType],
		"ll":         strings.Join([]string{lat, long}, ","),
		"radius":     "500",
		"sort":       "DISTANCE",
	}

	q := req.URL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", FSQ_API_KEY)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("Error")
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	if err2 != nil {
		panic("Error")
	}

	for _, v := range data["results"].([]interface{}) {
		v2 := v.(map[string]interface{})

		id := v2["fsq_id"].(string)
		rest_name := v2["name"].(string)
		rest_address := v2["location"].(map[string]interface{})["formatted_address"].(string)

		Imgurl := "https://api.foursquare.com/v3/places/" + id + "/photos?limit=1"
		req, _ := http.NewRequest(GET, Imgurl, nil)

		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", FSQ_API_KEY)

		imageResp, err3 := http.DefaultClient.Do(req)
		if err != nil {
			panic("Error")
		}

		RespBody, err4 := ioutil.ReadAll(imageResp.Body)
		RespBody = RespBody[1 : len(RespBody)-1]
		if err3 != nil || err4 != nil {
			panic(err3)
		}

		var ImgData map[string]string

		json.Unmarshal(RespBody, &ImgData)

		var prefix string
		var suffix string
		var imageURL string

		prefix = ImgData["prefix"]
		suffix = ImgData["suffix"]

		if prefix != "" && suffix != "" {
			imageURL = prefix + "300x300" + suffix
		}

		if imageURL == "" {
			imageURL = "http://pixabay.com/get/8926af5eb597ca51ca4c/1433440765/cheeseburger-34314_1280.png?direct"
		}

		restaurantInfo := BuildResponse{
			Name:  rest_name,
			Image: imageURL,
		}

		restaurantInfo.Address = append(restaurantInfo.Address, rest_address)
		responses = append(responses, restaurantInfo)
	}

	fmt.Println("process finished")
	return responses
}
