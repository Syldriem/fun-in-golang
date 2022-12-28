package main

import (
	"errors"
	"net/http"
	"strings"
)



func (apiCfg apiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    uuid := strings.TrimPrefix(path, "/posts/")
    if uuid == "" {
        respondWithError(w, http.StatusBadRequest, errors.New("no uuid provided in path"))
        return
    }
    posts, err := apiCfg.dbClient.GetPosts(uuid)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, err)
        return
    }
    respondWithJSON(w, http.StatusOK, posts)

}
