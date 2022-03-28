package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	info "github.com/peterargue/flow-info"
)

const (
	bootstrapDir = "./bootstrap"
)

func main() {
	var sporkName string
	flag.StringVar(&sporkName, "spork-name", "", "spork name (e.g. mainnet16, testnet33, etc)")
	flag.Parse()

	if sporkName == "" {
		fmt.Println("Missing --spork-name")
		flag.Usage()
		return
	}

	spork, err := info.LoadSpork(sporkName)
	if err != nil {
		log.Fatalf("Error loading spork: %v", err)
	}

	dir := fmt.Sprintf("%s/public-root-information", bootstrapDir)
	err = os.MkdirAll(dir, 0755)

	err = info.DownloadToFile(
		spork.StateArtefacts.RootProtocolStateSnapshot,
		fmt.Sprintf("%s/root-protocol-state-snapshot.json", dir),
	)

	if err != nil {
		log.Fatalf("Error downloading root-protocol-state-snapshot: %v", err)
	}

	err = info.DownloadToFile(
		spork.StateArtefacts.NodeInfo,
		fmt.Sprintf("%s/node-infos.pub.json", dir),
	)

	if err != nil {
		log.Fatalf("Error downloading node-infos: %v", err)
	}
}
