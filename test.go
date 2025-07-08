package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// var value string
	// fmt.Print("This is a print function")
	// fmt.Scan(&value)
	// fmt.Println("%s is the value you are talking about", value)

	// //we will test the lookup function
	// names, err := net.LookupHost("mc.004545.xyz")
	// if err != nil {
	// 	fmt.Printf("there is an errror", err)
	// }
	// // let's iterate over it using for loop

	// for index, value := range names {
	// 	fmt.Println(index, value)
	// }

	// let's use sha256
	msg := []byte("hello this is a test")
	val := sha256.Sum256(msg)
	res := hex.EncodeToString(val[:])
	fmt.Println(res)
}
