package helpers

import (
    "strings"
	"webapp/types"
    "webapp/constants"
)
var author = types.Author {
    Path: "anil-gaikwad.jpg",
    Label: "Anil Gaikwad",
    Link: "anil-gaikwad",
}

var categories = []types.Category{
    { Link: "engineering", Label: "Engineering"},
    { Link: "cyber-security", Label: "Cyber Security"},
    { Link: "code-bugs", Label: "Code Bugs"},
    { Link: "coding-trends", Label: "Coding Trends"},
    
}

//Common variables
func MergeWithCommons(page string, documents []types.Document, captcha bool) types.Site {
    
    if page == "contact-us" {
        return types.Site{Title: "Phish Bug", Categories: categories, Pages: makePagesLinks(), IsContact: true}
    }

    if page == "page" {
        return types.Site{Title: "Phish Bug", Categories: categories, Pages: makePagesLinks(), Post: documents[0], IsContact: captcha}    
    }
	
    return types.Site{Title: "Phish Bug", Categories: categories, Pages: makePagesLinks(), Data: documents, IsContact: captcha}
}

//Make pages like privacy, contact, etc
func makePagesLinks() []types.Page {
    // Create a slice to hold Page structs
    var pages []types.Page

    // Using range to split the string and append to the slice
    for _, value := range strings.Split(constants.Pages, "|") {
        result := types.Page{HREF: value, Title: strings.ReplaceAll(value, "-", " ")}
        pages = append(pages, result) // Correctly use append function
    }
    return pages
}