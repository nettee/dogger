package main

import (
	"encoding/json"
	"net/http"
)

type InspectContainerResult struct {
}

func HandleInspectContainer(w http.ResponseWriter, r *http.Request) {
	BeforeHandle(w, r)
	w.WriteHeader(http.StatusOK)
	result := InspectContainerResult{}
	_ = json.NewEncoder(w).Encode(result)
}
