package main

import (
    "fmt"
    "net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    // Serve a simple HTML login form
    html := `
    <!DOCTYPE html>
    <html>
    <head><title>Go Login App</title></head>
    <body>
        <h1>Simple Go Login Page</h1>
        <form action="/login" method="post">
            Username: <input type="text" name="username"><br><br>
            Password: <input type="password" name="password"><br><br>
            <input type="submit" value="Login">
        </form>
    </body>
    </html>
    `
    fmt.Fprintf(w, html)
}

func main() {
    // Handler for the main page (login form)
    http.HandleFunc("/", loginHandler)

    port := "8081"
    fmt.Printf("Starting Go server on http://localhost:%s\n", port)
    // Start the server on port 8081
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        panic(err)
    }
}
