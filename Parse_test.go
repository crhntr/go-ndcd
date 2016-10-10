package nationaldrugcodedirectory_test

import (
	"testing"

	ndcd "github.com/crhntr/nationaldrugcodedirectory"
)

func TestParse(t *testing.T) {
	products, labelers := ndcd.Parse("testdata/product.txt", -1)
	// fmt.Printf("products: %d, labelers: %d\n", len(products), len(labelers))

	for _, value := range labelers {
		str := value.String()
		if str == "" {
			t.Fail()
		}
		break
	}
	for _, value := range products {
		str := value.String()
		if str == "" {
			t.Fail()
		}
		break
	}
}
