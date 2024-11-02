package repository

import (
	"database/sql"
	"testing"
	"time"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_SavePost_Success(t *testing.T) {
	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateTweetsTable)

	if err != nil {
		t.Fatal("Error creating tweets table")
	}

	postRepo := NewPostRepository(Repository{Db: db})

	tweet := models.Tweets{
		ID:       uuid.New(),
		User:     "pepe",
		Content:  "Hola mundo",
		PostedAt: time.Now(),
	}

	//act

	err = postRepo.SavePost(tweet)

	//asserts

	assert.Nil(t, err)
}

func Test_SavePost_Error(t *testing.T) {
	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateTweetsTable)

	if err != nil {
		t.Fatal("Error creating tweets table")
	}

	postRepo := NewPostRepository(Repository{Db: db})

	tweet := models.Tweets{
		ID:       uuid.New(),
		User:     "pepe",
		PostedAt: time.Now(),
	}

	//act

	err = postRepo.SavePost(tweet)

	//asserts

	assert.Error(t, err)
	assert.Equal(t, "content cannot be empty", err.Error())
}

func Test_GetTimeline_Success(t *testing.T) {
	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateTweetsTable)

	if err != nil {
		t.Fatal("Error creating tweets table")
	}

	_, err = db.Exec(config.CreateFollowerTable)

	if err != nil {
		t.Fatal("Error creating followers table")
	}

	postRepo := NewPostRepository(Repository{Db: db})

	tweet := models.Tweets{
		ID:       uuid.New(),
		User:     "pepe",
		Content:  "Hola mundo",
		PostedAt: time.Now(),
	}
	_, err = db.Exec(config.CreatePostQuery, tweet.ID, tweet.User, tweet.Content, tweet.PostedAt)

	if err != nil {
		t.Fatal("Error creating post")
	}

	//act

	_, err = postRepo.GetTimeline("pepe")

	//asserts

	assert.NoError(t, err)

}

func Test_GetTimeline_Error(t *testing.T) {
	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateTweetsTable)

	if err != nil {
		t.Fatal("Error creating tweets table")
	}

	_, err = db.Exec(config.CreateFollowerTable)

	if err != nil {
		t.Fatal("Error creating followers table")
	}

	postRepo := NewPostRepository(Repository{Db: db})

	tweet := models.Tweets{
		ID:       uuid.New(),
		User:     "pepe",
		Content:  "Hola mundo",
		PostedAt: time.Now(),
	}
	_, err = db.Exec(config.CreatePostQuery, tweet.ID, tweet.User, tweet.Content, tweet.PostedAt)

	if err != nil {
		t.Fatal("Error creating post")
	}

	//act

	tweets, err := postRepo.GetTimeline("alicia")

	//asserts

	assert.NoError(t, err)
	assert.Len(t, tweets, 0)
}
