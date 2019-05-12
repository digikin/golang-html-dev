package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	One, Two, Three string
}

type time struct {
	Serving string
	Items   []item
}

type Time []time

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	x := Time{
		time{
			Serving: "Breakfast",
			Items: []item{
				item{
					One:   "Eggs",
					Two:   "toast",
					Three: "cerial",
				},
			},
		},
		time{
			Serving: "Lunch",
			Items: []item{
				item{
					One:   "Sandwich",
					Two:   "Wrap",
					Three: "hotdogs",
				},
			},
		},
		time{
			Serving: "Dinner",
			Items: []item{
				item{
					One:   "Steak",
					Two:   "Pork",
					Three: "Turkey",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, x)
	if err != nil {
		log.Fatalln(err)
	}
}
