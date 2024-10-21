package main

import (
    "net/http"
    // "webapp/controller"
    "webapp/elastic"
    // "webapp/constants"
    "webapp/helpers"
    // "webapp/auth"

    "github.com/gorilla/mux"

)



func main() {
    //Initiate Router
    r := mux.NewRouter()

    //Pages
    r.HandleFunc("/page/{page}", elastic.GetPage).Methods("GET")

    //Contact us Form
    r.HandleFunc("/contact-us}", elastic.GetPage).Methods("POST")

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

    // Catch-all for 404 responses
    r.NotFoundHandler = http.HandlerFunc(helpers.NotFoundHandler)
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.HandlerFunc(staticFileHandler)))
    
    // Start the server
    http.Handle("/", r)

    http.ListenAndServe("0.0.0.0:80", nil)
}


func staticFileHandler(w http.ResponseWriter, r *http.Request) {
    // Set cache headers
    w.Header().Set("Cache-Control", "public, max-age=86400") // Cache for 1 day

    // Serve the file
    http.ServeFile(w, r, "static/"+r.URL.Path)
}