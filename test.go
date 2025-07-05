// package main

// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	var value string
// 	fmt.Print("This is a print function")
// 	fmt.Scan(&value)
// 	fmt.Println("%s is the value you are talking about", value)

// 	//we will test the lookup function
// 	names, err := net.LookupHost("mc.004545.xyz")
// 	if err != nil {
// 		fmt.Printf("there is an errror", err)
// 	}
// 	// let's iterate over it using for loop

// 	for index, value := range names {
// 		fmt.Println(index, value)
// 	}
// }
