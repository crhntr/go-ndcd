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

type labelerNode struct {
	l    *Labeler
	next *labelerNode
}
type labelerList struct {
	first  *labelerNode
	length int
}

func (ll *labelerList) Len() int {
	return ll.length
}
func (ll *labelerList) Push(l *Labeler) {
	ll.first = &labelerNode{l: l, next: ll.first}
	ll.length++
}
