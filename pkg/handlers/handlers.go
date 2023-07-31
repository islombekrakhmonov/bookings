package handlers

import (
	"net/http"

	"github.com/islombekrakhmonov/bookings/pkg/config"
	"github.com/islombekrakhmonov/bookings/pkg/model"
	"github.com/islombekrakhmonov/bookings/pkg/render"
)



var Repo *Repository

type Repository struct{
	App *config.AppConfig	
}

func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository){
	Repo = r
}

// Home is the about page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &model.TemplateData{})
}

// About is the about page handler

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"


	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")


	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &model.TemplateData{
		StringMap: stringMap,
	})


}