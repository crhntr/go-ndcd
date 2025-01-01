package ndcd

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func Parse(path string, limit int) (map[string]*Product, map[string]*Labeler) {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	reader := bufio.NewReader(file)
	reader.ReadString('\n')

	products := make(map[string]*Product)
	labelers := make(map[string]*Labeler)

	// loading product lines from file

	for linesParsed := 0; linesParsed < limit || limit == -1; linesParsed++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panic(err)
		}
		p := parseLine(line)
		hash := strings.ToLower(p.LabelerName + "<<" + p.ProprietaryName)
		if existingProduct, ok := products[hash]; !ok {
			products[hash] = &p
		} else {
			existingProduct.Variations = append(existingProduct.Variations, p.Variations...)
			products[hash] = existingProduct
		}
	}

	for _, product := range products {
		if labeler, ok := labelers[product.LabelerName]; ok {
			labeler.Products = append(labeler.Products, product)
			labelers[product.LabelerName] = labeler
		} else {
			labelers[product.LabelerName] = &Labeler{Name: product.LabelerName, Products: []*Product{product}}
		}
	}
	return products, labelers
}
