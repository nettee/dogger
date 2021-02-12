package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

type HttpHandler struct {
}

func (handler HttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("> %v %v\n", request.Method, request.URL)
	fmt.Printf("%v\n", request.Body)
	_, _ = fmt.Fprintf(response, "hello, docker\n")
}

func serve(sockFile string) {
	_ = os.Remove(sockFile)

	listener, _ := net.Listen("unix", sockFile)
	// You can also use TCP socket
	//listener, _ := net.Listen("tcp", "localhost:9000")

	fmt.Printf("listen to %v...\n", sockFile)

	http.Handle("/", HttpHandler{})
	_ = http.Serve(listener, nil)
}

func main() {
	sockFile := "/tmp/dogger.sock"
	serve(sockFile)
}
