package main

import(
	"fmt"
	"net/http"
	"html/template"
)

func main(){
	server := http.Server{
        Addr: "127.0.0.1:8080",
	}
	//at / call handler of type Handler interface
	//HandlerFunc(f) returns a "Handler" that calls f
	// http.Handle("/" , http.HandlerFunc(handlerMethod))
	//or we can use helper method HandleFunc which takes a method instead of a Handler
    http.HandleFunc("/", handlerMethod)
    server.ListenAndServe()
}

//func of the Handler interface
func handlerMethod(res http.ResponseWriter , req *http.Request){
	t,_ := template.ParseFiles("index.html")
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	name := req.PostFormValue("name")
	fmt.Println(req.Form)
	// fmt.Println(req.PostForm) this gets only post values
	// Or we can write to the response like this 
	// res.Header().Set("Content-type"  , "text/plain")
	// fmt.Fprintln(res , "pass some html here")

	//Execute the template and write it to the response
	t.Execute(res , name)
}	