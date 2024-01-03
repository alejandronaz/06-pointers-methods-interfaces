package main

import "fmt"

func factory(typeProd ProductType, price float64) Product {
	switch typeProd {
	case SMALL:
		return SmallProduct{price}
	case MEDIUM:
		return MediumProduct{price}
	case LARGE:
		return LargeProduct{price}
	default:
		return nil
	}
}

func main() {

	var products []Product
	products = append(products, factory(SMALL, 100.0))
	products = append(products, factory(MEDIUM, 100.0))
	products = append(products, factory(LARGE, 100.0))

	for _, prod := range products {
		fmt.Println(prod.Price())
	}

}
