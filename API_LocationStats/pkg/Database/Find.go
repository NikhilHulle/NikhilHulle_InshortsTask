package Database

import (
	"context"
	"fmt"
	"net/http"

	"API_LocationStats/pkg/Structs"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Here i find the covid stats for india and the location which is determined by the coordinates given by the user that is passed as a parameter during the api call

func FindStats(c echo.Context) error {

	location := c.QueryParam("location")
	ctx := context.Background()
	// Options to the database.
	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts) // making connection to the Database
	if err != nil {
		fmt.Println(err)
		return err
	}

	db := client.Database(DBName)
	coll := db.Collection(notesCollection)
	err = coll.Drop(ctx)
	if err != nil {
		coll = db.Collection(notesCollection)
	}
	AddInDB()
	ResultState := coll.FindOne(ctx, bson.M{"loc": location}) // querying the database
	if err := ResultState.Err(); err != nil {
		fmt.Println(err)
		return err
	}
	n := Structs.Note{}
	err = ResultState.Decode(&n)
	if err != nil {
		fmt.Println(err)
		return err
	}

	ResultCountry := coll.FindOne(ctx, bson.M{"loc": "India"})
	if err1 := ResultCountry.Err(); err1 != nil {
		fmt.Println(err1)
		return err1
	}
	n1 := Structs.Note{}
	err1 := ResultCountry.Decode(&n1)
	if err1 != nil {
		fmt.Println(err1)
		return err1
	}

	res := []Structs.Stat{}

	target1 := Structs.Stat{
		Total:            n1.Total,
		Loc:              n1.Loc,
		Deaths:           n1.Deaths,
		Discharged:       n1.Discharged,
		LastRefreshed:    n1.LastRefreshed,
		LastOriginUpdate: n1.LastOriginUpdate,
	}
	res = append(res, target1)

	target := Structs.Stat{
		Total:            n.Total,
		Loc:              n.Loc,
		Deaths:           n.Deaths,
		Discharged:       n.Discharged,
		LastRefreshed:    n1.LastRefreshed,
		LastOriginUpdate: n1.LastOriginUpdate,
	}
	res = append(res, target)

	return c.JSON(http.StatusAccepted, res) // returning string that holds the required information

}
