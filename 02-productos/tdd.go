package main

type ProductType = int

const (
	SMALL  ProductType = iota // 0
	MEDIUM                    // 1
	LARGE                     // 2
)

type Product interface {
	Price() float64
}

// -------------------------------------
type SmallProduct struct {
	price float64
}

func (p SmallProduct) Price() float64 {
	return p.price
}

// -------------------------------------
type MediumProduct struct {
	price float64
}

func (p MediumProduct) Price() float64 {
	return p.price + (p.price * 0.03)
}

// -------------------------------------
type LargeProduct struct {
	price float64
}

func (p LargeProduct) Price() float64 {
	return p.price + (p.price * 0.06) + 2500.0
}

// -------------------------------------
