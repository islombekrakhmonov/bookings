package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/islombekrakhmonov/bookings/pkg/config"
	"github.com/islombekrakhmonov/bookings/pkg/model"
)
	

var app *config.AppConfig

func NewTemplateCache(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *model.TemplateData) *model.TemplateData{
	return td
}


func RenderTemplate(w http.ResponseWriter, tmpl string, td *model.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache{
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil{
		log.Println(err)
	}

	//render the template

	_, err = buf.WriteTo(w)
	if err != nil{
		log.Println(err)
	}
	
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files name *page.tmpl from ./templates 

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil{
		return myCache, err
	}

	// range through all files ending with *.page.tmpl

	for _, page := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil{
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil{
			return myCache, err
		}

		if len(matches) >0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil{
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template in our cache

// 	_, inMap := tc[t] 
// 	if !inMap {
// 		// need to create the template 
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil{
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil{
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t), 
// 		"./templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil{
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil
// }


// parsedTemplate, err := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	// if err != nil {
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	// err = parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error executing template:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// }