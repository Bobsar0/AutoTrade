package webClient

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"github.com/bobsar0/AutoTrade/model"
	"github.com/pkg/errors"
)

// appTemplate is a user login-aware wrapper for a html/template.
type appTemplate struct {
	t *template.Template
}

// NewAppTemplate applies a given file to the body of the base template
// This produces a new version of base template to be rendered to the UI when the filename path is requested
func NewAppTemplate(filename string) *appTemplate {
	baseTmpl := template.Must(template.ParseFiles("webClient/templates/base.gohtml"))

	// Put the named file into a template called "body"
	path := filepath.Join("webClient/templates", filename)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("could not read template: %v", err))
	}
	template.Must(baseTmpl.New("body").Parse(string(b)))

	return &appTemplate{baseTmpl.Lookup("base.gohtml")}
}

// Execute writes the template using the provided data, adding login and user
// information to the base template.
func (tmpl *appTemplate) Execute(w http.ResponseWriter, r *http.Request, data interface{}) error {
	d := struct {
		Data        interface{}
		AuthEnabled bool
		Profile     *model.User
		LoginURL    string
		LogoutURL   string
	}{
		Data:        data,
		//AuthEnabled: app.OAuthConfig != nil,
		LoginURL:    "/login?redirect=" + r.URL.RequestURI(),
		LogoutURL:   "/logout?redirect=" + r.URL.RequestURI(),
	}

	if d.AuthEnabled {
		// Ignore any errors.
		//d.Profile = profileFromSession(r)
	}
	if err := tmpl.t.Execute(w, d); err != nil {
		return errors.Wrapf(err, "could not write template: %+v", err)
	}
	return nil
}
