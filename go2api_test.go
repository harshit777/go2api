package go2api_test

import (
	"testing"

	"github.com/harshit777/go2api"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFindARestaurant(t *testing.T) {
	FindARestaurant := go2api.FindARestaurant
	expected := `[{"meta":{"code":200},"response":{"venues":[{"id":"533c1b8a498e7a99cde2ec22","name":"Momo's point","contact":{},"location":{"lat":26.847177540263544,"lng":80.93591544655929,"labeledLatLngs":[{"label":"display","lat":26.847177540263544,"lng":80.93591544655929}],"distance":1019,"cc":"IN","country":"India","formattedAddress":["India"]},"categories":[{"id":"4bf58dd8d48988d16e941735","name":"Fast Food Restaurant","pluralName":"Fast Food Restaurants","shortName":"Fast Food","icon":{"prefix":"https:\/\/ss3.4sqi.net\/img\/categories_v2\/food\/fastfood_","suffix":".png"},"primary":true}],"verified":false,"stats":{"tipCount":0,"usersCount":0,"checkinsCount":0,"visitsCount":0},"beenHere":{"count":0,"lastCheckinExpiredAt":0,"marked":false,"unconfirmedCount":0},"hereNow":{"count":0,"summary":"Nobody here","groups":[]},"referralId":"v-1609489423","venueChains":[],"hasPerk":false},{"id":"562d0109498e33b9f939993f","name":"Nainital Momos","contact":{},"location":{"lat":26.85237,"lng":81.003384,"labeledLatLngs":[{"label":"display","lat":26.85237,"lng":81.003384}],"distance":5717,"cc":"IN","country":"India","formattedAddress":["India"]},"categories":[{"id":"4bf58dd8d48988d10f941735","name":"Indian Restaurant","pluralName":"Indian Restaurants","shortName":"Indian","icon":{"prefix":"https:\/\/ss3.4sqi.net\/img\/categories_v2\/food\/indian_","suffix":".png"},"primary":true}],"verified":false,"stats":{"tipCount":0,"usersCount":0,"checkinsCount":0,"visitsCount":0},"beenHere":{"count":0,"lastCheckinExpiredAt":0,"marked":false,"unconfirmedCount":0},"hereNow":{"count":0,"summary":"Nobody here","groups":[]},"referralId":"v-1609489423","venueChains":[],"hasPerk":false}]}}

	]`
	Convey("Should convert correctly", t, func() {
		Convey("Small Number should convert correctly",
			func() {
				So(FindARestaurant("Momos", "Lucknow"), ShouldContain, expected)

			})

	})
}
