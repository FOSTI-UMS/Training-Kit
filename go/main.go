package main

import "fmt"

func main()  {
	fmt.Println("FOSTI UMS")
	hasil := linearSearch([]int{1, 5, 4, 3, 9, 3, 5}, 90)
	fmt.Println(hasil)

	//menambah function untuk encode dan decode base64
	encode()
	decode()
}