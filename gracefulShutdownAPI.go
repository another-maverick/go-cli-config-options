package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"github.com/braintree/manners"
)

func main() {
	handler := newHandler()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go shutmeDown(ch)

	manners.ListenAndServe(":12345", handler)


}

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	query := req.URL.Query()
	name := query.Get("name")

	fmt.Fprintf(res, "My name is %s", name)
}

func shutmeDown(ch <-chan os.Signal){
	<-ch
	manners.Close()
}


// Test by using ---- curl -k http://localhost:12345/?&name=Steph

