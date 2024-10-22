package user

import(
    "webapp/constants"
    "webapp/types"
    "webapp/helpers"
    "net/http"
    "html/template"
)

func LoginPage(w http.ResponseWriter, r *http.Request){

    //Set Templates
    tmpl := template.Must(template.ParseFiles(
        constants.GetTemplatePath() + "view/layout.gohtml",
        constants.GetTemplatePath() + "view/footer.gohtml",
        constants.GetTemplatePath() + "view/login.gohtml",
    ))

    tmpl.ExecuteTemplate(w, "layout.gohtml", helpers.MergeWithCommons("", []types.Document{}, true))
}