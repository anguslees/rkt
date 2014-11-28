package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/coreos-inc/rkt/app-container/discovery"
)

var (
	cmdDiscover = &Command{
		Name:        "discover",
		Description: "Discover the download URLs for an app",
		Summary:     "Discover the download URLs for one or more app container images",
		Usage:       "APP...",
		Run:         runDiscover,
	}
)

func init() {
	cmdDiscover.Flags.BoolVar(&transportFlags.Insecure, "insecure", false,
		"Allow insecure non-TLS downloads over http")
}

func runDiscover(args []string) (exit int) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "discover: at least one name required")
	}

	for _, name := range args {
		labels, err := appFromString(name)
		if err != nil {
			stderr("%s: %s", name, err)
			return 1
		}
		eps, err := discovery.DiscoverEndpoints(labels["name"], labels["ver"], labels["os"], labels["amd64"], transportFlags.Insecure)

		if err != nil {
			stderr("error fetching %s: %s", name, err)
			return 1
		}
		for _, list := range [][]string{eps.Sig, eps.ACI, eps.Keys} {
			if len(list) != 0 {
				fmt.Println(strings.Join(list, ","))
			}
		}
	}

	return
}