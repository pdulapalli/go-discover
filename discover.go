package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/hashicorp/mdns"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	required := []string{"i", "s", "d"}
	networkInterface := flag.String("i", "wlan0", "the network interface to use")
	serviceType := flag.String("s", "_workstation._tcp", "the broadcasted service type to seek")
	discoveryDurationSec := flag.Int("d", 5, "duration of discovery in seconds")

	// Validate flags
	flag.Parse()
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			fmt.Fprintf(os.Stderr, "missing required argument/flag: -%s\n", req)
			os.Exit(2) // the same exit code flag.Parse uses
		}
	}

	// Disable logging
	log.SetOutput(ioutil.Discard)

	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Println(prettyPrint(entry))
		}
	}()

	// Get interface obj
	nwInterface, err := net.InterfaceByName(*networkInterface)
	if err != nil {
		panic(err)
	}

	var queryOpts *mdns.QueryParam = &mdns.QueryParam{}
	queryOpts.Service = *serviceType
	queryOpts.Domain = "local"
	queryOpts.Timeout = time.Duration(*discoveryDurationSec) * time.Second
	queryOpts.Interface = nwInterface
	queryOpts.Entries = entriesCh
	queryOpts.WantUnicastResponse = false

	mdns.Query(queryOpts)
	close(entriesCh)
}
