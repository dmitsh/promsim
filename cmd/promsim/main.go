package main

import (
	"fmt"
	"os"

	"github.com/dmitsh/promsim/pkg/target"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		address, step string
		sets          int
	)

	tgt := kingpin.Command("target", "Start metrics generating target")
	tgt.Flag("address", "scraping endpoint address.").Short('a').Default(":8080").StringVar(&address)
	tgt.Flag("sets", "number of time series sets.").Short('n').Default("1").IntVar(&sets)
	tgt.Flag("rate", "time interval between two metric updates.").Short('r').Default("1s").StringVar(&step)

	switch kingpin.Parse() {

	case "target":
		if err := target.StartTarget(address, step, sets); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}
