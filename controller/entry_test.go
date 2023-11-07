package controller

import (
	"net/http"
	"testing"

	"github.com/nerd500/diary_api/model"

	"github.com/stretchr/testify/assert"
)

func TestAddEntry(t *testing.T) {
	newEntry := model.Entry{
		Content: "This is a test entry :)",
	}
	writer := makeRequest("POST", "/api/entry", newEntry, true)
	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestGetAllEntries(t *testing.T) {
	writer := makeRequest("GET", "/api/entry", nil, true)
	assert.Equal(t, http.StatusOK, writer.Code)
}
