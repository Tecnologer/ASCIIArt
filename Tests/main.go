package main

import (
	"fmt"

	"github.com/tecnologer/Strings"
)

func main() {
	var word = "hola mundo!"

	fmt.Println(chars.Parse(word))
	fmt.Println(chars.Parsef(word, '♠'))
	fmt.Println(chars.ParsefBackground(word, '-', '♠'))
}
