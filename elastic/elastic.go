package elastic

import (
    "webapp/helpers"
    "webapp/constants"
    "webapp/opensearch"
    "webapp/types"
    "net/http"
    "html/template"
    "strings"
    "github.com/gorilla/mux"
)

// Handler function that takes additional parameters
func GetPage(w http.ResponseWriter, r *http.Request) {
    // Retrieve URL parameters
    vars := mux.Vars(r) // Get the URL parameters
    isPage := helpers.InArrayStrings(strings.Split(constants.Pages, "|"), vars["page"])

    // Here, you can check for specific conditions to send a 404
    if !isPage {
        helpers.PageNotFound(w)
        return
    }
    //Set Templates
    tmpl := template.Must(template.ParseFiles(
        constants.GetTemplatePath() + "view/layout.gohtml",
        constants.GetTemplatePath() + "view/footer.gohtml",
        constants.GetTemplatePath() + "view/pages/" + strings.ReplaceAll(vars["page"], "-", "_") + ".gohtml",
    ))

    tmpl.ExecuteTemplate(w, "layout.gohtml", helpers.MergeWithCommons(vars["page"], []types.Document{}, true))
}

// Home functions
func Home(w http.ResponseWriter, r *http.Request) {
    //Get Docs
    documents := opensearch.GetHomeSearchQuery()

    //Set Templates
    tmpl := template.Must(template.ParseFiles(
        constants.GetTemplatePath() + "view/layout.gohtml",
        constants.GetTemplatePath() + "view/footer.gohtml",
        constants.GetTemplatePath() + "view/main.gohtml",
    ))

    tmpl.ExecuteTemplate(w, "layout.gohtml", helpers.MergeWithCommons("", documents, false))
}


func GetPost(w http.ResponseWriter, r *http.Request) {
    // Retrieve URL parameters
    vars := mux.Vars(r) // Get the URL parameters
    
    //Get Docs
    documents := opensearch.GetPostQuery(vars["post"])

    //Set Templates
    tmpl := template.Must(template.ParseFiles(
        constants.GetTemplatePath() + "view/layout.gohtml",
        constants.GetTemplatePath() + "view/footer.gohtml",
        constants.GetTemplatePath() + "view/post.gohtml",
    ))

    tmpl.ExecuteTemplate(w, "layout.gohtml", helpers.MergeWithCommons("", documents, false))
}

// func ElasticSearchPing(w http.ResponseWriter, r *http.Request, opensearchURL string, index string){
//      fmt.Println("Start");

//     client, err := newOpenSearchClient([]string{opensearchURL}, opensearchURL)
    
//     if err != nil {
//             fmt.Println("Client Error");
//     }
    

//     // Search for the document.
//     content := strings.NewReader(`{
//         "size": 50,
//         "query": {
//             "match_all": {}
//         },
//         "sort": [
//             {
//               "_id": {
//                 "order": "desc"
//               }
//             }
//         ],
//         "_source": ["title", "content", "createdat", "slug"]
//     }`)

//     ctx := context.Background()

//     searchResp, err := client.Search(
//         ctx,
//         &opensearchapi.SearchReq{
//             Indices: []string{index},
//             Body: content,
//         },
//     )
//     if err != nil {
//         fmt.Printf("Error Search")
//         fmt.Printf(err.Error());
//         return
//     }


//     fmt.Printf("Search hits: %v\n", searchResp.Hits.Total.Value)

//     var documents []Document
//     for _, hit := range searchResp.Hits.Hits {
//         var doc Document
//         // Assuming the source is in JSON format and needs to be decoded
//         if err := json.Unmarshal(hit.Source, &doc); err != nil {
//             continue
//         }

//         maxLength := 225
        
//         if len(doc.Content) > maxLength {
//             doc.Content = doc.Content[:maxLength]
//         }

//         t := time.Unix(doc.Createdat, 0)
//         doc.TimeDate = t.Format("January 2, 2006")
//         doc.Author =   author
//         fmt.Println(doc.TimeDate)

//         // Add the document to the slice
//         documents = append(documents, doc)
//     }
//     p := Site{Title: "Phish Bug", Categories: categories, Data: documents}

//     templatePath := filepath.Join("/home/ec2-user/view", "*.gohtml")

//     tmpl, err := template.ParseGlob(templatePath)

//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }


//     tmpl.ExecuteTemplate(w, "layout.gohtml", p)

// }



// func IndexDocument(w http.ResponseWriter, r *http.Request, opensearchURL string, IndexName string) {
    
//     //Check Post Method
//     checkPost(r, w)

//     var document Document
    
//     err := json.NewDecoder(r.Body).Decode(&document)
    
//     if err != nil {
//         http.Error(w, "Bad request", http.StatusBadRequest)
//         return
//     }

//     //Seo Slug
//     document.Slug = seo.GenerateSlug(document.Title)

//     fmt.Println("Start");

//     client, err := newOpenSearchClient([]string{opensearchURL}, opensearchURL)
    
//     if err != nil {
//             fmt.Println("Client Error");
//     }

//     ctx := context.Background()

//     insertResp, err := client.Index(
//         ctx,
//         opensearchapi.IndexReq{
//             Index:      IndexName,
//             Body:       opensearchutil.NewJSONReader(&document),
//             Params: opensearchapi.IndexParams{
//                 Refresh: "true",
//             },
//         },
//     )
//     if err != nil {
//         w.Write([]byte(fmt.Sprintf("Error Creating Document %s\n", err.Error())))
//     }

//     w.Write([]byte(fmt.Sprintf("Created document in %s\n  ID: %s\n", insertResp.Index, insertResp.ID)))
// }

// func newOpenSearchClient(addresses []string, opensearchURL string) (*opensearchapi.Client, error) {
        
//         client, err := opensearchapi.NewClient(
//             opensearchapi.Config{
//                 Client: opensearch.Config{
//                     Transport: &http.Transport{
//                         TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // For testing only. Use certificate for validation.
//                     },
//                     Addresses: []string{opensearchURL},
//                     Username:  "kunalthool", // For testing only. Don't store credentials in code.
//                     Password:  "P@ssw0rd007bonde",
//                 },
//             },
//         )
//         if err != nil {
//                 return nil, fmt.Errorf("Error creating OpenSearch client: %s", err)
//         }
//         return client, nil
// }


// func CreateIndex(w http.ResponseWriter, r *http.Request, opensearchURL string, index string) {
    
//     //Check Post Method
//     checkPost(r, w)

//     var document Document
    
//     err := json.NewDecoder(r.Body).Decode(&document)
    
//     if err != nil {
//         http.Error(w, "Bad request", http.StatusBadRequest)
//         return
//     }

//     //Seo Slug
//     document.Slug = seo.GenerateSlug(document.Title)

//     fmt.Println("Start");

//     client, err := newOpenSearchClient([]string{opensearchURL}, opensearchURL)
    
//     if err != nil {
//             fmt.Println("Client Error");
//     }
        
    
//     if err != nil {
//         fmt.Println("Error");
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }

//     ctx := context.Background()


//    // Define index mapping.
//     // Note: these particular settings (eg, shards/replicas)
//     // will have no effect in AWS OpenSearch Serverless
//     mapping := strings.NewReader(`{
//         "settings": {
//             "index": {
//                 "number_of_shards": 1
//             }
//         }
//     }`)

//     // Create an index with non-default settings.
//     createIndexResponse, err := client.Indices.Create(
//         ctx,
//         opensearchapi.IndicesCreateReq{
//             Index: index,
//             Body:  mapping,
//         },
//     )

//     var opensearchError *opensearch.StructError

//     // Load err into opensearch.Error to access the fields and tolerate if the index already exists
//     if err != nil {
//         if errors.As(err, &opensearchError) {
//             if opensearchError.Err.Type != "resource_already_exists_exception" {
//                 http.Error(w, err.Error(), http.StatusInternalServerError)
//                 fmt.Println("Start R resource_already_exists_exception");
//                 return
//             }
//         } else {
//             fmt.Println("Start R resource_already_exists_exception");
//             return
//         }
//     }

//     fmt.Printf("Created Index: %s\n  Shards Acknowledged: %t\n", createIndexResponse.Index, createIndexResponse.ShardsAcknowledged)

//     insertResp, err := client.Index(
//         ctx,
//         opensearchapi.IndexReq{
//             Index:      index,
//             Body:       opensearchutil.NewJSONReader(&document),
//             Params: opensearchapi.IndexParams{
//                 Refresh: "true",
//             },
//         },
//     )


//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         fmt.Println("Start R Before Create");
//         return
//     }

//     w.Write([]byte(fmt.Sprintf("Created document in %s\n  ID: %s\n", insertResp.Index, insertResp.ID)))
// }

// func DeleteDocument(w http.ResponseWriter, r *http.Request, opensearchURL string, index string, docId string) {

//     if r.Method != http.MethodPost {
//             http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//             return
//     }

//     fmt.Println("Start");

//     client, err := newOpenSearchClient([]string{opensearchURL}, opensearchURL)
    
//     if err != nil {
//         fmt.Println("Client Error")
//         return
//     }

//     ctx := context.Background()

//     // Delete the document.
//     deleteReq := opensearchapi.DocumentDeleteReq{
//         Index:      index,
//         DocumentID: docId,
//     }

//     deleteResponse, err := client.Document.Delete(ctx, deleteReq)
//     if err != nil {
//         fmt.Println("Error In deleting");
//         return
//     }
//     fmt.Printf("Deleted document: %t\n", deleteResponse.Result == "deleted")
// }

// func DeleteIndices(w http.ResponseWriter, r *http.Request, opensearchURL string) {

//     checkPost(r, w)

//     indexName := r.URL.Query().Get("index");

//     ctx := context.Background()

//     // Delete previously created index.
//     deleteIndex := opensearchapi.IndicesDeleteReq{Indices: []string{indexName}}

//     client, err := newOpenSearchClient([]string{opensearchURL}, opensearchURL)

//     deleteIndexResp, err := client.Indices.Delete(ctx, deleteIndex)


//     fmt.Printf("Deleted index: %t\n", deleteIndexResp.Acknowledged)

//     // Try to delete the index again which fails as it does not exist
//     _, err = client.Indices.Delete(ctx, deleteIndex)

//     var opensearchError *opensearch.StructError

//     // Load err into opensearchapi.Error to access the fields and tolerate if the index is missing
//     if err != nil {
//         if errors.As(err, &opensearchError) {
//             if opensearchError.Err.Type != "index_not_found_exception" {
//                 fmt.Printf("Deleted index: Error Index not found")
//                 http.Error(w, "Index not found", http.StatusInternalServerError)
//                 return
//             }
//         } else {
//             http.Error(w, "Some Error", http.StatusInternalServerError)
//             return
//         }
//     }

//     w.Write([]byte(fmt.Sprintf("Deleted index: Deleted: %s", indexName)))
    
// }


// func checkPost(r *http.Request, w http.ResponseWriter) {

//     // Check if the request method is POST
//     if r.Method != http.MethodPost {
//         http.Error(w, "Method Not Allowed", http.StatusInternalServerError)
//         return
//     }
// }
