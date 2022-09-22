package repository

import (
	 _ "github.com/mattn/go-sqlite3"

	 "board/model/entity"
)


func InsertThreadInfo(title string) error {
	_, err := db.Exec(`INSERT INTO threads (title) VALUES(?)`, title)
	return err
}


func GetThreadByTitle(title string) (entity.Thread, error) {
	var ret entity.Thread
	err := db.QueryRow(
		`SELECT 
			thread_id, 
			title
		 FROM 
		 	threads
		 WHERE 
		 	title = ?`,
		 title,
	).Scan(
		&ret.ThreadId, 
		&ret.Title,
	)
	return ret, err
}


func GetThreadByThreadId(threadId int) (entity.Thread, error) {
	var ret entity.Thread
	err := db.QueryRow(
		`SELECT
			thread_id, 
			title
		 FROM
		 	threads
		 WHERE
		 	thread_id = ?`,
		threadId,
	).Scan(
		&ret.ThreadId,
		&ret.Title,
	)
	return ret, err
}


func GetAllThreads() ([]entity.Thread, error) {
	var threads []entity.Thread
	rows, err := db.Query(
		`SELECT
		 	thread_id,
		 	title
		 FROM
		 	threads
		 ORDER BY
		 	thread_id`,
	)

	for rows.Next() {
		r := entity.Thread{}
		err = rows.Scan(
			&r.ThreadId,
			&r.Title,
		)
		if err != nil {
			break
		}
		threads = append(threads, r)
	}
	return threads, err
}



