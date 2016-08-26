package nationaldrugcodedirectory

import "fmt"

type Labeler struct {
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
