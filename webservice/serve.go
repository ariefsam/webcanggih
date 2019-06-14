package webservice

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type webCanggihHanlder struct{}

func (app webCanggihHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!")
}

func StartService() {

	var myHandler webCanggihHanlder

	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
