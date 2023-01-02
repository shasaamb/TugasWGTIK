package main

import "fmt"

type text [2022]rune

func main() {
	var kata text
	var s rune
	var i, len int
	var palindrom bool

	len = 0
	fmt.Scanf("%c", &s)
	for s != '.' {
		kata[len] = s
		len++
		fmt.Scanf("%c", &s)
	}

	palindrom = true
	for i = 0; i < len && palindrom; i++ {
		palindrom = kata[i] == kata[len-i-1]
	}
	fmt.Println(palindrom)

}
