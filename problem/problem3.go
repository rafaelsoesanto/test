package problem

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Problem3 discounts product price by 50% (i.e. price/2)

	URL: http://localhost:9000/problem3

	Expectation: price becomes cheaper!
	Reality: but it failed!
*/

// Problem3 discounts product price by 50% (i.e. price/2)
func Problem3(w http.ResponseWriter, r *http.Request) {
	// ...might as well copy-paste from problem two
	products := []product{
		{"Pen", 3000, 4},
		{"Pineapple", 2000, 3},
		{"Apple", 5000, 2},
		{"Pen", 6000, 1},
	}

	// time to multiply the price
	for i := 0; i < len(products); i++ {
		products[i].Price = products[i].Price / 2
		fmt.Printf("%s's price becomes %d\n", products[i].Name, products[i].Price) // log the calculation just to make sure!
	}

	// marshal the json
	result, err := json.Marshal(products)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(result)
}
