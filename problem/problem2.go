package problem

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Problem2 marshals struct into json

	URL: http://localhost:9000/problem2

	Expectation: it should become proper json!
	Reality: WHERE IS MY DATA!?
*/

type product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Qty   int    `json:"qty"`
}

// Problem2 marshals struct into json
func Problem2(w http.ResponseWriter, r *http.Request) {
	products := []product{
		{"Pen", 3000, 4},
		{"Pineapple", 2000, 3},
		{"Apple", 5000, 2},
		{"Pen", 6000, 1},
	}

	result, err := json.Marshal(products)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	fmt.Println(products)
	fmt.Println(string(result))
	w.Write(result)
}
