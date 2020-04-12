//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//

package main

import "fmt"

func main() {
	var (
		phones map[string]string
		// Key        : Last name
		// Element    : Phone number

		// Key        : Product ID
		// Element    : Available / Unavailable
		products map[int]bool

		multiPhones map[string][]string
		// Key        : Last name
		// Element    : Phone numbers

		basket map[int]map[int]int
		// Key        : Customer ID
		// Element Key:
		//   Key: Product ID Element: Quantity
	)

	fmt.Printf("phones     : %#v\n", phones)
	fmt.Printf("products   : %#v\n", products)
	fmt.Printf("multiPhones: %#v\n", multiPhones)
	fmt.Printf("basket     : %#v\n", basket)
}
