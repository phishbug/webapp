package controller

import (
    "context"
    "encoding/json"
    "net/http"
    "html/template"
    "webapp/model"
    "fmt"
    "os"

    "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/v8/esapi"
)

type Person struct {
    Name string
    Age  int
    Title string
}

type Response struct {
    Hits struct {
        Hits []struct {
            Source model.Item `json:"_source"`
        } `json:"hits"`
    } `json:"hits"`
}

func Index(w http.ResponseWriter, r *http.Request) {

    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    fmt.Println("Current working directory:", dir)

    // Render the view with the items
    p := Person{Name: "John", Age: 30, Title: "Phish Bug"}

   // Render the view with the items
    tmpl, err := template.ParseFiles("view/layout.gohtml", "view/items_view.gohtml")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    tmpl.ExecuteTemplate(w, "layout.gohtml", p)
}


func GetItems(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Start")
    // Create a new Elasticsearch client
    es, err := elasticsearch.NewDefaultClient()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
     fmt.Println("Start 1")
    // Search for items in the "items" index
    res, err := esapi.SearchRequest{
        Index: []string{"posts"},
        Body:  nil,
    }.Do(context.Background(), es)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    // Check if the response is an error
    if res.IsError() {
        fmt.Println("Start  IS Error null")
        http.Error(w, res.String(), http.StatusInternalServerError)
        return
    }

    // Decode the response into our Response struct
    var response Response
    if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }


    // Prepare the data for rendering
    items := make([]model.Item, len(response.Hits.Hits))
    for i, hit := range response.Hits.Hits {
        items[i] = hit.Source
    }

    // Render the view with the items
    tmpl, err := template.ParseFiles("view/layout.gohtml", "view/items_view.gohtml")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := struct {
        Title string
        Items []model.Item
    }{
        Title: "Code Base",
        Items: items,
    }

    tmpl.ExecuteTemplate(w, "layout.gohtml", data)
}



