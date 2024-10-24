package opensearch

import(
	"log"
	"context"
	"encoding/json"

    
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func GetIndexed()  []map[string]interface{}{
    client, err := getOpenSearchClient()

    if err != nil {
        log.Fatalf("Error creating the client: %s", err)
    }

    // Prepare the request to get all indices
    req := opensearchapi.CatIndicesRequest{
        Format: "json",
    }

    // Execute the request
    res, err := req.Do(context.Background(), client)

    if err != nil {
        log.Fatalf("Error getting indices: %s", err)
    }
    defer res.Body.Close()

    // Check for a successful response
    if res.IsError() {
        log.Fatalf("Error: %s", res.String())
    }

    // Parse the response body
    var indices []map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&indices); err != nil {
        panic(err)
    }

   return indices
}


func GetIndexDetails(name string) (string, string, error){
	// Initialize the OpenSearch client
     client, err := getOpenSearchClient()
    if err != nil {
        log.Fatalf("Error creating OpenSearch client: %s", err)
    }

    // Specify the index name
    indexName := name // Replace with your actual index name

    // Get index settings
    settingsResponse, err := getIndexSettings(client, indexName)
    if err != nil {
        log.Fatalf("Error getting index settings: %s", err)
    }

    // Get index mappings
    mappingsResponse, err := getIndexMappings(client, indexName)
    if err != nil {
        log.Fatalf("Error getting index mappings: %s", err)
    }

    return mappingsResponse, settingsResponse, err
}

// Function to get index settings
func getIndexSettings(client *opensearch.Client, index string) (string, error) {
    req := opensearchapi.IndicesGetSettingsRequest{
        Index: []string{index},
    }

    // Perform the request
    res, err := req.Do(context.Background(), client)
    if err != nil {
        return "", err
    }
    defer res.Body.Close()

    var settings map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&settings); err != nil {
        return "", err
    }

    settingsJSON, _ := json.MarshalIndent(settings, "", "  ")
    return string(settingsJSON), nil
}

// Function to get index mappings
func getIndexMappings(client *opensearch.Client, index string) (string, error) {
    req := opensearchapi.IndicesGetMappingRequest{
        Index: []string{index},
    }

    // Perform the request
    res, err := req.Do(context.Background(), client)
    if err != nil {
        return "", err
    }
    defer res.Body.Close()

    var mappings map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&mappings); err != nil {
        return "", err
    }

    mappingsJSON, _ := json.MarshalIndent(mappings, "", "  ")
    return string(mappingsJSON), nil
}
