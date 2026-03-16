package main

import(
	"fmt"
	"net/http"
)

func main(){
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hello, DevOps! This is a Go binary running.")
        })

        fmt.Println("Server starting on :8020...")
        http.ListenAndServe(":8020", nil)
}