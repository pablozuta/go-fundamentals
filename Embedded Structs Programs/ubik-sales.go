// In this program, there are five struct types: person, employee, ubikProduct, ubikDiscount and ubikSale. The employee struct embeds the person struct using the syntax person. The ubikSale struct embeds both the ubikProduct and ubikDiscount structs using the syntax ubikProduct and ubikDiscount.

// In the main function, a new employee instance is created with a name of "Joe Chip", an age of 32, a department of "Runciter Associates", and a salary of 50000. A new ubikSale instance is also created with a product name of "Ubik Spray", a price of 20, a weight of 50, and a discount percentage of 10. The program then prints out Joe's information and the ubikSale information using fmt.Printf.
package main

import "fmt"

type personU struct {
	name string
	age  int
}

type employeeU struct {
	personU
	department string
	salary     int
}

type ubikProduct struct {
	name   string
	price  int
	weight int
}

type ubikDiscount struct {
	percentage int
}

type ubikSale struct {
	ubikProduct
	ubikDiscount
}

func main() {
	// create a new employee instance
	joe := employeeU{
		personU: personU{
			name: "Joe Chip",
			age:  32,
		},
		department: "Runciter Asociation",
		salary:     50_000,
	}

	// create a new ubikSale instance
	ubikSale := ubikSale{
		ubikProduct: ubikProduct{
			name:   "Ubik Spray",
			price:  20,
			weight: 50,
		},
		ubikDiscount: ubikDiscount{
			percentage: 10,
		},
	}

	// print out Joe's information
	fmt.Printf("Employee Name: %s\n", joe.name)
	fmt.Printf("Employee Age: %d\n", joe.age)
	fmt.Printf("Employee Department: %s\n", joe.department)
	fmt.Printf("Salary: $%d\n", joe.salary)

	fmt.Println()
	// print out the ubikSale information
	fmt.Printf("Product Name: %s\n", ubikSale.name)
	fmt.Printf("Product Price: $%d\n", ubikSale.price)
	fmt.Printf("Product Weight: %d\n", ubikSale.weight)
	fmt.Printf("Product Discount Percentage: %d\n", ubikSale.ubikDiscount.percentage)

}
