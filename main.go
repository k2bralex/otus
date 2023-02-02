package main

import (

	/*
		"Otus/hw_1"
		"Otus/hw_2"
		"Otus/hw_3"

	*/

	"Otus/hw_3"
)

func main() {

	//HM_3

	str := "Firstly, similar to the sorting by key method, we have to obtain a slice of all the keys. " +
		"Now we want to sort the keys according to the values, for doing that, we use the SliceStable " +
		"function in the sort module. The slices table function takes in the slice and we can provide " +
		"less function. We can simply provide an anonymous lambda fiction that checks for the " +
		"comparison of the values of the provided slice."

	hw_3.TopStringWords(str, 10)

	/*

		//HM_2
		fmt.Println(hw_2.StringParse("a4bc2d5e"))
		fmt.Println(hw_2.StringParse("abcd"))
		fmt.Println(hw_2.StringParse("45"))
		fmt.Println(hw_2.StringParse("qwe\\4\\5"))
		fmt.Println(hw_2.StringParse("qwe\\45"))
		fmt.Println(hw_2.StringParse("qwe\\\\5"))

		//HM_1
		fmt.Println(hw_1.MyTime())

	*/
}
