package sporks

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/peterargue/flow-info/internal"
)

const (
	SporksJson = "https://raw.githubusercontent.com/onflow/flow/master/sporks.json"
)

// Load loads details about a specific spork from the official spork.json file.
func Load() (*SporkInfo, error) {
	data, err := internal.Download(SporksJson)
	if err != nil {
		return nil, fmt.Errorf("error downloading sporks json: %w", err)
	}

	var sporkInfoMap map[string]interface{}
	err = json.Unmarshal(data, &sporkInfoMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling sporks json: %w", err)
	}

	networks, ok := sporkInfoMap["networks"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error parsing sporks json: missing/invalid networks field")
	}

	si := newSporkInfo()
	for name, networkBlob := range networks {
		si.Networks[name] = *newSporks()

		network, ok := networkBlob.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("error parsing sporks json: invalid network field")
		}

		for sporkName, sporkData := range network {
			sporkMap, ok := sporkData.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("error parsing spork %s json: invalid spork field", sporkName)
			}

			spork, err := parseSpork(sporkMap)
			if err != nil {
				return nil, fmt.Errorf("error parsing spork %s: %w", sporkName, err)
			}

			si.Networks[name].Sporks[sporkName] = spork
		}
	}

	return si, nil
}

func parseSpork(spork map[string]interface{}) (Spork, error) {
	var err error
	s := Spork{
		ID:                  uint64(internal.Extract[float64](spork, "id")),
		Live:                internal.Extract[bool](spork, "live"),
		Name:                internal.Extract[string](spork, "name"),
		RootParentID:        internal.Extract[string](spork, "rootParentId"),
		RootStateCommitment: internal.Extract[string](spork, "rootStateCommitment"),
		GitCommitHash:       internal.Extract[string](spork, "gitCommitHash"),
	}

	s.SporkTime, err = time.Parse(time.RFC3339, internal.Extract[string](spork, "sporkTime"))
	if err != nil {
		return Spork{}, fmt.Errorf("error parsing spork time: %w", err)
	}

	s.RootHeight, err = strconv.ParseUint(internal.Extract[string](spork, "rootHeight"), 10, 64)
	if err != nil {
		return Spork{}, fmt.Errorf("error parsing root height: %w", err)
	}

	s.StateArtefacts, err = parseStateArtefacts(spork)
	if err != nil {
		return Spork{}, fmt.Errorf("error parsing state artefacts: %w", err)
	}

	s.Tags, err = parseTags(spork)
	if err != nil {
		return Spork{}, fmt.Errorf("error parsing tags: %w", err)
	}

	s.SeedNodes, err = parseSeedNodes(spork)
	if err != nil {
		return Spork{}, fmt.Errorf("error parsing seed nodes: %w", err)
	}

	s.AccessNodes, err = parseAccessNodes(spork)
	if err != nil {
		return Spork{}, fmt.Errorf("error parsing access nodes: %w", err)
	}

	return s, nil
}

func parseStateArtefacts(spork map[string]interface{}) (StateArtefacts, error) {
	stateArtefacts, ok := spork["stateArtefacts"].(map[string]interface{})
	if !ok {
		return StateArtefacts{}, fmt.Errorf("stateArtefacts missing/invalid")
	}

	gcp, ok := stateArtefacts["gcp"].(map[string]interface{})
	if !ok {
		return StateArtefacts{}, fmt.Errorf("gcp missing/invalid")
	}

	return StateArtefacts{
		RootCheckpointFile:                 internal.Extract[string](gcp, "rootCheckpointFile"),
		RootProtocolStateSnapshot:          internal.Extract[string](gcp, "rootProtocolStateSnapshot"),
		RootProtocolStateSnapshotSignature: internal.Extract[string](gcp, "rootProtocolStateSnapshotSignature"),
		NodeInfo:                           internal.Extract[string](gcp, "nodeInfo"),
		ExecutionStateBucket:               internal.Extract[string](gcp, "executionStateBucket"),
		ProtocolDBArchive:                  internal.Extract[string](gcp, "protocolDBArchive"),
		ProtocolDBArchiveChecksum:          internal.Extract[string](gcp, "protocolDBArchiveChecksum"),
		ExecutionStateArchive:              internal.Extract[string](gcp, "executionStateArchive"),
		ExecutionStateArchiveChecksum:      internal.Extract[string](gcp, "executionStateArchiveChecksum"),
	}, nil
}

func parseTags(spork map[string]interface{}) (map[string]string, error) {
	tags, ok := spork["tags"]
	if !ok {
		return map[string]string{}, nil
	}

	tagList, ok := tags.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("tags invalid")
	}

	parsed := map[string]string{}
	for tag, value := range tagList {
		parsed[tag] = value.(string)
	}
	return parsed, nil
}

func parseSeedNodes(spork map[string]interface{}) ([]Node, error) {
	seeds, ok := spork["seedNodes"]
	if !ok {
		return nil, nil
	}

	seedList, ok := seeds.([]interface{})
	if !ok {
		return nil, fmt.Errorf("seedNodes invalid")
	}

	seedNodes := make([]Node, 0)
	for i, node := range seedList {
		n, ok := node.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("seedNode %d invalid", i)
		}

		seedNodes = append(seedNodes, Node{
			Address: internal.Extract[string](n, "address"),
			Key:     internal.Extract[string](n, "key"),
		})
	}
	return seedNodes, nil
}

func parseAccessNodes(spork map[string]interface{}) ([]string, error) {
	ans, ok := spork["accessNodes"]
	if !ok {
		return nil, nil
	}

	anList, ok := ans.([]interface{})
	if !ok {
		return nil, fmt.Errorf("accessNodes invalid")
	}

	accessNodes := make([]string, 0)
	for _, an := range anList {
		accessNodes = append(accessNodes, an.(string))
	}
	return accessNodes, nil
}

func networkName(sporkName string) string {
	switch {
	case strings.HasPrefix(sporkName, "mainnet") || strings.HasPrefix(sporkName, "candidate"):
		return "mainnet"
	case strings.HasPrefix(sporkName, "testnet") || strings.HasPrefix(sporkName, "devnet"):
		return "testnet"
	}

	return ""
}
