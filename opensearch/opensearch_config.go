package opensearch

import(
	"webapp/constants"
	"log"
	
	"github.com/opensearch-project/opensearch-go"
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