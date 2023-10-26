package forum

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundHandler(t *testing.T) {
	// Create a test HTTP request with a path that should return a 404
	request := httptest.NewRequest("GET", "/nonexistent", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the NotFoundHandler function
	handled := NotFoundHandler(response, request)

	// Check that the response code is 404
	assert.Equal(t, http.StatusNotFound, response.Code)

	// Check that the function returns true (indicating that the request was handled)
	assert.True(t, handled)

	// Create a test HTTP request with a path that should not return a 404
	request = httptest.NewRequest("GET", "/", nil)

	// Reset the response recorder
	response = httptest.NewRecorder()

	// Call the NotFoundHandler function
	handled = NotFoundHandler(response, request)

	// Check that the response code is not 404
	assert.NotEqual(t, http.StatusNotFound, response.Code)

	// Check that the function returns false (indicating that the request was not handled)
	assert.False(t, handled)
}
