package main

import "fmt"

func main() {
	lat, long := GetGeocodeLocation("Dallas, Texas")
	fmt.Println(lat, ",", long)
}
