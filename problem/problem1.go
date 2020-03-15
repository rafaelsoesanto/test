package problem

import (
	"fmt"
	"net/http"
	"time"
)

type person struct {
	name string
	age  int
}

/*
	Problem1 benchmarks performance between goroutine and no goroutine

	URL: http://localhost:9000/problem1

	Expectation: data in server should be consistent, with or without goroutine
	Reality: ...why the data becomes unpredictable after using goroutine?
*/

// Problem1 benchmarks performance between goroutine and no goroutine
func Problem1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Benchmark Start!")

	group := []person{
		{"Alpha", 10},
		{"Beta", 12},
		{"Gamma", 14},
		{"Delta", 16},
		{"Epsilon", 18},
		{"Zeta", 20},
		{"Theta", 22},
		{"Eta", 24},
		{"Iota", 26},
		{"Kappa", 28},
	}

	start1 := time.Now()
	for _, p := range group {
		p.PrintSelf()
	}
	duration1 := fmt.Sprintf("%d", time.Since(start1).Nanoseconds())

	fmt.Println("==========================")

	start2 := time.Now()
	for _, p := range group {
		go p.PrintSelf() // please do not change this line, there should be some other way!
	}
	duration2 := fmt.Sprintf("%d", time.Since(start2).Nanoseconds())

	result := "No goroutine: " + duration1 + "ns\nGoroutine: " + duration2 + "ns"

	w.Write([]byte(result))
}

// PrintSelf prints person's detail in a proper sentence
func (p *person) PrintSelf() {
	fmt.Printf("%s is %d years old.\n", p.name, p.age)
}
