package main

import "os"

func main() {
	data1 := []byte("hello\n  Dev")
	err := os.WriteFile("data.txt", data1, 0644)

	if err != nil {
		panic(err)
	}

	f, err := os.Create("employeeName")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	data2 := []byte("Sira\n Manee")
	os.WriteFile("employeeName.txt", data2, 0644)
}
