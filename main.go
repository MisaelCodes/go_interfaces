package main

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/MisaelCodes/go_interfaces/advanced_level"
	"github.com/MisaelCodes/go_interfaces/basic_level"
	"github.com/MisaelCodes/go_interfaces/intermediate_level"
	"github.com/MisaelCodes/go_interfaces/professional_level"
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
	// ------------- Advanced --------//
	afl := advanced_level.NewFileLogger("/home/misa/Documents/learning/golang/specifics/go_interfaces/advanced.log")
	acl := advanced_level.ConsoleLogger{}
	amw := io.MultiWriter(afl, &acl)
	asr := strings.NewReader("something to be read")
	if _, err := io.Copy(amw, asr); err != nil {
		fmt.Println(err)
	}

	emailListener := &advanced_level.EmailListener{}
	advanced_level.ListenSomething(emailListener)
	inMCache := advanced_level.NewInMemoryCache(map[string]string{"1": "Hello Misa"})
	user := advanced_level.UserInfo{C: inMCache}
	fmt.Println(user.Get("1"))
	user.Get("some uuid")

	// At this point the io.Readers arent implemented as well as they should
	// the idea would be to read an insert in the given slice
	// there should be a pointer that says up to what index was read
	// so that the caller knows when a problem happened, like at what byte everything failed.
	// Another thing to do is always send the EOF when the data is fully readed.
	// the implemented reader has access to that data, it might be through some of its properties.

	// the caller can process the number of bytes read and if there's an error when processing return that.

	// for io.Writer  just take care of returning error when an error occurs in the middle of the writing
	// for example if we encounter a weird token or something, I'll still have to see more writers before
	// actually mapping all its use cases.

	circle := advanced_level.Circle{Radious: 2}
	rectangle := advanced_level.Rectangle{Width: 2, Height: 2}
	advanced_level.PrintArea(&circle)
	advanced_level.PrintArea(&rectangle)

	// Server so that we can implement an http.Handler
	// serverTest := &http.Server{
	//   Addr: ":8080",
	//   Handler: &advanced_level.HelloHandler{Data: "Hello http"},
	// }
	// log.Fatal(serverTest.ListenAndServe())
    plugingRegistry := []professional_level.Plugin{&professional_level.ReversePluging{}, &professional_level.UpperCasePlugin{}}
    for _, pluging := range plugingRegistry{
        resP := pluging.Execute("hola")
        fmt.Println(resP)
    }
}
