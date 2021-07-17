package Database

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"API_LocationStats/pkg/Structs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// global variables
var state string
var total, discharged, deaths int

// This function makes a get request call to https://api.rootnet.in/covid19-in/stats/latest and consequently stores the data in the mongodb database.
// I have used mongoDB atlas.
func AddInDB() {

	ctx := context.Background()
	// Options to the database.
	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts) // making connection to the Database
	if err != nil {
		fmt.Println(err)
		return
	}
	db := client.Database(DBName)
	coll := db.Collection(notesCollection)
	fmt.Println(coll.Name())

	// The get request call
	response, err := http.Get("https://api.rootnet.in/covid19-in/stats/latest")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(responseData), &data)

	if err != nil {
		panic(err)
	}

	// reading required fields from the response
	items_india := data["data"].(map[string]interface{})["summary"]

	for key1, value1 := range items_india.(map[string]interface{}) {

		if key1 == "confirmedCasesIndian" {

			total = int(value1.(float64))
		}

		if key1 == "discharged" {

			discharged = int(value1.(float64))

		}

		if key1 == "deaths" {

			deaths = int(value1.(float64))

		}

		state = "India"
	}
	note := Structs.Note{}
	note.ID = primitive.NewObjectID()
	note.Loc = state
	note.Total = total
	note.Deaths = deaths
	note.Discharged = discharged

	str := data["lastRefreshed"].(string)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		fmt.Println(err)
	}
	RefreshTime := t.Format("2006-01-02 15:04:05")
	str = data["lastOriginUpdate"].(string)
	t, err = time.Parse(time.RFC3339, str)
	if err != nil {
		fmt.Println(err)
	}
	OriginTime := t.Format("2006-01-02 15:04:05")
	note.LastRefreshed = RefreshTime
	note.LastOriginUpdate = OriginTime
	result, err := coll.InsertOne(ctx, note) // storing the entry in the database
	if err != nil {
		fmt.Println(err)
		return
	}
	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)

	items := data["data"].(map[string]interface{})["regional"]

	for key, value := range items.([]interface{}) {
		fmt.Println(key)
		for key1, value1 := range value.(map[string]interface{}) {

			if key1 == "confirmedCasesIndian" {

				total = int(value1.(float64))
			}

			if key1 == "discharged" {

				discharged = int(value1.(float64))

			}

			if key1 == "deaths" {

				deaths = int(value1.(float64))

			}
			if key1 == "loc" {

				state = string(value1.(string))

			}

			// An ID for MongoDB.
		}
		note := Structs.Note{}
		note.ID = primitive.NewObjectID()
		note.Loc = state
		note.Total = total
		note.Deaths = deaths
		note.Discharged = discharged
		note.LastRefreshed = RefreshTime
		note.LastOriginUpdate = OriginTime

		result, err := coll.InsertOne(ctx, note) // storing the entry in the database
		if err != nil {
			fmt.Println(err)
			return
		}
		// ID of the inserted document.
		objectID := result.InsertedID.(primitive.ObjectID)
		fmt.Println(objectID)

	}

}
