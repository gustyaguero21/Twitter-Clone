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

func TestFollowUserController(t *testing.T) {
	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateFollowerTable)

	if err != nil {
		t.Fatal("Error creating followers table")
	}

	followService := services.NewFollowerService(&repository.Repository{Db: db})
	followController := NewFollowerController(*followService)

	follow := models.Followers{
		FollowingUsername: "alicia",
	}

	followJSON, _ := json.Marshal(follow)

	//act

	req, _ := http.NewRequest("POST", "/follow-user/pepe", bytes.NewBuffer(followJSON))

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router := gin.Default()

	router.POST("/follow-user/:username", followController.FollowUserController)

	router.ServeHTTP(w, req)

	var response models.Followers

	err = json.Unmarshal(w.Body.Bytes(), &response)

	if err != nil {
		t.Fatal("Error unmarshaling response")
	}

	//asserts

	assert.Equal(t, response.FollowingUsername, follow.FollowerUsername)
}
