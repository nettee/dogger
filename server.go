package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
)

func beforeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("> %v %v\n", r.Method, r.URL)
	vars := mux.Vars(r)
	for i, e := range vars {
		fmt.Printf("%v='%v'\n", i, e)
	}
	fmt.Printf("%v\n", r.Body)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	beforeHandle(w, r)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	//w.WriteHeader(http.StatusInternalServerError)
	//_, _ = w.Write([]byte("not implemented by dogger"))
}

func handleListContainers(w http.ResponseWriter, r *http.Request) {
	beforeHandle(w, r)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("not implemented by dogger"))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	beforeHandle(w, r)

	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("not implemented by dogger"))
}

type Route struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var routes = []Route{
	{"/_ping", handlePing},
	{"/v1.24/containers/{id}/json", handleListContainers},
}

func serve(sockFile string) {
	_ = os.Remove(sockFile)

	listener, _ := net.Listen("unix", sockFile)
	// You can also use TCP socket
	//listener, _ := net.Listen("tcp", "localhost:9000")

	fmt.Printf("listen to %v...\n", sockFile)

	router := mux.NewRouter()
	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler)
	}
	router.NotFoundHandler = http.HandlerFunc(handleRoot)
	http.Handle("/", router)

	_ = http.Serve(listener, nil)
}

func main() {
	sockFile := "/tmp/dogger.sock"
	serve(sockFile)
}
