package opensearch

import(
	"net/http"
	"encoding/json"
	"context"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/opensearch-project/opensearch-go/opensearchapi"

)

func GetIndexes(w http.ResponseWriter, r *http.Request){
    indices := GetIndexed();

    // Set response header to application/json
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(indices)
}

//Get Posts for admin section
func GetDocs(w http.ResponseWriter, r *http.Request) {
	// Retrieve URL parameters
    vars := mux.Vars(r) // Get the URL parameters

   client, _ := getOpenSearchClient()

    // Example: Search for documents
    searchReq := opensearchapi.SearchRequest{
        Index:  []string{vars["index"]},
        Body:  getSearchQueryForMainPage(""),
    }

    fmt.Print(vars["index"])
    searchRes, err := searchReq.Do(context.Background(), client)
    
    if err != nil {
        fmt.Print(err)
    }

    // fmt.Print(index)

   // res := getOpenSourceDoc(response)

   // fmt.Print(res)
    // Set response header to application/json
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(searchRes)
}

func GetIndex(w http.ResponseWriter, r *http.Request){
    // Retrieve URL parameters
    vars := mux.Vars(r) // Get the URL parameters

    maping, setting, _ := GetIndexDetails(vars["name"]);

    // Set response header to application/json
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode([]string{maping, setting})
}