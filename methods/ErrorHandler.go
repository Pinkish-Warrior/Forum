package forum

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var tmpl *template.Template

func init() {

	currentWorkingDir, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("currentWorkingDir from Error Handler:", currentWorkingDir)

	templatePath := filepath.Join(currentWorkingDir, "templates/error.html")
	// templatePath := filepath.Join(currentWorkingDir, "templates/error.html")
	fmt.Println("template path from ErrorHandler:", templatePath)

	fmt.Println("Current Working Directory from ErrorHandler:", currentWorkingDir)

	tmpl, err = template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Template Path after parsing:", templatePath)

}

// func init() {
// 	// Determine the directory of the Go source code or binary
// 	currentWorkingDir, err := os.Executable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Construct the relative path to your template directory
// 	templateDir := filepath.Join(filepath.Dir(currentWorkingDir), "templates")

// 	// Construct the template pattern based on the template directory
// 	templatePattern := filepath.Join(templateDir, "error.html")

// 	// Parse templates using the pattern
// 	tmpl, err = template.ParseGlob(templatePattern)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	//ALERT Parsing the template file in every request can be inefficient and slow down the response time. By parsing the template file once during initialization and storing it in a global variable, we can improve the performance of the application by avoiding unnecessary file parsing operations.

	// tmpl, err := template.ParseFiles("templates/error.html")
	// if err != nil {
	// 	http.Redirect(w, r, "/error/500", http.StatusSeeOther)
	// 	return
	// }

	var errorMessage string
	statusCode := http.StatusNotFound

	switch r.URL.Path {
	case "/error/404":
		errorMessage = "Page not found"
		statusCode = http.StatusNotFound
	case "/error/500":
		errorMessage = "Internal Server Error"
		statusCode = http.StatusInternalServerError
	case "/error/400":
		errorMessage = "Bad Request"
		statusCode = http.StatusBadRequest
	case "/error/405":
		errorMessage = "Method Not Allowed"
		statusCode = http.StatusMethodNotAllowed
	default:
		errorMessage = "Service not Available"
		statusCode = http.StatusServiceUnavailable
	}

	data := ErrorData{
		Error:     errorMessage,
		ErrorCode: fmt.Sprintf("%d", statusCode),
	}

	w.WriteHeader(statusCode)

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// package forum

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"text/template"
// )

// var tmpl *template.Template

// func init() {
// 	// Determine the directory of the Go source code or binary
// 	currentWorkingDir, err := os.Executable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Construct the relative path to your template directory
// 	templateDir := filepath.Join(filepath.Dir(currentWorkingDir), "templates")

// 	// Construct the template path based on the template directory
// 	templatePath := filepath.Join(templateDir, "error.html")

// 	tmpl, err = template.ParseFiles(templatePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func init() {
// 	// Construct the relative path to the template directory and the template file
// 	templateDir := "templates" // The directory containing your template
// 	templatePath := filepath.Join(templateDir, "error.html")

// 	tmpl, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(tmpl)
// }
