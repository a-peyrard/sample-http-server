package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

var appVersion = "0.0.1"

func getHealth(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"version": appVersion,
		"http":    "ok",
	}
	bytes, _ := json.Marshal(data)
	_, _ = w.Write(bytes)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", getHealth)

	err := http.ListenAndServe(":8333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
