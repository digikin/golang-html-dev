package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	One, Two, Three string
}

type timea struct {
	Serving string
	Items   []item
}

type time []timea

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	x := time{
		timea{
			Serving: "Breakfast",
			Items: []item{
				item{
					One:   "Eggs",
					Two:   "toast",
					Three: "cerial",
				},
			},
		},
		timea{
			Serving: "Lunch",
			Items: []item{
				item{
					One:   "Sandwich",
					Two:   "Wrap",
					Three: "hotdogs",
				},
			},
		},
		timea{
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
