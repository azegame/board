package entity


//masterテーブル(データをupdateしない)
type Thread struct {
	ThreadId int `db:"thread_id"`
	Title string `db:"title"`
	CreatedAt string `db:"created_at"`
}



