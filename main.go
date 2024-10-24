package main

import (
    "net/http"
    "webapp/elastic"
    "webapp/opensearch"
    "webapp/helpers"
    "webapp/user"
    "fmt"
    "webapp/auth"

    "github.com/gorilla/mux"

)


func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, modify as needed
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        // Handle preflight requests
        if r.Method == http.MethodOptions {
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

// Middleware to print the request URI
func printRequestURI(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("Requested Method: %s, Requested URI: %s\n", r.Method, r.RequestURI) // Print method and URI
        next.ServeHTTP(w, r) // Call the next handler
    })
}

// Custom handler for method not allowed
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func main() {
    //Initiate Router
    r := mux.NewRouter()

    // Use the middleware
    r.Use(printRequestURI)

    //Home 
    r.HandleFunc("/", elastic.Home).Methods("GET")

    //Pages
    r.HandleFunc("/page/{page}",   elastic.GetPage).Methods("GET")

    //Contact us Form
    r.HandleFunc("/contact-us",   elastic.GetPage).Methods("POST")

    //Login Page
    r.HandleFunc("/signup-login", user.LoginPage).Methods("GET")

    //Pages
    r.HandleFunc("/{post}",   elastic.GetPost).Methods("GET")



/////////////////////////////////////////////Admin Routes///////////////////////////////////////
    //AdminPages
    r.HandleFunc("/api-phish-bug/login",  auth.LoginHandler).Methods("POST")

    // Define protected route with the AuthMiddleware
    r.Handle("/api-phish-bug/indicess", auth.AuthMiddleware(http.HandlerFunc(opensearch.GetIndexes))).Methods("GET")

    // Define protected route with the AuthMiddleware
    r.Handle("/api-phish-bug/index/{name}", auth.AuthMiddleware(http.HandlerFunc(opensearch.GetIndex))).Methods("GET")

    // Define protected route with the AuthMiddleware
    r.Handle("/api-phish-bug/docs/{index}", auth.AuthMiddleware(http.HandlerFunc(opensearch.GetDocs))).Methods("GET")

/////////////////////////////////////////////Admin Routes End///////////////////////////////////////


    // // Serve the robots.txt file
    // http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
    //         // Set the Content-Type header to text/plain
    //         w.Header().Set("Content-Type", "text/plain")

    //         // Write the content of robots.txt
    //         _, err := w.Write([]byte(`User-agent: *
    //         Disallow: /private/
    //         Allow: /public/
    //         Sitemap: https://www.example.com/sitemap.xml
    //         `))
    //     if err != nil {
    //         http.Error(w, "Unable to write response", http.StatusInternalServerError)
    //     }
    // })
  
    // http.HandleFunc("/", controller.Index)


    // http.HandleFunc("/admin-e-data", func(w http.ResponseWriter, r *http.Request) {
    //     elastic.ElasticSearchPing(w, r, opensearchURL, index)
    // })

    // http.HandleFunc("/admin-e-index-delete", auth.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
    //     elastic.DeleteIndices(w, r, opensearchURL)
    // }))
    

    // //Add New Document
    // http.HandleFunc("/admin-e-data-add", auth.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
    //     elastic.IndexDocument(w, r, opensearchURL, index)
    // }))

    // http.HandleFunc("/admin-e-data-create-index", auth.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
    //     elastic.CreateIndex(w, r, opensearchURL, index)
    // }))
    
     
    // fs := http.FileServer(http.Dir(constants.TemplatePath + "styles"))


    // // Set cache headers for static assets
    // http.HandleFunc("/styles/", func(w http.ResponseWriter, r *http.Request) {
    //     w.Header().Set("Cache-Control", "public, max-age=31536000, immutable") // 1 year cache
    //     fs.ServeHTTP(w, r) // Call the file server
    // })
    
    // http.HandleFunc("/items", controller.GetItems)

    // //Can be any
    // http.HandleFunc("/admin-phish-bug-login-jwt", auth.LoginHandler)

     // Custom handler for 405 errors
    r.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedHandler)
    
    // Catch-all for 404 responses
    r.NotFoundHandler = http.HandlerFunc(helpers.NotFoundHandler)
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.HandlerFunc(staticFileHandler)))
    
    // Start the server
    http.Handle("/",  enableCORS(r))

    http.ListenAndServe("0.0.0.0:80", nil)
}


func staticFileHandler(w http.ResponseWriter, r *http.Request) {
    // Set cache headers
    w.Header().Set("Cache-Control", "public, max-age=86400") // Cache for 1 day

    // Serve the file
    http.ServeFile(w, r, "static/"+r.URL.Path)
}
