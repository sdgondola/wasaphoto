package database

import (
	"strings"

	"github.com/sdgondola/wasaphoto/service/globaltime"
)

func (db *appdbimpl) CommentPost(user string, postID int64, comment string) (int64, error) {
	user = strings.ToLower(user)
	exists, err := db.UserExists(user)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrUserNotFound
	}
	exists, err = db.PostExists(postID)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, ErrPostNotFound
	}
	var op string
	err = db.c.QueryRow("select author from Posts where postID = ?", postID).Scan(&op)
	if err != nil {
		return 0, err
	}
	blocked, err := db.IsBlockedBy(user, op)
	if err != nil {
		return 0, err
	}
	if blocked {
		return 0, ErrUserIsBlocked
	}
	ins, err := db.c.Prepare("insert into Comments values (?, ?, ?, ?) returning commentID")
	if err != nil {
		return 0, err
	}
	res, err := ins.Exec(globaltime.Now(), user, postID, comment)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
