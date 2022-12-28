package main

import (
	"errors"
	"net/http"
	"strings"
)



func (apiCfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    email := strings.TrimPrefix(path, "/users/")
    if email == "" {
        respondWithError(w, http.StatusBadRequest, errors.New("no email provided in path"))
        return
    }
    user, err := apiCfg.dbClient.GetUser(email)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, err)
        return
    }
    respondWithJSON(w, http.StatusOK, user)

}
