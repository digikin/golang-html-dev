package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	One, Two, Three string
}

type shift struct {
	Serving string
	Items   []item
}

type time []shift

type resturant struct {
	Name string
	Menu time
}

type resturants []resturant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	x := resturants{
		resturant{
			Name: "Denny's",
			Menu: time{
				shift{
					Serving: "Breakfast",
					Items: []item{
						item{
							One:   "Eggs",
							Two:   "toast",
							Three: "cerial",
						},
					},
				},
				shift{
					Serving: "Lunch",
					Items: []item{
						item{
							One:   "Sandwich",
							Two:   "Wrap",
							Three: "hotdogs",
						},
					},
				},
				shift{
					Serving: "Dinner",
					Items: []item{
						item{
							One:   "Steak",
							Two:   "Pork",
							Three: "Turkey",
						},
					},
				},
			},
		},
		resturant{
			Name: "IHOP",
			Menu: time{
				shift{
					Serving: "Breakfast",
					Items: []item{
						item{
							One:   "Eggs",
							Two:   "toast",
							Three: "cerial",
						},
					},
				},
				shift{
					Serving: "Lunch",
					Items: []item{
						item{
							One:   "Sandwich",
							Two:   "Wrap",
							Three: "hotdogs",
						},
					},
				},
				shift{
					Serving: "Dinner",
					Items: []item{
						item{
							One:   "Steak",
							Two:   "Pork",
							Three: "Turkey",
						},
					},
				},
			},
		},
		resturant{
			Name: "Waffle House",
			Menu: time{
				shift{
					Serving: "Breakfast",
					Items: []item{
						item{
							One:   "Eggs",
							Two:   "toast",
							Three: "cerial",
						},
					},
				},
				shift{
					Serving: "Lunch",
					Items: []item{
						item{
							One:   "Sandwich",
							Two:   "Wrap",
							Three: "hotdogs",
						},
					},
				},
				shift{
					Serving: "Dinner",
					Items: []item{
						item{
							One:   "Steak",
							Two:   "Pork",
							Three: "Turkey",
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, x)
	if err != nil {
		log.Fatalln(err)
	}
}
