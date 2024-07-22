package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var userName string
	fmt.Println(userName)
	fmt.Printf("the type is: %T \n", userName)

	var numbers [3]int = [3]int{1, 2, 3}         // Array
	var scores []int = []int{95, 85, 75}         // Slice
	var john = Person{Name: "Mayank", Age: 6700} // Struct

	fmt.Println(numbers, scores, john)

	var channels chan int = make(chan int) // Channel

	var studentGrades map[string]int = map[string]int{"John": 90, "Jane": 85} // Map

	fmt.Println(channels, studentGrades)

	pointers()
	maps()
}

func pointers() {
	var number int = 42
	var ptr *int = &number // Declare a pointer to an integer and assign it the address of 'number'

	// Print the pointer and the value it points to
	fmt.Println("Pointer address:", ptr) // Memory address of 'number'
	fmt.Println("Pointer value:", *ptr)  // Value stored at the address
}

func maps() {
	// Declare and initialize a map with string keys and integer values
	var studentGrades map[string]int = map[string]int{"John": 90, "Jane": 85}
	var isPresent map[string]bool = map[string]bool{"Mayank": true, "Tiwari": false}

	fmt.Println("Student Grades:", studentGrades)

	// Add a new key-value pair to the map
	studentGrades["Alice"] = 95
	isPresent["gojo"] = true
	// Retrieve and print a value from the map
	fmt.Println("Jane's grade:", studentGrades["Jane"])

	// Delete a key-value pair from the map
	delete(studentGrades, "John")

	fmt.Println("Updated Student Grades:", studentGrades)

	fmt.Print(isPresent)
	fmt.Print("the status is : ", isPresent["gojo"])
}
