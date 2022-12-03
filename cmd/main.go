package main

import (
	"fmt"
	"github.com/rustingoff/hex-arch-go/internal/adapters/core/arithmetic"
	"log"
)

func main() {
	//

	arithAdaper := arithmetic.NewArithmeticAdapter()
	result, err := arithAdaper.Addition(1, 1)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)
}
