package nationaldrugcodedirectory

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Labeler struct {
	ID       bson.ObjectId
	Name     string
	Products []*Product
}

func (l Labeler) String() string {
	str := "\n  ["
	for _, l := range l.Products {
		str += "\n     " + l.String()
	}
	str += "\n  ]"

	return fmt.Sprintf("{ %s%s \n}", l.Name, str)
}
