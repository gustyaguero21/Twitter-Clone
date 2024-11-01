package config

//Router params

const (
	Port = ":8080"
)

//Database params

const (
	Driver = "sqlite"
	DbPath = "./internal/data/database.db"
)

//Database queries

const (
	CreateUserTable = `CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY NOT NULL,
		username TEXT NOT NULL);`

	CreateFollowsTable = `CREATE TABLE IF NOT EXISTS follows (
		id TEXT PRIMARY KEY NOT NULL,
		follower_id INTEGER NOT NULL, 
		followed_id INTEGER NOT NULL, 
		FOREIGN KEY (follower_id) REFERENCES users(id), 
		FOREIGN KEY (followed_id) REFERENCES users(id) );`

	CreateTweetsTable = `CREATE TABLE IF NOT EXISTS tweets (
		id TEXT PRIMARY KEY NOT NULL,
		user_id TEXT NOT NULL,
		content TEXT NOT NULL,
		posted_at INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id));`
)
