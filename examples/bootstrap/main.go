package main

import (
	"fmt"
	"log"
	"os"

	"github.com/peterargue/flow-info/pkg/info"
	"github.com/peterargue/flow-info/pkg/sporks"
)

// This example downloads the snapshot data needed to bootstrap an observer node.

const (
	bootstrapDir = "./bootstrap"
	sporkName    = "mainnet16"
)

func main() {
	// load the latest spork info
	sporksInfo, err := sporks.Load()
	if err != nil {
		log.Fatalf("Error loading sporks: %v", err)
	}

	// load info for the specific spork network
	spork, err := sporksInfo.Spork(sporkName)
	if err != nil {
		log.Fatalf("Error loading spork %s: %v", sporkName, err)
	}

	dir := fmt.Sprintf("%s/public-root-information", bootstrapDir)
	err = os.MkdirAll(dir, 0755)

	err = info.Save(
		spork.StateArtefacts.RootProtocolStateSnapshot,
		fmt.Sprintf("%s/root-protocol-state-snapshot.json", dir),
	)

	if err != nil {
		log.Fatalf("Error downloading root-protocol-state-snapshot: %v", err)
	}
}
