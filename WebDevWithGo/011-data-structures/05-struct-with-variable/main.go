package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	//sages := []string{"Ghandi", "MLK", "Buddha", "Jesus", "Muhammad"}

	// sages := map[string]string{
	// 	"India":    "Ghandi",
	// 	"America":  "MLK",
	// 	"Meditate": "Buddha",
	// 	"Love":     "Jesus",
	// 	"Prophet":  "Muhammad",
	// }

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
