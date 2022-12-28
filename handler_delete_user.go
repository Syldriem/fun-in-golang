package main

import (
	"errors"
	"net/http"
	"strings"
)



func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    email := strings.TrimPrefix(path, "/users/")
    if email == "" {
        respondWithError(w, http.StatusBadRequest, errors.New("no email provided in path"))
        return
    }
    err := apiCfg.dbClient.DeleteUser(email)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, err)
        return
    }
    respondWithJSON(w, http.StatusOK, struct{}{})
}
