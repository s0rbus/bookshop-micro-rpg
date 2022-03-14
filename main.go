package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/s0rbus/bookshop-micro-rpg/api"

	"github.com/alecthomas/kong"
	"github.com/guptarohit/asciigraph"
)

//Build variables
var (
	version    = "undefined"
	githash    = "undefined"
	buildstamp = "undefined"
)

var cli struct {
	Verbose      bool   `help:"Print more information." default:"false"`
	NumAttempts  int    `help:"number of attempts" default:"3"`
	Pause        int    `help:"set number of milliseconds pause between attempts" default:"0"`
	FundPatience bool   `help:"have a chance at spending 2 money to refill patience by 10"`
	Monthly      bool   `help:"landlord visits every 30 days, taking 30 in money"`
	Plot         bool   `help:"use asciigraph to plot money history for each attempt"`
	PluginDir    string `help:"folder containing expansion plugins"`
	Expansion    string `help:"name of expansion to use"`
	Version      bool   `help:"Show version and build info and exit" default:"false"`
}

var eventMaps []map[int][]api.Action
var catscores map[string]int

//var eventLog []Event
var verbose bool

type Event struct {
	EventType string
	Timestamp time.Time
	Score     int
}

func versionInfo() {
	fmt.Fprintf(os.Stderr, "App version: %v\n", version)
	fmt.Fprintf(os.Stderr, "Git commit: %v\n", githash)
	fmt.Fprintf(os.Stderr, "Build time: %v\n", buildstamp)
}

func setup() {
	eventMaps = make([]map[int][]api.Action, 3)
	for i := 0; i < 3; i++ {
		eventMaps[i] = make(map[int][]api.Action)
	}
	//Customer
	eventMaps[0][1] = []api.Action{{Score: -1, Category: "PATIENCE", Description: "Wants the bathroom"}}
	eventMaps[0][2] = []api.Action{{Score: -1, Category: "MONEY", Description: "Shoplifter"}}
	eventMaps[0][3] = []api.Action{{Score: -1, Category: "TIME", Description: "Wants a book you don't have at the moment"}}
	eventMaps[0][4] = []api.Action{{Score: -1, Category: "TIME", Description: "A literal wild animal"}}
	eventMaps[0][5] = []api.Action{{Score: -1, Category: "PATIENCE", Description: "Has a complaint"}}
	eventMaps[0][6] = []api.Action{{Score: 1, Category: "MONEY", Description: "Buys a book"}}
	//Crisis
	eventMaps[1][1] = []api.Action{{Score: -1, Category: "PATIENCE", Description: "You run out of tea"}}
	eventMaps[1][2] = []api.Action{{Score: -2, Category: "TIME", Description: "The printer breaks"}}
	eventMaps[1][3] = []api.Action{{Score: -3, Category: "TIME", Description: "You can't find a book"}}
	eventMaps[1][4] = []api.Action{{Score: -3, Category: "PATIENCE", Description: "Someone is haggling"}, {Score: 1, Category: "MONEY", Description: "ditto"}}
	eventMaps[1][5] = []api.Action{{Score: -2, Category: "TIME", Description: "The phone rings"}}
	eventMaps[1][6] = []api.Action{{Score: -2, Category: "MONEY", Description: "You bought more books"}}
	//Peculiarity
	eventMaps[2][1] = []api.Action{{Score: -2, Category: "TIME", Description: "Mysterious noises"}}
	eventMaps[2][2] = []api.Action{{Score: -1, Category: "PATIENCE", Description: "A feeling of dread"}}
	eventMaps[2][3] = []api.Action{{Score: 1, Category: "PATIENCE", Description: "A long blissful silence"}}
	eventMaps[2][4] = []api.Action{{Score: -1, Category: "MONEY", Description: "Books fall off a shelf"}}
	eventMaps[2][5] = []api.Action{{Score: 1, Category: "MONEY", Description: "You find a missing book"}}
	eventMaps[2][6] = []api.Action{{Score: -3, Category: "MONEY", Description: "Unexpected bills"}}
	catscores = make(map[string]int)
	catscores["TIME"] = 10
	catscores["PATIENCE"] = 10
}

func getActions(r1, r2 int) []api.Action {
	if verbose {
		fmt.Printf("rolled %d,%d\n", r1, r2)
	}
	var a []api.Action
	switch r1 {
	case 1, 2:
		if verbose {
			fmt.Println("a customer event")
		}
		a = eventMaps[0][r2-1]
	case 3, 4:
		if verbose {
			fmt.Println("a crisis event")
		}
		a = eventMaps[1][r2-1]
	case 5, 6:
		if verbose {
			fmt.Println("a peculiarity event")
		}
		a = eventMaps[2][r2-1]
	}
	return a
}

func Run(na int, pause int, plot bool, fp bool, mon bool, exp api.ExpansionStruct) error {
	moneyData := make([][]int, na)
	llperiod := 10
	rent := 10
	if mon {
		llperiod = 30
		rent = 30
	}
	for i := range moneyData {
		moneyData[i] = make([]int, 0)
	}
	for attempt := 1; attempt <= na; attempt++ {
		fmt.Printf("Attempt %d ", attempt)
		setup()
		patienceStart := 10
		viable := true
		maxMoney := 0
		day := 1
		for viable {
			if verbose {
				fmt.Printf("Day %d starting with patience level %d.........\n", day, patienceStart)
			}
			catscores["TIME"] = 10
			catscores["PATIENCE"] = patienceStart
			open := true
			for open {
				eventRoll1 := rand.Intn(6) + 1
				eventRoll2 := rand.Intn(6) + 1
				actions := getActions(eventRoll1, eventRoll2)
				if exp.Name != nil {
					numThrows := exp.GetRequiredThrows()
					expRows := make([]int, numThrows)
					for i := 0; i < numThrows; i++ {
						expRows[i] = rand.Intn(6) + 1
					}
					expActions, err := exp.Run(day, expRows)
					if err != nil {
						fmt.Printf("Error running expansion %s: %v\n", exp.Name(), err)
					} else {
						for _, a := range expActions {
							action := api.Action{}
							err := json.Unmarshal([]byte(a), &action)
							if err != nil {
								fmt.Printf("error unmarshalling action")
							} else {
								actions = append(actions, action)
							}
						}
					}
				}
				for _, a := range actions {
					if verbose {
						fmt.Printf("Adjust %v by %d. Reason: %v\n", a.Category, a.Score, a.Description)
					}
					catscores[a.Category] += a.Score
				}
				if catscores["TIME"] <= 0 || catscores["PATIENCE"] <= 0 {
					open = false
					if catscores["PATIENCE"] <= 0 {
						patienceStart -= 1
					}
				}
			}
			if (day % llperiod) == 0 {
				if verbose {
					fmt.Printf("Day %d, landlord needs his money, you have %d\n", day, catscores["MONEY"])
				}
				if catscores["MONEY"] < rent {
					viable = false
				} else {
					catscores["MONEY"] -= rent
				}
			}
			if viable {
				day++
				if verbose {
					fmt.Printf("a new day (number %v)\n", day)
				}
				//if fp selected and shop still viable and patienceStart < 10 and have enough money, randomly decide whether to fund patience.
				//larger amount of money and smaller patience, better chance of doing it
				if fp && patienceStart < 10 && catscores["MONEY"] >= 2 && (rand.Intn(catscores["MONEY"])+1 >= catscores["PATIENCE"]) { //(catscores["MONEY"] / 2)) {
					patienceStart = 10
					catscores["MONEY"] -= 2
					if verbose {
						fmt.Printf("Spent some money on fun. Money is now %d\n", catscores["MONEY"])
					}
				}
			}
			moneyData[attempt-1] = append(moneyData[attempt-1], catscores["MONEY"])
			if catscores["MONEY"] > maxMoney {
				maxMoney = catscores["MONEY"]
			}
		}
		t := fmt.Sprintf("%d days", day)
		if day > 365 {
			t = fmt.Sprintf("%v years, %d days", day/365, day%365)
		}
		fmt.Printf("The shop is no longer viable. You survived for %v, but now the business is closing for good. Your maximum amount of money was %d and you were left with %d\n", t, maxMoney, moneyData[attempt-1][len(moneyData[attempt-1])-1])
		copy := moneyData[attempt-1]
		sfx := ""
		if len(copy) > 50 {
			copy = copy[:50]
			sfx = " (snipped)"
		}
		fmt.Printf("Attempt %d money: %v%s\n", attempt, copy, sfx)
		if plot {

			fd := make([]float64, len(copy))
			for j, v := range copy {
				fd[j] = float64(v)
			}
			graph := asciigraph.Plot(fd)
			fmt.Println(graph)
		}
		if pause > 0 {
			time.Sleep(time.Duration(pause) * time.Millisecond)
		}
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	kong.Parse(&cli,
		kong.Name("bookshop-micro-rpg"),
		kong.Description("Based on a description by Oliver Darkshire"))
	if cli.Version {
		versionInfo()
		os.Exit(0)
	}
	verbose = cli.Verbose
	var exp api.ExpansionStruct
	var err error
	/* if cli.PluginDir != "" && cli.Expansion != "" {
		exp, err = LoadPlugins(cli.PluginDir, cli.Expansion)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} */
	if cli.PluginDir != "" && cli.Expansion != "" {
		exp, err = LoadExpansion(cli.PluginDir, cli.Expansion)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	Run(cli.NumAttempts, cli.Pause, cli.Plot, cli.FundPatience, cli.Monthly, exp)
}
