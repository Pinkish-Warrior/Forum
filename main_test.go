package main

import (
	"database/sql"
	"fmt"
	forum "forum/methods"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// func setup() {
// 	// Initialize the test database connection
// 	testDB, err := sql.Open("sqlite3", "file::memory:?cache=shared")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db = testDB
// }

// func teardown() {
// 	// Close the test database connection
// 	if err := db.Close(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func TestMain(m *testing.M) {
// 	setup()
// 	defer teardown()

// 	// Run the tests
// 	exitCode := m.Run()

// 	os.Exit(exitCode)
// }

// func Router() *mux.Router {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/", forum.HandleMain).Methods("GET")
// 	// router.HandleFunc("/nonexistent", forum.HandleMain).Methods("GET")
// 	// router.HandleFunc("/", forum.HandleMain).Methods("POST")
// 	return router
// }

// func TestRouter(t *testing.T) {
// 	router := Router()
// 	// Create a test HTTP request
// 	request := httptest.NewRequest("GET", "/", nil)

// 	// Create a test HTTP response recorder
// 	response := httptest.NewRecorder()

// 	// Serve the HTTP request to our router. This will populate the response.
// 	router.ServeHTTP(response, request)

//		// Check the status code is what we expect.
//		if response.Code != http.StatusOK {
//			t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
//		}
//	}
//
// Declare a global variable for the database connection

// func teardown() {
// 	// Close the test database connection
// 	if err := db.Close(); err != nil {
// 		log.Fatal(err)
// 	}
// }
// func TestMain(m *testing.M) {
// 	// Initialize the test database connection
// 	testDB, err := sql.Open("sqlite3", "file::memory:?cache=shared")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Assign the test database to the global db variable
// 	db = testDB

// 	setup()
// 	defer teardown()

// 	// Run the tests
// 	exitCode := m.Run()

// 	os.Exit(exitCode)
// 	// Close the test database when testing is complete
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()
// }
// func Router() *mux.Router {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/", forum.HandleMain).Methods("GET")
// 	// router.HandleFunc("/nonexistent", forum.HandleMain).Methods("GET")
// 	// router.HandleFunc("/", forum.HandleMain).Methods("POST")
// 	return router
// }

// func TestRouter(t *testing.T) {
// 	router := Router()
// 	// Create a test HTTP request
// 	request := httptest.NewRequest("GET", "/", nil)

// 	// Create a test HTTP response recorder
// 	response := httptest.NewRecorder()

// 	// Serve the HTTP request to our router. This will populate the response.
// 	router.ServeHTTP(response, request)

// 	// Check the status code is what we expect.
// 	if response.Code != http.StatusOK {
// 		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
// 	}
// }

// package main

// import (
// 	"database/sql"
// 	_ "github.com/mattn/go-sqlite3"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"github.com/gorilla/mux"
// )

var testDB *sql.DB
var testUser forum.User // Declare a test user
// Declare the global db variable

func setup() {
	// Initialize the test database connection
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatal(err)
	}

	testDB = db // Assign the test database to the global db variable

	testUser = forum.GenerateFakeUser()
}

func teardown() {
	// Close the test database connection
	if err := testDB.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {

	setup()
	defer teardown() // Ensure teardown is called at the end of the tests

	// Initialize the main application's database
	var err1 error
	db, err1 = sql.Open("sqlite3", "database.db")

	if err1 != nil {
		log.Fatal(err1.Error())
	}
	defer db.Close()

	err1 = db.Ping()
	if err1 != nil {
		fmt.Println("error verifying connection with Ping")
	}
	// Run the tests
	exitCode := m.Run()

	os.Exit(exitCode)
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", forum.HandleMain).Methods("GET")
	// Add other routes as needed
	return router
}

func TestRouter(t *testing.T) {
	router := Router()
	// Create a test HTTP request
	request := httptest.NewRequest("GET", "/", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Serve the HTTP request to our router. This will populate the response.
	router.ServeHTTP(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}
}

// IMPORTANT

func TestHandleMain(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("GET", "/", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the HandleMain function
	forum.HandleMain(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Welcome to the Forum"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestRegisterPage(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("GET", "/register", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the RegisterPage function
	forum.RegisterPage(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Register"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestRegistration2(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("POST", "/register2", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the Registration2 function
	forum.Registration2(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Registration Successful"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestLoginHandler(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("POST", "/login", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the LoginHandler function
	forum.LoginHandler(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Login Successful"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestAddPost(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("POST", "/post-added", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the AddPost function
	forum.AddPost(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Post Added Successfully"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestLogoutHandler(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("GET", "/logout", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the LogoutHandler function
	forum.LogoutHandler(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Logout Successful"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestFilterPosts(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("GET", "/filter", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the FilterPosts function
	forum.FilterPosts(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Filtered Posts"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestProfilePage(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("GET", "/profile", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the ProfilePage function
	forum.ProfilePage(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Profile"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestAddCommentHandler(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("POST", "/add-comment", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the AddCommentHandler function
	forum.AddCommentHandler(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Comment Added Successfully"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestLikePostHandler(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("POST", "/like", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the LikePostHandler function
	forum.LikePostHandler(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Post Liked Successfully"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestDislikePostHandler(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("POST", "/dislike", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the DislikePostHandler function
	forum.DislikePostHandler(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Post Disliked Successfully"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestLikeDislikeHandler(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("POST", "/like-dislike", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the LikeDislikeHandler function
	forum.LikeDislikeHandler(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Post Liked/Disliked Successfully"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestLikeCommentHandler(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("POST", "/like-comment", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the LikeCommentHandler function
	forum.LikeCommentHandler(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Comment Liked Successfully"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
}

func TestDislikeCommentHandler(t *testing.T) {
	// Create a test HTTP request
	request := httptest.NewRequest("POST", "/dislike-comment", nil)

	// Create a test HTTP response recorder
	response := httptest.NewRecorder()

	// Call the DislikeCommentHandler function
	forum.DislikeCommentHandler(response, request)

	// Check the status code is what we expect.
	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP status %d, but got %d", http.StatusOK, response.Code)
	}

	// Check that the response body contains the expected text
	expectedText := "Comment Disliked Successfully"
	if !strings.Contains(response.Body.String(), expectedText) {
		t.Errorf("Expected response body to contain '%s', but got '%s'", expectedText, response.Body.String())
	}
	gofakeit.Name()
}