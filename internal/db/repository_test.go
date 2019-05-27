package db

import (
	"cadence-service/internal/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	dbURL = "devuser:liverpool123@tcp(127.0.0.1:3306)/players"
)

func TestNewMySQLRepository(t *testing.T) {

	testRepo, err := NewMySQLRepository(dbURL)

	if err != nil {

		errorMsg := fmt.Errorf("cannot get repo %v", err)

		assert.Fail(t, errorMsg.Error())

	}

	fmt.Println(testRepo)
	assert.Equal(t, nil, err)

}

func TestReadOneReturnsModel(t *testing.T) {

	testRepo, err := NewMySQLRepository(dbURL)

	if err != nil {
		assert.Fail(t, "Cannot get repo")
	}

	actual, err := testRepo.ReadOne("gerrard")

	expected := model.DatastoreModel{Name: "gerrard", Team: "liverpool", Number: 8}

	if err != nil {
		assert.Fail(t, err.Error())

	}

	if actual == nil {

	}

	assert.Equal(t, expected, actual)

}

func TestInsertPersistsRecord(t *testing.T) {

	testRepo, err := NewMySQLRepository(dbURL)

	if err != nil {
		assert.Fail(t, "Cannot get repo")
	}

	testPlayer := model.DatastoreModel{Name: "gerrard", Team: "liverpool", Number: 8}

	err = testRepo.Insert(&testPlayer)

	if err != nil {

		assert.Fail(t, "Cannot insert player")

	}

	fmt.Println("Inserted")
	assert.Equal(t, nil, err)

}
