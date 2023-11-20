package main

import (
	"fmt"

	"rsc.io/quote"
)

// So main should just init the server and start accepting HTTP requests.
func main() {
	fmt.Println(quote.Go())
}
