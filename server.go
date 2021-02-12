package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("> %v %v\n", r.Method, r.URL)
	fmt.Printf("%v\n", r.Body)

	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("not implemented by dogger"))
}

func serve(sockFile string) {
	_ = os.Remove(sockFile)

	listener, _ := net.Listen("unix", sockFile)
	// You can also use TCP socket
	//listener, _ := net.Listen("tcp", "localhost:9000")

	fmt.Printf("listen to %v...\n", sockFile)

	http.HandleFunc("/", handleRoot)
	_ = http.Serve(listener, nil)
}

func main() {
	sockFile := "/tmp/dogger.sock"
	serve(sockFile)
}
