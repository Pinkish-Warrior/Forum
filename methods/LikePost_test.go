package forum

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestLikePost(t *testing.T) {
	// Create a new mock DB connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	defer db.Close()

	// Replace the global 'db' variable with the mock
	globalDB := db
	defer func() {
		globalDB = nil
	}()
	fmt.Println(globalDB)
	// Create a test case
	testCases := []struct {
		name              string
		userID, postID    int
		likeStatus        bool
		queryExpectations func()
		expectedError     bool
	}{
		{
			name:       "User likes a post",
			userID:     1,
			postID:     123,
			likeStatus: true,
			queryExpectations: func() {
				mock.ExpectQuery("SELECT like_status FROM likes_dislikes").
					WithArgs(1, 123).
					WillReturnRows(sqlmock.NewRows([]string{"like_status"}).AddRow(true))

				mock.ExpectExec("DELETE FROM likes_dislikes").
					WithArgs(1, 123).
					WillReturnResult(driver.ResultNoRows)

				mock.ExpectExec("UPDATE likes_dislikes").
					WithArgs(true, 1, 123).
					WillReturnResult(driver.ResultNoRows)
			},
			expectedError: false,
		},
		{
			name:       "User dislikes a post",
			userID:     1,
			postID:     123,
			likeStatus: false,
			queryExpectations: func() {
				mock.ExpectQuery("SELECT like_status FROM likes_dislikes").
					WithArgs(1, 123).
					WillReturnRows(sqlmock.NewRows([]string{"like_status"}).AddRow(false))

				mock.ExpectExec("DELETE FROM likes_dislikes").
					WithArgs(1, 123).
					WillReturnResult(driver.ResultNoRows)

				mock.ExpectExec("UPDATE likes_dislikes").
					WithArgs(true, 1, 123).
					WillReturnResult(driver.ResultNoRows)
			},
			expectedError: false,
		},
		{
			name:       "User has not liked the post",
			userID:     1,
			postID:     123,
			likeStatus: true,
			queryExpectations: func() {
				mock.ExpectQuery("SELECT like_status FROM likes_dislikes").
					WithArgs(1, 123).
					WillReturnError(sql.ErrNoRows)

				mock.ExpectExec("INSERT INTO likes_dislikes").
					WithArgs(1, 123, true, 1).
					WillReturnResult(driver.ResultNoRows)
			},
			expectedError: false,
		},
		{
			name:   "Database query error",
			userID: 1,
			postID: 123,
			queryExpectations: func() {
				mock.ExpectQuery("SELECT like_status FROM likes_dislikes").
					WithArgs(1, 123).
					WillReturnError(sql.ErrConnDone)
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up query expectations
			tc.queryExpectations()

			// Call the function
			err := LikePost(tc.userID, tc.postID)

			// Check the error result
			if (err != nil) != tc.expectedError {
				t.Errorf("Expected error: %v, got error: %v", tc.expectedError, err)
			}

			// Ensure all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
