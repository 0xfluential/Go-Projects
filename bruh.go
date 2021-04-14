package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter any number: ")
	fmt.Scanln(&n)
	i := 1
	/* for loop as a go's while */
	for {
		if(i>10){
			break;
		}
		fmt.Println(n," X ",i," = ",n*i)
		i++
	}
}