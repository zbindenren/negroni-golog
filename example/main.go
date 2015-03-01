package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/zbindenren/negroni-golog"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "success!\n")
	})

	n := negroni.New()
	// example with timehop's standard logger
	n.Use(negronigolog.NewLogger())
	// n.Use(negronigolog.NewLoggerWithPrefix("myapp")) // example with custom prefix
	// n.Use(negronigolog.NewLoggerWithPrefixAndFlags("myapp", log.FlagsDate|log.FlagsTime|log.FlagsPrecisionTime)) // example with custom prefix and time format
	n.UseHandler(r)

	n.Run(":3000")
}
