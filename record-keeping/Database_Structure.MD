## Database Structure:

## Table: users

Columns:
id (Primary Key)
email (Unique)
username
password

## Table: posts

Columns:
id (Primary Key)
title
content
user_id (Foreign Key referencing users)
category

## Table: comments

Columns:
id (Primary Key)
content
user_id (Foreign Key referencing users)
post_id (Foreign Key referencing posts)

## Table: likes_dislikes

Columns:
id (Primary Key)
user_id (Foreign Key referencing users)
post_id (Foreign Key referencing posts)
like_status (Boolean)

## Table: categories

Columns:
id (Primary Key)
name

/forum_project
│
├── app
│ ├── handlers # Request handlers
│ │ ├── auth.go # Authentication handlers (register, login,etc.)
│ │ ├── posts.go # Post related handlers
│ │ ├── comments.go # Comment related handlers
│ │ ├── filters.go # Filtering related handlers
│ │ └── ...
│ │
│ ├── models # Data models and database interactions
│ │ ├── user.go # User model
│ │ ├── post.go # Post model
│ │ ├── comment.go # Comment model
│ │ ├── category.go # Category model
│ │ └── ...
│ │
│ ├── templates # HTML templates
│ │ ├── index.html
│ │ ├── login.html
│ │ ├── register.html
│ │ ├── post.html
│ │ └── ...
│ │
│ ├── static # Static files (CSS, JS, images, etc.)
│ │ ├── style.css
│ │ ├── script.js
│ │ └── ...
│ │
│ ├── main.go # Main application entry point
│ ├── config.go # Application configuration
│ └── router.go # HTTP router setup
│
├── migrations # Database schema migrations
│ ├── 001_initial.sql
│ └── ...
│
├── Dockerfile # Docker configuration
├── docker-compose.yml # Docker compose configuration
├── go.mod
└── README.md
