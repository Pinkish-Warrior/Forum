## Tree structure suggestion:

forum/
├── main.go
├── methods/
│ ├── AddComment.go
│ ├── AddCommentHandler.go
│ ├── ...
│ ├── ClearSession_test.go
│ ├── CreateSession_test.go
│ ├── ...
├── static/
│ ├── ...
├── templates/
│ ├── ...
└── tests/
├── handlers/
│ ├── AddCommentHandler_test.go
│ ├── AddPostHandler_test.go
│ ├── ...
├── database/
│ ├── AddCommentToDatabase_test.go
│ ├── AddPostToDatabase_test.go
│ ├── ...
├── helpers/
│ ├── testhelpers.go
├── main_test.go

## Tree structure suggestion:

forum/
├── main.go
├── database/
│ ├── database.go
│ ├── database_test.go
├── methods/
│ ├── your_methods.go
│ ├── your_methods_test.go
├── templates/
│ ├── your_templates.go
├── mock/
│ ├── mock_database.go
└── go.mod
