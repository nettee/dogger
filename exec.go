package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

type Exec struct {
	Cmd []string
}

var execInstances = make(map[string]Exec)

type CreateExecRequest struct {
	User         string
	Privileged   bool
	Tty          bool
	AttachStdin  bool
	AttachStdout bool
	AttachStderr bool
	Detach       bool
	DetachKeys   string
	Env          string
	WorkingDir   string
	Cmd          []string
}

type CreateExecResult struct {
	Id string
}

func HandleCreateExec(w http.ResponseWriter, r *http.Request) {
	BeforeHandle(w, r)

	var request CreateExecRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("request: %v\n", request)

	exec := Exec{Cmd: request.Cmd}
	id := fmt.Sprintf("%x", uuid.New().ID())
	execInstances[id] = exec

	w.WriteHeader(http.StatusOK)
	result := CreateExecResult{
		Id: id,
	}
	_ = json.NewEncoder(w).Encode(result)
}

func execute(exe Exec) {
	cmd := exe.Cmd
	fmt.Printf("cmd = %v\n", cmd)
	command := exec.Command(cmd[0], cmd[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		log.Fatalf("command failed: %s\n", err.Error())
	}
	fmt.Println(string(output))
}

func HandleStartExec(w http.ResponseWriter, r *http.Request) {
	BeforeHandle(w, r)

	vars := mux.Vars(r)
	id := vars["id"]
	exec1 := execInstances[id]
	execute(exec1)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("{}"))
}
