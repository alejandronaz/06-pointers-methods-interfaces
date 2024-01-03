package main

import "fmt"

func factory(typeProd ProductType, price float64) Product {
	switch typeProd {
	case SMALL:
		return SmallProduct{price} // SmallProduct "implementa" la interface Product
	case MEDIUM:
		return MediumProduct{price}
	case LARGE:
		return LargeProduct{price}
	default:
		return nil
	}
}

func main() {

	var prod1 Product = factory(SMALL, 100.0) // prod1 en realidad es una instancia de SmallProduct
	fmt.Println(prod1.Price())

	var prod2 Product = factory(MEDIUM, 100.0)
	fmt.Println(prod2.Price())

	var prod3 Product = factory(LARGE, 100.0)
	fmt.Println(prod3.Price())

}
