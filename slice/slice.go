package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"java", "pythin"}
	fmt.Println(courseName)
	courseName = append(courseName, "c", "C#", "html", "css", "Javascript")
	fmt.Println(courseName)

	courseWeb := courseName[4:7]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
