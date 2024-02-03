package api

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	Google_Map_API_KEY string
	Foursquare_API_KEY string
)

func RetrieveAPIKeys() {
	Google_Map_API_KEY = os.Getenv("Google_Map_API_KEY")
	Foursquare_API_KEY = os.Getenv("Foursquare_API_KEY")

	if Google_Map_API_KEY == "" || Foursquare_API_KEY == "" {
		fmt.Println("Please make sure you have all necessary API keys before you proceed")
	}
}
