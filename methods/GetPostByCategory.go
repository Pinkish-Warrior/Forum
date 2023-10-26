package forum

func GetPostsByCategory(category string) ([]Post, error) {
	// Query posts with the specified category, including username
	rows, err := db.Query(`
        SELECT p.id, p.title, p.content, p.user_id, p.category, u.username 
        FROM posts p
        JOIN users u ON p.user_id = u.id
        WHERE category=?
    `, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	// Iterate over the results and populate the posts slice
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.Category, &post.Username)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	// Check for any error during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
