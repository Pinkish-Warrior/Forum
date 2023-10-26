package forum

import (
	"log"
	"net/http"
	"strconv"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Get the authenticated user data
		userData := GetAuthenticatedUserData(r)

		// If the user is not authenticated, redirect to the login page
		if !userData.IsAuthenticated {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		postID, err := strconv.Atoi(r.FormValue("post_id"))
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}
		content := r.FormValue("content")

		// Retrieve the user ID from the session
		userID, ok := GetAuthenticatedUserID(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		err = AddComment(postID, userID, content)
		if err != nil {
			http.Redirect(w, r, "/error/500", http.StatusSeeOther)
			log.Println(err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/error/405", http.StatusSeeOther)
}
