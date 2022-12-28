package main

import (
	"encoding/json"
	"fmt"
	database "fun-in-golang/internal"
	"log"
	"net/http"
	"time"
)

type apiConfig struct {
    dbClient database.Client
}

func main()  {
    serveMux := http.NewServeMux()

    dbClient := database.NewClient("db.json")
    err := dbClient.EnsureDB()
    if err != nil {
        log.Fatal(err)
    }

    apiCfg := apiConfig{
        dbClient: dbClient,
    }

    serveMux.HandleFunc("/users/", apiCfg.endpointUsersHandler)
    serveMux.HandleFunc("/users", apiCfg.endpointUsersHandler)

    const addr = "localhost:8000"
    server := http.Server{
        Handler: serveMux,
        Addr: addr,
        WriteTimeout: 30 * time.Second,
        ReadTimeout: 30 * time.Second,
    }
    fmt.Println("server started on", addr)
    err = server.ListenAndServe()
    log.Fatal(err)
}

type errorBody struct {
    Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, err error) {
    if err == nil {
        log.Println("don't call respondWithError with a nil err!")
        return
    }
    log.Println(err)
    respondWithJSON(w, code, errorBody{
        Error: err.Error(),
    })
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    if payload != nil {
        response, err := json.Marshal(payload)
        if err != nil {
            log.Println("error marshalling", err)
            w.WriteHeader(500)
            response, _ := json.Marshal(errorBody{
                Error: "error marshalling",
            })
            w.Write(response)
            return

        }
        w.WriteHeader(code)
        w.Write(response)
    }
}
