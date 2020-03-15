package problem

import (
	"fmt"
	"net/http"
)

/*
	Problem4 plays around with errors

	URL: http://localhost:9000/problem4

	Expectation: browser shows the error which stops the main flow
	Reality: everything is fine (?)
*/

// Problem4 plays around with errors
func Problem4(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		message := ""
		if err != nil {
			message = err.Error()
		} else {
			message = "I'm happy, you're happy, we're all happy~"
		}
		w.Write([]byte(message))
	}()

	// no error here
	if err := NoError(); err != nil {
		fmt.Println("You should not see this error:", err.Error())
	}

	// this should have error
	if err = FatalError(); err != nil {
		fmt.Println("You must see this error:", err.Error())
		return
	}

	fmt.Println("No error, yippee!")
}

// NoError returns no error (nil)
func NoError() error {
	return nil
}

// FatalError returns error in which the main flow should stop if it happens
func FatalError() error {
	return fmt.Errorf("fatal error please stop")
}
