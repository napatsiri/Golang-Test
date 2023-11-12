package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product = ", product)
	product["macbook"] = 400000
	product["ipad"] = 40000
	product["iphone"] = 4000
	fmt.Println("product = ", product)

	delete(product, "ipad")
	fmt.Println("product = ", product)

	product["macbook"] = 300000
	fmt.Println("product = ", product)

	value1 := product["iphone"]
	fmt.Println("value1 =", value1)

	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println("coures = ", courseName)
}
