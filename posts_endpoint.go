package main

import (
	"errors"
	"net/http"
)



func (apiCfg apiConfig) endpointPostsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        apiCfg.handlerGetPosts(w, r)
    case http.MethodPost:
        apiCfg.handlerCreatePost(w, r)
    case http.MethodDelete:
        apiCfg.handlerDeletePost(w, r)
    default:
        respondWithJSON(w, 404, errors.New("method not supported"))
    }
}
