package main

import (
	"fmt"

	"github.com/ducminhgd/zalo-go-sdk/x/pkce"
)

func main() {
	result := pkce.GetCodeChallenge("hello")
	fmt.Println(result)
}
