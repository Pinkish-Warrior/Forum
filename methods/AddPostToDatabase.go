package forum

func AddPostToDatabase(category, title, content string, userID int) error {
	// Insert the post into the database
	_, err := db.Exec("INSERT INTO posts (title, content, user_id, category) VALUES (?, ?, ?, ?)",
		title, content, userID, category)
	return err
}
