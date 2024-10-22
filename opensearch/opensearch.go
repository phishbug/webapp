package opensearch

import (
    "context"
    "log"
    "strings"
    "time"
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
    searchResp := getSearchRequest(getSearchQueryForMainPage)

    return getOpenSourceDoc(searchResp)
}

func getSearchRequest(op types.Operation) *opensearchapi.Response{

    client, err := getOpenSearchClient()

    // Example: Search for documents
    searchReq := opensearchapi.SearchRequest{
        Index:  []string{constants.GetENVKey("OPEN_SEARCH_INDEX")},
        Body:  op(),
    }

    searchRes, err := searchReq.Do(context.Background(), client)
    
    if err != nil {
        log.Fatalf("Error searching: %s", err)
    }

    defer searchRes.Body.Close()

    // Check the response
    if searchRes.IsError() {
        log.Fatalf("Error in response:", searchRes)
        panic(searchRes.IsError)
    }

    return searchRes
}

//Generate Query For main page
func getSearchQueryForMainPage() *strings.Reader{
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

//Get Results from the doc
func getOpenSourceDoc(searchResp *opensearchapi.Response) []types.Document{
    var response map[string]interface{}
    var documents []types.Document

    if err := json.NewDecoder(searchResp.Body).Decode(&response); err != nil {
        log.Fatalf("Error parsing response:", err)
        panic(err)
    }

    // Extract the hits
    hits := response["hits"].(map[string]interface{})["hits"].([]interface{})

    for _, hit := range hits {
        var doc types.Document
        
        // Each hit is a map, so we need to type assert it
        if hitMap, ok := hit.(map[string]interface{}); ok {
            // Extract the "_id" field
            if id, exists := hitMap["_id"]; exists {
                t := time.Unix(doc.Createdat, 0)
                doc.TimeDate = t.Format(constants.DATE_FORMAT)
                doc.Author   =   constants.GetAuthor()
                doc.ID       =   id.(string)
                documents    = append(documents, doc)
            }
        }
    }

    return documents
}