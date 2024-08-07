package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetFeed(user int64) ([]int64, error) {
	feed := make([]int64, 0)
	exists, err := db.UserExists(user)
	if err != nil {
		return feed, err
	}
	if !exists {
		return feed, ErrUserNotFound
	}
	q, err := db.c.Query("select postID from Posts where author in (select following from Follows where follower = ?) order by pub_time desc", user)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return feed, err
	}
	for q.Next() {
		var p int64
		err = q.Err()
		if err != nil {
			return feed, err
		}
		err := q.Scan(&p)
		if err != nil {
			return feed, err
		}
		feed = append(feed, p)
	}
	return feed, nil
}
