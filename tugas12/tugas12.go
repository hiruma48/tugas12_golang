package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "aku SUka KAmu saYAN9"
	var regexp, err = regexp.Compile(`[A-Z]+\S`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var hasil = regexp.FindAllString(text, 3)
	fmt.Println(hasil)

	var index1 = regexp.FindStringIndex(text)
	fmt.Println(index1)
}
