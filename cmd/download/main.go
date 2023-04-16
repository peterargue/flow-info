package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/peterargue/flow-info/pkg/info"
	"github.com/peterargue/flow-info/pkg/sporks"
)

func main() {
	var sporkName,
		rootCheckpointFile,
		rootProtocolStateSnapshot,
		rootProtocolStateSnapshotSignature,
		nodeInfo string

	flag.StringVar(&sporkName, "spork-name", "", "spork name (e.g. mainnet22, testnet43, etc)")
	flag.StringVar(&rootCheckpointFile, "root-checkpoint", "", "path where rootCheckpointFile will be written")
	flag.StringVar(&rootProtocolStateSnapshot, "root-protocol-state-snapshot", "", "path where rootProtocolStateSnapshot will be written")
	flag.StringVar(&rootProtocolStateSnapshotSignature, "root-protocol-state-snapshot-sig", "", "path where rootProtocolStateSnapshotSignature will be written")
	flag.StringVar(&nodeInfo, "node-info", "", "path where nodeInfo will be written")
	flag.Parse()

	if sporkName == "" {
		fmt.Println("Missing --spork-name")
		flag.Usage()
		return
	}

	if rootCheckpointFile == "" && rootProtocolStateSnapshot == "" && rootProtocolStateSnapshotSignature == "" && nodeInfo == "" {
		fmt.Println("At least one of --root-checkpoint, --root-protocol-state-snapshot, --root-protocol-state-snapshot-sig, --node-info must be specified")
		flag.Usage()
		return
	}

	sporkInfo, err := sporks.Load()
	if err != nil {
		log.Fatalf("error loading sporks: %v", err)
	}

	spork, err := sporkInfo.Spork(sporkName)
	if err != nil {
		log.Fatalf("error loading spork: %v", err)
	}

	if rootCheckpointFile != "" {
		err = info.Save(spork.StateArtefacts.RootCheckpointFile, rootCheckpointFile)
		if err != nil {
			log.Fatalf("error downloading root-checkpoint: %v", err)
		}
		log.Printf("wrote root-checkpoint to %s", rootCheckpointFile)
	}

	if rootProtocolStateSnapshot != "" {
		err = info.Save(spork.StateArtefacts.RootProtocolStateSnapshot, rootProtocolStateSnapshot)
		if err != nil {
			log.Fatalf("error downloading root-protocol-state-snapshot: %v", err)
		}
		log.Printf("wrote root-protocol-state-snapshot to %s", rootProtocolStateSnapshot)
	}

	if rootProtocolStateSnapshotSignature != "" {
		err = info.Save(spork.StateArtefacts.RootProtocolStateSnapshotSignature, rootProtocolStateSnapshotSignature)
		if err != nil {
			log.Fatalf("error downloading root-protocol-state-snapshot-sig: %v", err)
		}
		log.Printf("wrote root-protocol-state-snapshot-sig to %s", rootProtocolStateSnapshotSignature)
	}

	if nodeInfo != "" {
		err = info.Save(spork.StateArtefacts.NodeInfo, nodeInfo)
		if err != nil {
			log.Fatalf("error downloading node-info: %v", err)
		}
		log.Printf("wrote node-info to %s", nodeInfo)
	}
}
