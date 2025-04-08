package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/MisaelCodes/go_interfaces/basic_level"
	"github.com/MisaelCodes/go_interfaces/intermediate_level"
)

func main() {
	b := basic_level.Book{Title: "Go Programming", Author: "John Doe"}
	fmt.Println(b)

	res, err := basic_level.Divide(10.0, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %v\n", res)
	}
	res, err = basic_level.Divide(10.0, 0)
	if err != nil {
		fmt.Println(err)
	}
	hr := basic_level.HelloReader{}
	var res2 int
	bArr := make([]byte, 5)
	res2, err = io.ReadFull(hr, bArr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res2)
		fmt.Println(bArr)
		fmt.Println(string(bArr))
	}
	cl := basic_level.ConsoleLogger{}
	cl.Write(bArr)
	basic_level.MakeSound(basic_level.Dog{})

	//----------- Intermediate -----//

	product1 := intermediate_level.Product{Name: "Invoice", Price: 1.00}
	product2 := intermediate_level.Product{Name: "Invoice", Price: 2.00}
	product3 := intermediate_level.Product{Name: "Invoice", Price: 3.00}
	product4 := intermediate_level.Product{Name: "Invoice", Price: 4.00}
	product5 := intermediate_level.Product{Name: "Invoice", Price: 5.00}

	pl := intermediate_level.ProductList{product5, product2, product4, product3, product1}
	fmt.Println(pl)
	sort.Sort(pl)
	fmt.Println(pl)

	io.Copy(basic_level.ConsoleLogger{}, intermediate_level.UpperCaseReader{S: "hola go"})
	cl2 := intermediate_level.FileLogger{}
	intermediate_level.LoggingStuff(cl2, "hellow go")

    intermediate_level.DescribeType(11)
    intermediate_level.DescribeType("hello")
    intermediate_level.DescribeType(product1)

    db1 := intermediate_level.MockDB{}
    db2 := intermediate_level.PostgreDB{}
    intermediate_level.FetchData(&db1)
    intermediate_level.FetchData(&db2)

}
