package opensearch

import (
    "context"
    "log"
    "strings"
    // "time"
    "fmt"
    "encoding/json"
    "webapp/constants"
    "webapp/types"

    "github.com/opensearch-project/opensearch-go"
    "github.com/opensearch-project/opensearch-go/opensearchapi"
)

//Open search Client
func getOpenSearchClient() (*opensearch.Client, error) {
    // Define custom client options
    cfg := opensearch.Config{
        Addresses: constants.GetOpenSearchAddress(),
        Username: constants.GetENVKey("OPEN_SOURCE_USERNAME"), // Optional: for basic authentication
        Password: constants.GetENVKey("OPEN_SOURCE_PASSWORD"),  // Optional: for basic authentication
    }

    client, err := opensearch.NewClient(cfg)

    if err != nil {
        log.Fatalf("Error creating OpenSearch client: %s", err)
    }

    return client, err
}

//Get Query For OPen Search
func GetHomeSearchQuery() []types.Document{
    
    //Get Search Request Here
    searchResp := getSearchRequest(getSearchQueryForMainPage, "")

    return getOpenSourceDoc(searchResp)
}

func GetPostQuery(slug string) []types.Document{
fmt.Print("Print Slug");
    fmt.Print(slug);
    //Get Search Request Here
    searchResp := getSearchRequest(getSearchQueryForPostPage, slug)

    return getOpenSourceDoc(searchResp)
}

func getSearchRequest(op types.Operation, slug string) *opensearchapi.Response{

    client, err := getOpenSearchClient()

    // Example: Search for documents
    searchReq := opensearchapi.SearchRequest{
        Index:  []string{constants.GetENVKey("OPEN_SEARCH_INDEX")},
        Body:  op(slug),
    }

    searchRes, err := searchReq.Do(context.Background(), client)
    
    if err != nil {
        log.Fatalf("Error searching: %s", err)
    }

    // Check the response
    if searchRes.IsError() {
        log.Fatalf("Error in response: %s", searchRes)
        panic(searchRes.IsError)
    }

    return searchRes
}

//Generate Query For main page
func getSearchQueryForMainPage(slug string) *strings.Reader{
    return strings.NewReader(
            `{
                 "size": 50,
                 "query": {
                     "match_all": {}
                 },
                 "sort": [
                     {
                       "_id": {
                         "order": "desc"
                       }
                     }
                 ],
                 "_source": ["title", "content", "createdat", "slug"]
             }`,
        )
}

//Generate Query For main page
func getSearchQueryForPostPage(slug string) *strings.Reader{
    // Create the wildcard query
    // Create the query
    query := map[string]interface{}{
        "query": map[string]interface{}{
            "term": map[string]interface{}{
                "type": "post",  // Change to your slug value
            },
        },
        "_source": []string{"title", "content", "createdat", "slug"},
    }

    // Convert the query to JSON
    queryJSON, err := json.Marshal(query)
    if err != nil {
        log.Fatalf("Error marshaling JSON: %s", err)
    }
    return strings.NewReader(string(queryJSON))    
}

//Get Results from the doc
func getOpenSourceDoc(searchResp *opensearchapi.Response) []types.Document{

     // Parse the response body
    var response map[string]interface{}
    if err := json.NewDecoder(searchResp.Body).Decode(&response); err != nil {
        fmt.Println("Error parsing response:", err)
    }

    // Extract the hits
    hits := response["hits"].(map[string]interface{})["hits"].([]interface{})

    fmt.Print(response["hits"])

    var documents []types.Document
     // Loop through the hits and extract the data
    for _, hit := range hits {
        if hitMap, ok := hit.(map[string]interface{}); ok {
            // Extract the "_id" field
            id := hitMap["_id"].(string)

            fmt.Println(id);

            // Extract other fields from the "_source" if available
            source  := hitMap["_source"].(map[string]interface{})
            title   := source["title"].(string) // Adjust based on your document structure
            slug    := source["slug"].(string) // Adjust based on your document structure
            content := source["content"].(string) // Adjust based on your document structure

            // Create a new Document and append to the slice
            documents = append(documents, types.Document{
                ID:       id,
                Title:    title,
                Slug:     slug,
                Content:  content,
            })
        }
    }
    return documents
}