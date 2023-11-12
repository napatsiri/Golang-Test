package main

import "fmt"

type employee struct {
	employeeId   string
	employeeName string
	phone        string
}

func main() {
	employee1 := employee{
		employeeId:   "101",
		employeeName: "Natty",
		phone:        "098908005",
	}

	fmt.Println("employee1 = ", employee1)

	employeeList := [3]employee{}
	employeeList[0] = employee{
		employeeId:   "102",
		employeeName: "Juney",
		phone:        "094490198095",
	}
	employeeList[1] = employee{
		employeeId:   "103",
		employeeName: "froey",
		phone:        "0942400095",
	}
	employeeList[2] = employee{
		employeeId:   "104",
		employeeName: "Juree",
		phone:        "094443243325",
	}

	fmt.Println("employee = ", employeeList)

	employeeSlide := []employee{}
	employee5 := employee{
		employeeId:   "105",
		employeeName: "Juop",
		phone:        "094440909025",
	}
	employee6 := employee{
		employeeId:   "106",
		employeeName: "hdgyy",
		phone:        "90998978025",
	}

	employeeSlide = append(employeeSlide, employee5, employee6)
	fmt.Println("employee = ", employeeSlide)
}
