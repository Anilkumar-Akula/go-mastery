package main

import "fmt"

func main() {
	// age := 16
	// if age >= 18 {
	// 	fmt.Println("person is adult")
	// } else {
	// 	fmt.Println("Person is not adult")
	// }

	age := 10
	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is a kid")
	}

	var role = "admin"
	var hasPermissions = false
	if role == "admin" && hasPermissions {
		fmt.Println("Yes")
	}
	// age variable inside if construct
	if age := 20; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

}
