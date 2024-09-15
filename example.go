package main

import ("fmt")


func main () {
	var x int
	fmt.Println("Enter the avg: ")
	fmt.Scanln(&x)
	if x >= 50 {
		fmt.Println("success")
	} else {
		fmt.Println("faild")
	}
	


}