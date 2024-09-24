package compositionexample

import "fmt"

type Game struct {
	Name  string
	Price string
	Details
}

type Details struct {
	Genre         string
	GeneralRating string
	Reviews       string
}

func (d Details) showDetails() {
	fmt.Println(d.Genre)
	fmt.Println(d.GeneralRating)
	fmt.Println(d.Reviews)
}

func (g Game) ShowGameDetails() {
	fmt.Println(g.Name)
	fmt.Println(g.Reviews)
	g.showDetails()
}
