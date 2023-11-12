package main

import "fmt"

var productName [4]string
var price [4]float32

func main() {
	productName[0] = "macbook"
	productName[1] = "ipad"
	productName[2] = "iphone"
	productName[3] = "airpod"

	price := [4]float32{400000, 30000, 320000, 10000}
	fmt.Println(productName)
	fmt.Println(price)
}
