package main

import "fmt"

var grade int

func main() {

	fmt.Println("grade calculater")
	fmt.Scan("%d", &grade)
	fmt.Println("score = ", grade)
	if grade >= 80 {
		fmt.Println("A")
	} else if grade < 80 && grade >= 70 {
		fmt.Println("B")
	} else if grade < 70 && grade >= 60 {
		fmt.Println("C")
	} else if grade < 60 && grade >= 50 {
		fmt.Println("D")
	} else if grade < 50 {
		fmt.Println("F")
	}

	// myMoney := 100
	// if myMoney > 100 {
	// 	fmt.Println("คุณสามารถซื้อได้")
	// } else {
	// 	fmt.Println("เงินไม่พอ")
	// }

}
