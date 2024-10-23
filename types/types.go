
package types

import(
    "strings"
)

type Document struct {
    ID    string `json:"id"`
    Title    string  `json:"title"`
    Type string       `json:"type"`
    Content  string   `json:"content"`
    Published bool    `json:"published"`
    Category int      `json:"category"`
    Author Author      `json:"author"`
    AuthorObj Author 
    Createdat int64     `json:"createdat"`
    Updatedat int64   `json:"updatedat"`
    Slug string       `json:"slug"`
    Tags     []string `json:"tags"`
    TimeDate string `json:"timedate"`

}

type Category struct {
    Link string
    Label string
}

type Author struct {
    Link string
    Label string
    Id int
    Path string
}

type Page struct {
    HREF  string
    Title string
}

type Site struct {
    Title string
    Data  []Document
    ImageBaseUrl string
    Categories []Category
    Authors  []Author
    Post Document
    Pages []Page
    IsContact bool
}

type ContactForm struct {
    Name    string
    Email   string
    Message string
}

type Operation func(string) *strings.Reader

// ErrorResponse struct for sending error messages
type ErrorResponse struct {
    Error string `json:"error"`
}

// Response struct for JSON responses
type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"` // Optional field for additional data
}

// Struct for JSON data in POST requests
type LoginData struct {
    Username  string `json:"username"`
    Password string `json:"password"`
}

