package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(
		tracer.WithService("test"),
		tracer.WithEnv("dev"),
	)
	defer tracer.Stop()

	// Create a traced mux router
	mux := httptrace.NewServeMux()
	// Continue using the router as you normally would.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(100) >= 50 {
			fmt.Println("uhoh slow inference :( let's debug why this sucks")
			time.Sleep(10 * time.Second)
		}
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8080", mux)
}
