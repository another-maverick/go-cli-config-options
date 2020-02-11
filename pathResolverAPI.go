package main

import (
	"fmt"
	"net/http"
	"path"
)

func main(){

	pr := myNewPathResolver()

	// TEST: http://localhost:12345/
	pr.AddURI("GET /", basicHello)

	//  TEST: http://localhost:12345/hello
	pr.AddURI("GET /hello", customizedHello)

	// TEST: http://localhost:12345/random/
	pr.AddURI("* /random/", random)

	// TEST: http://localhost:12345/randomregex/asdasdasdasdadsdasd
	pr.AddURI("* /randomregex/*", randomRegex)

	// pr can be used since it is a struct that implements ServeHTTP method. and therefore it is a handler
	http.ListenAndServe(":12345", pr)


}

type pathResolver struct{
	myHandler map[string]http.HandlerFunc
}

func myNewPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

func (p *pathResolver) AddURI(path string, handler http.HandlerFunc){
	p.myHandler[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request){
	check := req.Method + " " + req.URL.Path
	for pathRegex, pathHandler := range p.myHandler{
		if ok, err := path.Match(pathRegex, check); ok && err  == nil {
			pathHandler(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}
	http.NotFound(res, req)
}


func basicHello(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res, "A Basic Hello")
}


func customizedHello(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res, "A Customized Hello")
}


func random(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "A random Function")
}

func randomRegex(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "A random Function with wild cards to the URL path")
}



