package main

import "fmt"

var name = "Aditya"

func main() {
	fmt.Println("Name from package scope: ", name)

	name := "Hajare" // 'name' from 'package scope' is shadowed
	fmt.Println("Name from main() block scope: ", name)

	{
		name := "Aditya Hajare" // 'name' from 'main() block scope' is shadowed
		fmt.Println("Name from inner block scope: ", name)
	}

	fmt.Println("Name from main() block scope: ", name) // 'name' from 'main() block scope' will be used
}
