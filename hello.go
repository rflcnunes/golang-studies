package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Rafael"
	fmt.Println("Hello,", name)

	version := 1.1
	fmt.Println("Version:", version)

	fmt.Println("Variable version type:", reflect.TypeOf(version))

	fmt.Println("1 - Init monitoring")
	fmt.Println("2 - See logs")
	fmt.Println("0 - Exit")

	var command int
	fmt.Scan(&command)

	fmt.Println("Your choice:", command)

	switch command {
	case 1:
		fmt.Println("Start monitoring")
	case 2:
		fmt.Println("Displaying logs")
	case 0:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Invalid command")
	}
}
