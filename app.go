package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/google/gops/agent"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/tokopedia/grace.v1"
	"gopkg.in/tokopedia/logging.v1"

	"github.com/tokopedia/gosample/hello"
	"github.com/tokopedia/gosample/problem"
	"github.com/tokopedia/logging/tracer"
)

func main() {

	flag.Parse()
	logging.LogInit()

	debug := logging.Debug.Println

	debug("app started") // message will not appear unless run with -debug switch

	if err := agent.Listen(&agent.Options{}); err != nil {
		log.Fatal(err)
	}

	hwm := hello.NewHelloWorldModule()

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/hello", hwm.SayHelloWorld)
	http.HandleFunc("/zip", hwm.TestFunc)
	// go logging.StatsLog()

	http.HandleFunc("/problem1", problem.Problem1)
	http.HandleFunc("/problem2", problem.Problem2)
	http.HandleFunc("/problem3", problem.Problem3)
	http.HandleFunc("/problem4", problem.Problem4)
	http.HandleFunc("/problem5", problem.Problem5)

	tracer.Init(&tracer.Config{Port: 8700, Enabled: true})

	log.Fatal(grace.Serve(":9000", nil))
}
