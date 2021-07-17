package geocoder

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	geo "github.com/kellydunn/golang-geo"
	"github.com/labstack/echo"
)

// i have used opencage for reversergeo coding
func GetLocation(c echo.Context) error {
	lat := c.QueryParam("lat")
	lng := c.QueryParam("lng")
	lat1, _ := strconv.ParseFloat(lat, 64)
	lng1, _ := strconv.ParseFloat(lng, 64)

	r := geo.NewPoint(lat1, lng1)
	k := new(geo.OpenCageGeocoder)
	var dist string

	// api key
	geo.SetOpenCageAPIKey("615d1ccf96124f329bed5a98f213d294")
	dist, _ = k.ReverseGeocode(r)
	res2 := strings.Fields(dist)
	fmt.Printf("location:%s\n", res2[6])

	s1 := res2[6]
	if last := len(s1) - 1; last >= 0 && s1[last] == ',' {
		s1 = s1[:last]
	}
	fmt.Println("s1:", s1)

	// using the API_LocationStats API
	uri := "http://localhost:2445/find?location=" + s1
	response, err := http.Get(uri)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	target := string(responseData)

	if err != nil {
		panic(err)
	}

	return c.String(http.StatusOK, target)
}
