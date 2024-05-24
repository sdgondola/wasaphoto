package api

import (
    "github.com/sdgondola/wasaphoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
    "strconv"
)

func (rt *_router) likeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    idc, err := r.Cookie("WASASESSIONID")
    if err == http.ErrNoCookie {
        http.Error(w, "Unauthenticated", http.StatusUnauthorized)
        return
    } else if err != nil {
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
        return
    }
    toBlock := ps.ByName("commentID")
    if toBlock == "" {
        http.Error(w, "Bad request: no username provided", http.StatusBadRequest)
        return
    }
    id := idc.Value
    commentID, err := strconv.ParseInt(ps.ByName("commentID"), 10, 64)
    if err != nil {
        http.Error(w, database.ErrPostNotFound.Error(), http.StatusNotFound)
        return
    }

    err = rt.db.LikeComment(id, commentID)
    if err == database.ErrUserNotFound {
        // This is kinda suspicious, likely a forged cookie
        http.Error(w, "Bad request: hacking attempt?!", http.StatusBadRequest)
    } else if err == database.ErrCommentNotFound {
        http.Error(w, database.ErrCommentNotFound.Error(), http.StatusNotFound)
    } else if err == database.ErrUserIsBlocked {
        http.Error(w, "Cannot like post: user blocked you!", http.StatusForbidden)
    } else if err != nil && err != database.ErrAlreadyLiked {    // We can safely ignore that as it's likely some duplicate request
        http.Error(w, "Internal server error: " + err.Error(), http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusCreated)
    }
}
