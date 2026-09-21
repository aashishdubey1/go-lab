package interfaces

import (
	"fmt"
	"sort"
)

// -> -> -> sort.Interface From Scratch
// Given type Product struct{ Name string; Price float64 },
// implement sort.Interface twice — once sorting by Price ascending,
// once by Name alphabetically — as two named slice types, and sort the same underlying data both ways.

type Product struct {
	Name  string
	Price float64
}

type ByPrice []Product

type ByName []Product

func (bp ByPrice) Len() int {
	return len(bp)
}

func (bp ByPrice) Swap(i, j int) {
	bp[i], bp[j] = bp[j], bp[i]
}

func (bp ByPrice) Less(i, j int) bool {
	return bp[i].Price < bp[j].Price
}

func (bn ByName) Len() int {
	return len(bn)
}

func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

func (bn ByName) Less(i, j int) bool {
	return bn[i].Name < bn[j].Name
}

func RunEx2() {
	data := []Product{
		{Name: "Laptop", Price: 999.99},
		{Name: "Apple", Price: 1.50},
		{Name: "Mouse", Price: 25.00},
	}

	sort.Sort(ByPrice(data))
	fmt.Println("Sorted by Price:", data)

	sort.Sort(ByName(data))
	fmt.Println("Sorted by Name:", data)

}
