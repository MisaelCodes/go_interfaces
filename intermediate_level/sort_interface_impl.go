package intermediate_level

import (
	"fmt"
	"math"
)

type Product struct {
	Name  string
	Price float64
}

func (p Product) String() string{
    return fmt.Sprintf("%s -- %v", p.Name, p.Price)
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

func (pl ProductList) String() string {
	s := ""
	for _, v := range pl {
		s += fmt.Sprintf("[%v] ", v)
	}
    return s
}
