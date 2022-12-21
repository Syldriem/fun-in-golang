package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
    "./internal/database.go"
)

func main()  {
   fmt.Println("Working")

   serveMux := http.NewServeMux()

   serveMux.HandleFunc("/" , func (w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "Application/json")
        w.WriteHeader(200)
        w.Write([]byte("{standard:message}"))
   })
   serveMux.HandleFunc("/hi", handler)
   const addr = "localhost:8000"
   server := http.Server{
       Handler: serveMux,
       Addr: addr,
       WriteTimeout: 30 * time.Second,
       ReadTimeout: 30 * time.Second,
   }
   fmt.Println("server started on", addr)
   err := server.ListenAndServe()
   log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "Application/json")
    w.WriteHeader(200)
    w.Write([]byte("{hi:hello}"))
}

