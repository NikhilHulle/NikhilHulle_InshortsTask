package main

import (
	"API_CovidSTATS/pkg/geocoder"

	"github.com/labstack/echo"
)

// @title API_CovidSTATS API
// @version 1.0
// @description This is a simple API that returns the covid statistics of the country(INDIA) and of the location determined by the co-ordinated entered by the user.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1453
// @BasePath /
// @schemes http

// http://localhost:1453/location?lat=18.3817844&lng=77.11884189999999   -----> This is an example of an api call, please make the call through your browzer in this manner
func main() {
	e := echo.New()

	// Routes
	e.GET("/location", geocoder.GetLocation)

	// Start server
	e.Logger.Fatal(e.Start(":1453"))

}
