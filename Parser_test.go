package nationaldrugcodedirectory_test

import "testing"

func TestParse(t *testing.T) {
	products, labelers := GoNDCD.Parse("./test/product.txt", 100)

	t.Logf("loaded %d products", products.Len())
	t.Logf("loaded %d labelers", labelers.Len())
}
