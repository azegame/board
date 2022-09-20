package entity


type Response struct {
	ResponseId int `db:"response_id"`
	ThreadId int `db:"thread_id"`
	UserName string `db:"user_name"`
	ResponseMessage string `db:"response_massage"`
	ResNo int `db:"response_no"`
	CreatedAt string `db:"created_at"`
}



