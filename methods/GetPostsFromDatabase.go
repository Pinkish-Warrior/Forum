package forum

import (
	"log"
)

func GetPostsFromDatabase() ([]Post, error) {
	rows, err := db.Query(`
        SELECT posts.id, posts.title, posts.content, posts.category, users.username 
        FROM posts
        JOIN users ON posts.user_id = users.id
        ORDER BY posts.id DESC
    `)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	// Iterate over the results and populate the posts slice
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Username)
		if err != nil {
			log.Println(err)
			continue
		}

		// Get like count for the post
		likeCount, err := GetLikeCountForPost(post.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		// Get dislike count for the post
		dislikeCount, err := GetDislikeCountForPost(post.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		post.LikeCount = likeCount
		post.DislikeCount = dislikeCount

		posts = append(posts, post)
	}

	return posts, nil
}
