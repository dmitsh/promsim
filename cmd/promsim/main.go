package main

import (
	"fmt"
	"os"

	"github.com/dmitsh/promsim/pkg/target"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	cfg := &target.Config{}
	tgt := kingpin.Command("target", "Start metrics generating target.")
	tgt.Flag("address", "scraping endpoint address.").Short('a').Default(":8080").StringVar(&cfg.Address)
	tgt.Flag("metrics", "metrics path.").Short('m').Default("/metrics").StringVar(&cfg.MetricsPath)
	tgt.Flag("job", "job name.").Short('j').StringVar(&cfg.JobName)
	tgt.Flag("sets", "number of time series sets.").Short('n').Default("1").IntVar(&cfg.Sets)
	tgt.Flag("rate", "time interval between two metric updates.").Short('r').Default("1s").StringVar(&cfg.UpdateRate)
	tgt.Flag("tls.enabled", "enable TLS.").Default("false").BoolVar(&cfg.TlsEnabled)
	tgt.Flag("tls.key", "path to the server key.").StringVar(&cfg.TlsKeyPath)
	tgt.Flag("tls.cert", "path to the server certificate.").StringVar(&cfg.TlsCertPath)

	switch kingpin.Parse() {

	case "target":
		if err := target.StartTarget(cfg); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}
