package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"twitter-clone/cmd/config"
	"twitter-clone/internal/models"
	"twitter-clone/internal/repository"
	"twitter-clone/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func Test_CreatePostController_Success(t *testing.T) {
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

	tweetService := services.NewTweetService(&repository.Repository{Db: db})
	tweetController := NewTweetController(*tweetService)

	tweet := models.Tweets{
		User:    "pepe",
		Content: "Hola mundo",
	}

	tweetJSON, _ := json.Marshal(tweet)

	//act

	req, _ := http.NewRequest("POST", "/create-tweet/pepe", bytes.NewBuffer(tweetJSON))

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router := gin.Default()

	router.POST("/create-tweet/:username", tweetController.CreatePostController)

	router.ServeHTTP(w, req)

	var response models.Tweets

	err = json.Unmarshal(w.Body.Bytes(), &response)

	if err != nil {
		t.Fatal("Error unmarshaling response")
	}

	//asserts

	assert.Equal(t, tweet.Content, response.Content)

}
