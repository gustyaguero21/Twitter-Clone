package config

//Router params

const (
	Port = ":8080"
)

//Database params

const (
	Driver = "sqlite"
	DbPath = "./internal/data/database.db"

	MaxContentLength = 280
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

	CreateTweetsTable = `CREATE TABLE IF NOT EXISTS tweets (
		id TEXT PRIMARY KEY NOT NULL,
		username TEXT NOT NULL,
		content VARCHAR(280) NOT NULL,
		posted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (username) REFERENCES users(username));`

	SaveUserQuery = `INSERT INTO users (id, username) VALUES (?,?);`

	ExistUserQuery = `SELECT EXISTS(SELECT 1 FROM users WHERE id = ? OR username = ?);`

	FollowUserQuery = `INSERT INTO followers (follower_username,following_username) VALUES (?,?);`

	ShowFollowersQuery = `SELECT following_username FROM followers WHERE follower_username = ?;`

	CreatePostQuery = `INSERT INTO tweets (id, username, content) VALUES (?,?,?);`

	TimelineQuery = `SELECT t.username, t.content, t.posted_at
		FROM tweets t
		JOIN followers f ON t.username = f.following_username
		WHERE f.follower_username = ?
		ORDER BY t.posted_at DESC;`
)
