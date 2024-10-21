package helpers

import (
	"net/http"
	"html/template"
	"webapp/constants"

)

// Custom 404 handler
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)

    tmpl := template.Must(template.ParseFiles(constants.TemplatePath + "pages/notfound.gohtml"))

    // fmt.Println(path);
    tmpl.ExecuteTemplate(w, "notfound.gohtml", nil)
    return
}

func PageNotFound(path string, w http.ResponseWriter) {
	 w.WriteHeader(http.StatusNotFound)

    tmpl := template.Must(template.ParseFiles(constants.TemplatePath + "pages/notfound.gohtml"))

    // fmt.Println(path);
    tmpl.ExecuteTemplate(w, "notfound.gohtml", nil)
    return
}
