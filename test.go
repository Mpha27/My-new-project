package main

import (
	"html/template"
	"net/http"
)

const htmlPage = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Task Page</title>
    <style>
        .buttons {
            position: fixed;
            top: 0;
            right: 0;
            display: flex;
            flex-direction: column;
        }
        .button-group {
            display: flex;
            justify-content: flex-end;
            margin-right: 1rem;
        }
        .middle-button {
            display: block;
            margin: 0 auto;
            margin-top: 5rem;
        }
    </style>
</head>
<body>
    <div class="buttons">
        <div class="button-group">
            <button>My tasks</button>
            <button>Book a task</button>
            <button>Account</button>
        </div>
    </div>
    <button class="middle-button">Book Your Next Task</button>
</body>
</html>
`

var templateHTML = template.Must(template.New("").Parse(htmlPage))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templateHTML.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}
