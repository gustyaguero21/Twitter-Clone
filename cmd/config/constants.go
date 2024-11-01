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

	CreateFollowerTable = `CREATE TABLE IF NOT EXISTS followers (
		follower_username TEXT,
		following_username TEXT,
		FOREIGN KEY (follower_username) REFERENCES users(username),
		FOREIGN KEY (following_username) REFERENCES users(username),
		PRIMARY KEY (follower_username, following_username)
		);`

	SaveUserQuery = `INSERT INTO users (id, username) VALUES (?,?);`

	FollowUserQuery = `INSERT INTO followers (follower_username,following_username) VALUES (?,?);`
)
