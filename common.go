package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func BeforeHandle(w http.ResponseWriter, r *http.Request) {
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
