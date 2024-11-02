package services

import (
	"context"
	"database/sql"
	"testing"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"
	"twitter-clone/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {

	//given

	ctx := context.Background()

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateTweetsTable)

	if err != nil {
		t.Fatal("Error creating tweets table")
	}

	tweet := models.Tweets{
		User:    "pepe",
		Content: "Hola mundo",
	}

	postService := NewTweetService(&repository.Repository{Db: db})

	//act

	createdPost, err := postService.CreatePost(ctx, tweet.User, tweet.Content)

	//asserts

	assert.NoError(t, err)

	assert.Equal(t, createdPost.Content, tweet.Content)
}
