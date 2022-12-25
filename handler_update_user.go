package main

import (
	"encoding/json"
	"net/http"
	"strings"
)



func (apiCfg apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
    type parameters struct {
        Password string `json:"password"`
        Name     string `json:"name"`
        Age      int    `json:"age"`
    }

    decoder := json.NewDecoder(r.Body)
    params := parameters{}
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, err)
        return
    }
    path := r.URL.Path
    email := strings.TrimPrefix(path, "http://localhost:8000/users/")
    updUser, err := apiCfg.dbClient.UpdateUser(email, params.Password, params.Name, params.Age)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, err)
        return
    }
    respondWithJSON(w, http.StatusOK, updUser)
}
