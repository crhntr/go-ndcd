package nationaldrugcodedirectory

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Parse(logger *log.Logger, path string, limit int) (productList, labelerList) {

	file, err := os.Open(path)
	if err != nil {
		logger.Panic(err)
	}
	reader := bufio.NewReader(file)
	reader.ReadString('\n')

	products := productList{}
	labelers := labelerList{}

	// loading product lines from file
	for products.length < limit || limit == -1 {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			logger.Panic(err)
		}
		p := parseLine(line)
		products.Push(&p)
	}
	if products.first == nil {
		logger.Panic("no products loaded")
	}

	//for p := products.first; p != nil; p = p.next{
	//	fmt.Printf("| %s ", p.p.ProprietaryName)
	//}

	duplicateProducts := 0
	firstProduct := products.first

	for firstProduct != nil && firstProduct.next != nil {
		potentialDuplicate := firstProduct

		for potentialDuplicate.next != nil {
			if firstProduct.p.ProprietaryName == potentialDuplicate.next.p.ProprietaryName {
				duplicateProducts++
				firstProduct.p.Variations = append(firstProduct.p.Variations, potentialDuplicate.next.p.Variations...)
				potentialDuplicate.next = potentialDuplicate.next.next
				products.length--
			} else {
				potentialDuplicate = potentialDuplicate.next
			}
		}

		firstProduct = firstProduct.next
	}

	for p := products.first; p != nil; p = p.next {
		l := &Labeler{Name: p.p.LabelerName, Products: []*Product{p.p}}
		p.p.Labeler = l
		labelers.Push(l)
	}

	firstLabeler := labelers.first
	duplicateLabelers := 0

	for firstLabeler != nil && firstLabeler.next != nil {
		potentialDuplicate := firstLabeler

		for potentialDuplicate.next != nil {
			if firstLabeler.l.Name == potentialDuplicate.next.l.Name {
				duplicateLabelers++
				firstLabeler.l.Products = append(firstLabeler.l.Products, potentialDuplicate.next.l.Products...)
				potentialDuplicate.next = potentialDuplicate.next.next
				labelers.length--
			} else {
				potentialDuplicate = potentialDuplicate.next
			}
		}
		firstLabeler = firstLabeler.next
	}

	//for l := labelers.first; l != nil; l = l.next{
	//	fmt.Println(*l.l)
	//}

	logger.Printf("duplicates removed | products: %7d ; labelers: %7d\n", duplicateProducts, duplicateLabelers)
	logger.Printf("total items parsed | products: %7d ; labelers: %7d\n", products.length, labelers.length)

	return products, labelers
}
