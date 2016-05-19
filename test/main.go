package main

import (
	"fmt"
	"log"

	"github.com/PI-Victor/passgen"
)

func main() {
	newPass, err := passgen.GenPass(8, false, "")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(newPass)
}
