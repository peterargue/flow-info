package main

import (
	"flag"
	"fmt"
	"log"

	info "github.com/peterargue/flow-info"
)

func main() {
	var sporkName,
		rootCheckpointFile,
		rootProtocolStateSnapshot,
		rootProtocolStateSnapshotSignature,
		nodeInfo string

	flag.StringVar(&sporkName, "spork-name", "", "spork name (e.g. mainnet16, testnet33, etc)")
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

	if rootCheckpointFile == "" &&
		rootProtocolStateSnapshot == "" &&
		rootProtocolStateSnapshotSignature == "" &&
		nodeInfo == "" {

		fmt.Println("At least one of --root-checkpoint, --root-protocol-state-snapshot, --root-protocol-state-snapshot-sig, --node-info must be specified")
		flag.Usage()
		return
	}

	spork, err := info.LoadSporkJSON(sporkName)
	if err != nil {
		log.Fatalf("Error loading spork: %v", err)
	}

	if rootCheckpointFile != "" {
		err = info.DownloadToFile(spork.StateArtefacts.RootCheckpointFile, rootCheckpointFile)
		if err != nil {
			log.Fatalf("Error downloading root-checkpoint: %v", err)
		}
		log.Printf("wrote root-checkpoint to %s", rootCheckpointFile)
	}

	if rootProtocolStateSnapshot != "" {
		err = info.DownloadToFile(spork.StateArtefacts.RootProtocolStateSnapshot, rootProtocolStateSnapshot)
		if err != nil {
			log.Fatalf("Error downloading root-protocol-state-snapshot: %v", err)
		}
		log.Printf("wrote root-protocol-state-snapshot to %s", rootProtocolStateSnapshot)
	}

	if rootProtocolStateSnapshotSignature != "" {
		err = info.DownloadToFile(spork.StateArtefacts.RootProtocolStateSnapshotSignature, rootProtocolStateSnapshotSignature)
		if err != nil {
			log.Fatalf("Error downloading root-protocol-state-snapshot-sig: %v", err)
		}
		log.Printf("wrote root-protocol-state-snapshot-sig to %s", rootProtocolStateSnapshotSignature)
	}

	if nodeInfo != "" {
		err = info.DownloadToFile(spork.StateArtefacts.NodeInfo, nodeInfo)
		if err != nil {
			log.Fatalf("Error downloading node-info: %v", err)
		}
		log.Printf("wrote node-info to %s", nodeInfo)
	}
}
