package main

import (
	"fmt"
	"os"

	"github.com/genkiroid/sp"
)

func main() {
	var yml = sp.New(os.Stdin)
	fmt.Printf("%s", yml)
}
