package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Println("Server is going to start...")
	err := http.ListenAndServe(":9090", nil)
	fmt.Println("Server is going to quit...")
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Test S2I process!\n")
}
