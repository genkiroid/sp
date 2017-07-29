package sp_test

import (
	"fmt"
	"strings"

	"github.com/genkiroid/sp"
)

func ExampleYamlDataSet() {
	r := strings.NewReader(`
id	name	age	created_at
1	Alice	30	2017-07-29 00:00:00
2	Bob	20	20170729000000
3	Carol	22.0	20170729
4	Dan	30	2017-07-29
`)
	var yml = sp.YamlDataSet{r}
	fmt.Printf("%s", yml)
	// Output:
	// -
	//   id: 1
	//   name: "Alice"
	//   age: 30
	//   created_at: "2017-07-29 00:00:00"
	// -
	//   id: 2
	//   name: "Bob"
	//   age: 20
	//   created_at: "20170729000000"
	// -
	//   id: 3
	//   name: "Carol"
	//   age: 22.0
	//   created_at: "20170729"
	// -
	//   id: 4
	//   name: "Dan"
	//   age: 30
	//   created_at: "2017-07-29"
}
