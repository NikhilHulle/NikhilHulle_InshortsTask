package main

import (
	"API_LocationStats/pkg/Database"

	"github.com/labstack/echo"
)

// @title API_LocationStats API
// @version 1.0
// @description This is a simple API that stores the covid statistics of the country(INDIA) and returns the stats of the location determined by the co-ordinated entered by the user.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:2445
// @BasePath /
// @schemes http

func main() {

	e := echo.New()

	// Routes
	e.GET("/find", Database.FindStats)

	// Start server
	e.Logger.Fatal(e.Start(":2445"))

}
