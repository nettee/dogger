package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	BeforeHandle(w, r)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	//w.WriteHeader(http.StatusInternalServerError)
	//_, _ = w.Write([]byte("not implemented by dogger"))
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	BeforeHandle(w, r)

	fmt.Println("not implemented")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("not implemented by dogger"))
}

type Route struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var routes = []Route{
	{"/_ping", HandlePing},
	{"/v1.24/containers/{id}/json", HandleInspectContainer},
	{"/v1.24/containers/{id}/exec", HandleCreateExec},
	{"/v1.24/exec/{id}/start", HandleStartExec},
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
	router.NotFoundHandler = http.HandlerFunc(HandleNotFound)
	http.Handle("/", router)

	_ = http.Serve(listener, nil)
}

func main() {
	sockFile := "/tmp/dogger.sock"
	serve(sockFile)
}
