// file: datamodels/movie.go

package viewmodels

import (
	"gopkg.in/mgo.v2/bson"
)

// Movie is our sample data structure.
// Keep note that the tags for public-use (for our web app)
// should be kept in other file like "web/viewmodels/movie.go"
// which could wrap by embedding the datamodels.Movie or
// declare new fields instead butwe will use this datamodel
// as the only one Movie model in our application,
// for the shake of simplicty.
type Todo struct {
	Id bson.ObjectId `bson:"_id" json:"id"`

	Name string `json:"name"`

	Completed bool `json:"completed"`
}

type TodoList []Todo
