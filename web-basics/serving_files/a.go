package main

import(
	"io"
	"net/http"
)

func main(){
	server := http.Server{
        Addr: "127.0.0.1:8080",
	}
	//Setup file server
	http.HandleFunc("/", handlerMethod)
	http.Handle("/assets/" , http.StripPrefix("/assets/" , http.FileServer(http.Dir("./assets"))))
	// http.Handle("/assets" , http.HandlerFunc(handlerMethod))

	server.ListenAndServe()
}

//func of the Handler interface
func handlerMethod(res http.ResponseWriter , req *http.Request){
	res.Header().Set("Content-type" , "text/html")
	io.WriteString(res , `<h1> wtf </h1>`)
}	