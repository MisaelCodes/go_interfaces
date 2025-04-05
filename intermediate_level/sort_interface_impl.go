package intermediate_level

import "math"

type Product struct {
	Name  string
	Price float64
}

type ProductList []Product

func (pl ProductList) Len() int {
	return len(pl)
}

func (pl ProductList) Less(i, j int) bool {
	return pl[i].Price < pl[j].Price || (math.IsNaN(pl[i].Price) && !math.IsNaN(pl[j].Price))
}

func (pl ProductList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}
