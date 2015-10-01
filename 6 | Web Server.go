package main 

import ( "fmt"
	"net/http"	
)

//Creating the handler functions

//w will assemble the server's response, and write it to the client; r will hold the client's request
//Fprintf(w, "some_string") will write to the client screen
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}
//Differs from the previous handler function only in response
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Earth\n")
}

func main() {
	//When someone accesses the default "/" folder or the nested "/earth" folder on your server, the corresponding handler function is called.
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)
	//Starts a HTTP server and specifies which port to listen to.
	http.ListenAndServe(":8080", nil)
}

