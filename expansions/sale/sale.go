package main

import (
	"github.com/s0rbus/bookshop-micro-rpg/api"
	"fmt"
)

//API

var saleStarted bool
var saleDaycount int
var currentDay int
var verbose bool

type plgExpansion struct{}

func GetExpansion() (e api.Expansion, err error) {
	e = plgExpansion{}
	fmt.Printf("[plugin GetExpansion] Returning expansion: %T %v\n", e, e)
	return
}

func (plgExpansion) SetVerbose(v bool) {
	verbose = v
}

func (plgExpansion) Name() string {
	return "simple, sale expansion"
}

//Tell app how many dice throws are required by this plugin
func (plgExpansion) GetRequiredThrows() int {
	return 2
}

func (p plgExpansion) Run(day int, throws ...int) ([]api.Action, error) {
	a := []api.Action{}
	var err error
	req := p.GetRequiredThrows()
	if len(throws) != req {
		return a, fmt.Errorf("you must provide %d dice throws", req)
	}
	status := ""
	if saleStarted && saleDaycount > 0 {
		status = "running"
	}
	//1 in 3 chance that sale starts if not already started
	if throws[0] < 3 {
		if !saleStarted {
			saleStarted = true
			saleDaycount = 6
			currentDay = day
			status = "starting"
			a = append(a, api.Action{Category: "MONEY", Score: -5, Description: "Bought more books for sale"},
				api.Action{Category: "PATIENCE", Score: -1, Description: "Start book sale"})
		}
	}
	if saleStarted {
		if day > currentDay { // a new day, decrement counter
			currentDay = day
			saleDaycount--
			if saleDaycount <= 0 {
				saleStarted = false
				a = append(a, api.Action{Category: "PATIENCE", Score: 1, Description: "End book sale"})
			}
		}
		//50/50 chance of selling a book during sale
		if (throws[1] % 2) == 0 {
			a = append(a, api.Action{Category: "MONEY", Score: 1, Description: "Sold a book in the sale"})
		}
	}
	if verbose {
		fmt.Printf("you threw %d,%d. sale is %v", throws[0], throws[1], status)

		if saleStarted {
			fmt.Printf(", on day %v\n", 15-saleDaycount)
		} else {
			fmt.Println()
		}
	}
	return a, err
}
