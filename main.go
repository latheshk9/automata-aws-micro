package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

var HelloMessage string

func helloWorld(w http.ResponseWriter, r *http.Request) {
	HelloMessage := "Hello World New"
    out, _ := exec.Command("bash", "-c", "hostname").Output()
    HelloMessage = HelloMessage + ": " + string(out)
    fmt.Fprintf(w, HelloMessage)
}

func main() {
	//Configuration()
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

type Config struct {
	Message string `json:"message"`
}

func Configuration() {
	file, err := os.Open("properties.json")
	if err != nil {
		fmt.Println("file error:", err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	configuration := Config{}

	erro := decoder.Decode(&configuration)
	if erro != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	HelloMessage = configuration.Message
}
