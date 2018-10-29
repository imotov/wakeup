package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("xdotool", "key", "Escape")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s\n", out.String())
}

func main() {
	http.HandleFunc("/wakeup", handler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
