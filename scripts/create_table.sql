CREATE TABLE IF NOT EXISTS threads (
	thread_id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT UNIQUE NOT NULL,
	created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);


CREATE TABLE IF NOT EXISTS responses (
	response_id INTEGER PRIMARY KEY AUTOINCREMENT,
	thread_id INTEGER NOT NULL,
	user_name TEXT NOT NULL,
	response_massage TEXT NOT NULL,
	response_no INTEGER DEFAULT 1,
	created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);



