package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
    "encoding/json"
    "io/ioutil"
	"net/http"
)

type PostParams struct {
    Image   string  `json:"image"`
    Caption string  `json:"caption"`
}

func (rt *_router) newPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    idc, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        http.Error(w, "Unauthenticated", http.StatusUnauthorized)
        return
    } else if err != nil {
    	http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    	return
    }
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    var postParams PostParams
    id := idc.Value
    err = json.Unmarshal(body, &postParams)
    if err != nil {
    	http.Error(w, "Bad request: malformed json: " + err.Error(), http.StatusBadRequest)
    	return
    }

    postID, err := rt.db.NewPost(id, postParams.Image, postParams.Caption)
    if err == database.ErrBadImage {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }
    if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("content-type", "text/plain")
    _, _ = w.Write([]byte(string(postID)))
}