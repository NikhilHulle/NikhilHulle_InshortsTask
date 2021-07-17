package Structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// These are the structures necessary for the code

type Stat struct {
	Loc              string `json:"loc"`
	Total            int    `json:"total"`
	Deaths           int    `json:"deaths"`
	Discharged       int    `json:"discharged"`
	LastOriginUpdate string `bson:"Created_at" json:"lastOriginUpdate,omitempty"`
	LastRefreshed    string `bson:"Updated_at" json:"lastRefreshed,omitempty"`
}

type Note struct {
	ID               primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Loc              string             `json:"Location"`
	Total            int                `json:"TotalCases"`
	Deaths           int                `json:"Deaths"`
	Discharged       int                `json:"Discharged"`
	LastOriginUpdate string             `bson:"Created_at" json:"lastOriginUpdate,omitempty"`
	LastRefreshed    string             `bson:"Updated_at" json:"lastRefreshed,omitempty"`
}
