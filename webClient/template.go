
package webClient

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"github.com/chidi150c/autotrade/model"
	"github.com/pkg/errors"
)


// appTemplate is a user login-aware wrapper for a html/template.
type appTemplate struct {
	t *template.Template
}


// parseTemplate applies a given file to the body of the base template.
func NewAppTemplate(filename string) *appTemplate {
	tmpl := template.Must(template.ParseFiles("webClient/templates/base.html"))

	// Put the named file into a template called "body"
	path := filepath.Join("webClient/templates", filename)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("could not read template: %v", err))
	}
	template.Must(tmpl.New("body").Parse(string(b)))

	return &appTemplate{tmpl.Lookup("base.html")}
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
		//AuthEnabled: bookshelf.OAuthConfig != nil,
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
