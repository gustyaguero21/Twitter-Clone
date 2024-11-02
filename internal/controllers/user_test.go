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

func Test_CreateUserController_Success(t *testing.T) {

	//given

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateUserTable)

	if err != nil {
		t.Fatal("Error creating user table")
	}

	userService := services.NewUserService(&repository.Repository{Db: db})
	userController := NewUserController(*userService)

	user := models.Users{
		Username: "pepe",
	}

	userJSON, _ := json.Marshal(user)

	//act

	req, _ := http.NewRequest("POST", "/create-user", bytes.NewBuffer(userJSON))

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router := gin.Default()

	router.POST("/create-user", userController.CreateUserController)

	router.ServeHTTP(w, req)

	//asserts

	assert.Equal(t, http.StatusOK, w.Code)
}
