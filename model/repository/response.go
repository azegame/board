package repository

import (
	 _ "github.com/mattn/go-sqlite3"

	 "board/model/entity"
)


func InsertResponseInfo(threadId int, user string, massage string, resNo int) error {
	_, err := db.Exec(
		`INSERT INTO responses (
			thread_id,
			user_name,
			response_massage,
			response_no
		 ) VALUES (?,?,?,?)`, 
		threadId,
		user,
		massage,
		resNo,
	)
	return err
}


func GetResponse(threadId int) ([]entity.Response, error) {
	var responses []entity.Response

	rows, err := db.Query(
		`SELECT
			user_name,
			response_massage,
			response_no,
			created_at
		 FROM 
			responses
		 WHERE 
			thread_id = ?`,
		threadId,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r := entity.Response{}
		err = rows.Scan(
			&r.UserName,
			&r.ResponseMessage,
			&r.ResNo,
			&r.CreatedAt,
		)
		if err != nil {
			break
		}
		responses = append(responses, r)
	}
	return responses, err
}


//投稿するレスのNoの取得
//(スレッドIDに一致したレス数をカウント(スレッドのレス数の検索)、
//その数+1したものを連番として使用)
func GetResponseNo(threadId int) (int, error) {
	var resNo int

	row := db.QueryRow(
		`SELECT
			COUNT(response_id) + 1
		 FROM
		 	responses
		 WHERE
		 	thread_id = ?`,
		threadId,
	)
	err := row.Scan(&resNo)
	return resNo, err
}


func CountResponse(threadId int) (int, error) {
	var count int

	row := db.QueryRow(
		`SELECT
			COUNT(response_id)
		 FROM
			responses
		 WHERE
			thread_id = ?`,
		threadId,
	)
	err := row.Scan(&count)
	return count, err
}



