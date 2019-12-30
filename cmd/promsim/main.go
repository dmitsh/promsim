package main

import (
	"fmt"
	"os"

	"github.com/dmitsh/promsim/pkg/target"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		prefix, step string
		port, sets   int
	)

	tgt := kingpin.Command("target", "Start metrics generating target")
	tgt.Flag("port", "scraping port").Short('p').Default("80").IntVar(&port)
	tgt.Flag("prefix", "metrics prefix").Short('m').StringVar(&prefix)
	tgt.Flag("sets", "number of time series sets").Short('n').Default("1").IntVar(&sets)
	tgt.Flag("rate", "time interval between two metric updates").Short('r').Default("1s").StringVar(&step)

	switch kingpin.Parse() {

	case "target":
		if err := target.StartTarget(fmt.Sprintf(":%d", port), prefix, step, sets); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}
