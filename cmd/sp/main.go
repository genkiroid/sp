package main

import (
	"fmt"
	"os"

	"github.com/genkiroid/sp"
)

func main() {
	yml := sp.NewYamlDataSet(os.Stdin)
	fmt.Printf("%s", yml)
}
