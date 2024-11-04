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

func Test_FollowUser_Success(t *testing.T) {

	//given

	ctx := context.Background()

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateFollowerTable)

	if err != nil {
		t.Fatal("Error creating tweets table")
	}

	follow := models.Followers{
		FollowerUsername:  "pepe",
		FollowingUsername: "alicia",
	}

	followService := NewFollowerService(&repository.Repository{Db: db})

	//act

	_, err = followService.FollowUser(ctx, follow)

	//asserts

	assert.NoError(t, err)

}

func Test_FollowUser_Error(t *testing.T) {

	//given

	ctx := context.Background()

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateFollowerTable)

	if err != nil {
		t.Fatal("Error creating tweets table")
	}

	follow := models.Followers{
		FollowerUsername:  "pepe",
		FollowingUsername: "alicia",
	}

	followService := NewFollowerService(&repository.Repository{Db: db})

	//act

	_, err = followService.FollowUser(ctx, follow)

	if err != nil {
		t.Fatal("Error following user")
	}

	_, err = followService.FollowUser(ctx, follow)

	//asserts

	assert.Error(t, err)

}

func Test_Following_Success(t *testing.T) {

	//given

	ctx := context.Background()

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateFollowerTable)

	if err != nil {
		t.Fatal("Error creating tweets table")
	}

	follow := models.Followers{
		FollowerUsername:  "pepe",
		FollowingUsername: "alicia",
	}

	followService := NewFollowerService(&repository.Repository{Db: db})

	_, err = followService.FollowUser(ctx, follow)

	if err != nil {
		t.Fatal("Error following user")
	}

	//act

	following, err := followService.ShowFollowers(ctx, "pepe")

	//asserts

	assert.NoError(t, err)
	assert.Len(t, following, 1)

}

func Test_Following_Error(t *testing.T) {

	//given

	ctx := context.Background()

	db, err := sql.Open(config.Driver, ":memory:")

	if err != nil {
		t.Fatal("Error creating database in memory")
	}

	defer db.Close()

	_, err = db.Exec(config.CreateFollowerTable)

	if err != nil {
		t.Fatal("Error creating tweets table")
	}

	followService := NewFollowerService(&repository.Repository{Db: db})

	//act

	following, err := followService.ShowFollowers(ctx, "pepe")

	//asserts

	assert.NoError(t, err)
	assert.Len(t, following, 0)

}
