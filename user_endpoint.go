package main

import (
	"errors"
	"net/http"
)



func (apiCfg apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        apiCfg.dbClient.GetUser("test1@example.com")
    case http.MethodPost:
        // handler
    case http.MethodPut:
        // handler
    case http.MethodDelete:
        // handler
    default:
        respondWithJSON(w, 404, errors.New("method not supported"))
    }
}
