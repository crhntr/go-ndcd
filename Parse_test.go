package nationaldrugcodedirectory

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	products, labelers := Parse("testdata/product.txt", -1)
	fmt.Printf("products: %d, labelers: %d\n", len(products), len(labelers))
}
