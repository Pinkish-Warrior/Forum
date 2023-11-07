package forum

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/error/405", http.StatusSeeOther)
		return
	}
	// Retrieve the category, title, and content from the form values
	category := r.FormValue("category")
	title := r.FormValue("title")
	content := r.FormValue("content")
	content = strings.TrimSpace(content)

	// Check if any of the fields are empty, if so, return a 400 error
	if category == " " || title == " " || content == " " {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Use a regular expression to check for non-empty content
	if ok, err := regexp.MatchString(`\S`, content); !ok || err != nil {
		http.Error(w, "🤔 CANNOT see your POST, make sure you write something cool!", http.StatusBadRequest)
		return
	}

	// Retrieve the user ID from the session cookie
	userID, isAuthenticated := GetAuthenticatedUserID(r)

	if !isAuthenticated {
		// If the user is not authenticated, redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Check if the category already exists
	var categoryID int
	err := db.QueryRow("SELECT id FROM categories WHERE name = ?", category).Scan(&categoryID)
	if err != nil {
		if err != sql.ErrNoRows {
			http.Redirect(w, r, "/error/500", http.StatusSeeOther)
			log.Fatal(err)
			return
		}

		// If the category doesn't exist, create it
		_, err = db.Exec("INSERT INTO categories (name) VALUES (?)", category)
		if err != nil {
			http.Redirect(w, r, "/error/500", http.StatusSeeOther)
			log.Println(err)
			return
		}

		// Retrieve the newly created category ID
		err = db.QueryRow("SELECT id FROM categories WHERE name = ?", category).Scan(&categoryID)
		if err != nil {
			http.Redirect(w, r, "/error/500", http.StatusSeeOther)
			log.Println(err)
			return
		}
	}

	// Add the post to the database
	err = AddPostToDatabase(category, title, content, userID)
	if err != nil {
		http.Redirect(w, r, "/error/500", http.StatusSeeOther)
		log.Println(err)
		return
	}

	// Redirect the user to the home page after successfully adding the post
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
