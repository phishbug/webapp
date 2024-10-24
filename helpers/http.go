package helpers

import(
	"encoding/json"
	"fmt"
    "io/ioutil"
    "net/http"
)

// Function to perform a POST request
func postCall() {

    url := "https://jsonplaceholder.typicode.com/posts" // Example URL
    payload := map[string]string{"title": "foo", "body": "bar", "userId": "1"}
    
    // Convert payload to JSON
    jsonData, err := json.Marshal(payload)
    if err != nil {
        fmt.Printf("Error marshaling JSON: %s\n", err)
        return
    }

    response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Printf("Error making POST request: %s\n", err)
        return
    }
    defer response.Body.Close()

    // Read and print the response
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("Error reading response body: %s\n", err)
        return
    }
    fmt.Printf("POST Response:\n%s\n", body)
}

// Function to perform a GET request
func getCall(w http.ResponseWriter, r *http.Request, url string) {

    url := fmt.Sprintf("%s%s", constants.GetENVKey("OPEN_SEARCH_ADDRESS"), "/posts/1")

    response, err := http.Get(url)

    if err != nil {
        fmt.Printf("Error making GET request: %s\n", err)
        return
    }
    defer response.Body.Close()

    // Read and print the response
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("Error reading response body: %s\n", err)
        return
    }


    fmt.Printf("GET Response:\n%s\n", body)

    // Set the content type to application/json
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    // Encode the response to JSON and send it
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}