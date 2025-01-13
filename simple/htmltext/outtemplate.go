package main

import (
	"html/template"
	"os"
)

func LoadingTemplate() {

	PrintTotal("Loading Template...")
	t, err := template.ParseFiles("templates/template.html")
	if err == nil {
		t.Execute(os.Stdout, &Kayak)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, &Kayak)
}

func LoadingMultipleTemplates() {

	PrintTotal("Loading Multiple Templates...")

	allTemplates, err := template.ParseGlob("templates/*.html")

	if err == nil {
		selectedTemplated := allTemplates.Lookup("template.html")
		err = Exec(selectedTemplated)

		// for _, t := range allTemplates.Templates() {
		// 	Printfln("Template name: %v", t.Name())
		// }
		// allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		// os.Stdout.WriteString("\n")
		// allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
	}
	// else {
	// 	Printfln("Error: %v %v", err.Error(), err.Error())
	// }
	
	if err != nil {
		Printfln("Error: %v %v", err.Error())
	}

}
