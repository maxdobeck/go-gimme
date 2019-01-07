package main

import "fmt"


func main() {
	const spinner = `. .. ... .... .....`
	fmt.Println("Hello World")
	i := 0
	for {
		fmt.Println(i%5) 
		i++
		if i == 5 {
			break
		}
	}
}
