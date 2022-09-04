package main

import (
	"flag"
	"fmt"

	"github.com/eljobe/nonmem/listmonk"
)

func main() {
	flag.Parse()
	fmt.Println("Making or updating a Non-Member's List.")
	fmt.Println("Talking to Listmonk on ", listmonk.ApiUrl)
}
