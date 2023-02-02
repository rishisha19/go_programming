package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	router = setupRoutes()
	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
// func getRouter() *gin.Engine {
// 	r := setupRoutes()

// 	return r
// }

func TestGetAlbumsRoute(t *testing.T) {
	//router := getRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	var response []album
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.NotEmpty(t, response)
	for i, a := range response {

		assert.Equal(t, a, albums[i])
	}
	assert.Nil(t, err)
	//assert.True(t, exists)
}

func TestGetAlbumsByIdRoute(t *testing.T) {
	//router := setupRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	var response album
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.NotEmpty(t, response)

	assert.Equal(t, response.ID, "1")
	assert.Equal(t, response.Title, "Blue Train")
	assert.Nil(t, err)
	//assert.True(t, exists)
}

func TestGetAlbumsByIdRoute_NotFound(t *testing.T) {
	//router := setupRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/10", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.NotEmpty(t, response)
	fmt.Println(response)
	assert.Equal(t, response["message"], "album not found with ID:10")
	// assert.Equal(t, response.Title, "Blue Train")
	assert.Nil(t, err)
	//assert.True(t, exists)
}

func TestAddAlbums(t *testing.T) {
	//router := setupRoutes()
	w := httptest.NewRecorder()
	jsonValue, _ := json.Marshal(album{ID: "5", Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99})

	req, _ := http.NewRequest("POST", "/albums", bytes.NewReader(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	// Convert the JSON response to a map
	var response album
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.NotEmpty(t, response)
	fmt.Println(response)
	assert.Equal(t, response.ID, "5")
	// assert.Equal(t, response.Title, "Blue Train")
	assert.Nil(t, err)
	//assert.True(t, exists)
}
