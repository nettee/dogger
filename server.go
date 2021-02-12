package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func beforeHandle(w http.ResponseWriter, r *http.Request) {
	// Print basic info
	fmt.Printf("> %v %v\n", r.Method, r.URL)
	vars := mux.Vars(r)
	for i, e := range vars {
		fmt.Printf("%v='%v'\n", i, e)
	}
	// Print body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error read body: %v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	fmt.Printf("body=%s\n", body)

}

func handlePing(w http.ResponseWriter, r *http.Request) {
	beforeHandle(w, r)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	//w.WriteHeader(http.StatusInternalServerError)
	//_, _ = w.Write([]byte("not implemented by dogger"))
}

type InspectContainerResult struct {
}

func handleInspectContainer(w http.ResponseWriter, r *http.Request) {
	beforeHandle(w, r)
	w.WriteHeader(http.StatusOK)
	result := InspectContainerResult{}
	_ = json.NewEncoder(w).Encode(result)
}

type CreateExecResult struct {
	Id string
}

func handleCreateExec(w http.ResponseWriter, r *http.Request) {
	beforeHandle(w, r)
	w.WriteHeader(http.StatusOK)
	result := CreateExecResult{
		Id: "aaaaaaa",
	}
	_ = json.NewEncoder(w).Encode(result)
}

func handleStartExec(w http.ResponseWriter, r *http.Request) {
	beforeHandle(w, r)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("{}"))
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	beforeHandle(w, r)

	fmt.Println("not implemented")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte("not implemented by dogger"))
}

type Route struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var routes = []Route{
	{"/_ping", handlePing},
	{"/v1.24/containers/{id}/json", handleInspectContainer},
	{"/v1.24/containers/{id}/exec", handleCreateExec},
	{"/v1.24/exec/{id}/start", handleStartExec},
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
	router.NotFoundHandler = http.HandlerFunc(handleNotFound)
	http.Handle("/", router)

	_ = http.Serve(listener, nil)
}

func main() {
	sockFile := "/tmp/dogger.sock"
	serve(sockFile)
}
