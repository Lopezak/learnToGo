package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	a := hotel{
		Name:    "Marriot",
		Address: "5528 Mariot Way",
		City:    "Las Vegas",
		Zip:     "02345",
		Region:  "Southern",
	}

	b := hotel{
		Name:    "Hilton",
		Address: "5528 Hilton Way",
		City:    "Rancho",
		Zip:     "92336",
		Region:  "Northern",
	}

	c := hotel{
		Name:    "La Mesa",
		Address: "5528 Bluff Drive",
		City:    "Utah",
		Zip:     "84087",
		Region:  "Central",
	}

	Hotels := []hotel{a, b, c}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
