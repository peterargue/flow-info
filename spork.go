package info

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	SporksJson = "https://raw.githubusercontent.com/onflow/flow/master/sporks.json"
)

var ErrNetworkNotFound = fmt.Errorf("network not found")
var ErrSporkNotFound = fmt.Errorf("spork not found")

type Spork struct {
	ID                  uint64
	Live                bool
	Name                string
	SporkTime           time.Time
	RootHeight          uint64
	RootParentID        string
	RootStateCommitment string
	GitCommitHash       string
	StateArtefacts      StateArtefacts
	Tags                map[string]string
	SeedNodes           []Node
	AccessNodes         []string
}

type StateArtefacts struct {
	RootCheckpointFile                 string
	RootProtocolStateSnapshot          string
	RootProtocolStateSnapshotSignature string
	NodeInfo                           string
	ExecutionStateBucket               string
	ProtocolDBArchive                  string
	ProtocolDBArchiveChecksum          string
	ExecutionStateArchive              string
	ExecutionStateArchiveChecksum      string
}

type Node struct {
	Address string
	Key     string
}

// LoadSpork loads details about a specific spork from the official spork.json file.
func LoadSpork(sporkName string) (*Spork, error) {
	data, err := Download(SporksJson)
	if err != nil {
		return nil, fmt.Errorf("error downloading sporks json: %w", err)
	}

	var sporkMap map[string]interface{}
	err = json.Unmarshal(data, &sporkMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling sporks json: %w", err)
	}

	spork, err := spork(sporkName, sporkMap)
	if err != nil {
		return nil, err
	}

	return parse(spork)
}

func parse(spork map[string]interface{}) (*Spork, error) {
	var err error
	s := &Spork{
		ID:                  uint64(spork["id"].(float64)),
		Live:                boolValue(spork, "live"),
		Name:                stringValue(spork, "name"),
		RootParentID:        stringValue(spork, "rootParentId"),
		RootStateCommitment: stringValue(spork, "rootStateCommitment"),
		GitCommitHash:       stringValue(spork, "gitCommitHash"),
	}

	s.SporkTime, err = time.Parse(time.RFC3339, stringValue(spork, "sporkTime"))
	if err != nil {
		return nil, fmt.Errorf("error parsing spork time: %w", err)
	}

	s.RootHeight, err = strconv.ParseUint(stringValue(spork, "rootHeight"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing root height: %w", err)
	}

	gcp := spork["stateArtefacts"].(map[string]interface{})["gcp"].(map[string]interface{})
	s.StateArtefacts = StateArtefacts{
		RootCheckpointFile:                 stringValue(gcp, "rootCheckpointFile"),
		RootProtocolStateSnapshot:          stringValue(gcp, "rootProtocolStateSnapshot"),
		RootProtocolStateSnapshotSignature: stringValue(gcp, "rootProtocolStateSnapshotSignature"),
		NodeInfo:                           stringValue(gcp, "nodeInfo"),
		ExecutionStateBucket:               stringValue(gcp, "executionStateBucket"),
		ProtocolDBArchive:                  stringValue(gcp, "protocolDBArchive"),
		ProtocolDBArchiveChecksum:          stringValue(gcp, "protocolDBArchiveChecksum"),
		ExecutionStateArchive:              stringValue(gcp, "executionStateArchive"),
		ExecutionStateArchiveChecksum:      stringValue(gcp, "executionStateArchiveChecksum"),
	}

	s.Tags = map[string]string{}
	for tag, value := range spork["tags"].(map[string]interface{}) {
		s.Tags[tag] = value.(string)
	}

	if seeds, has := spork["seedNodes"]; has {
		for _, node := range seeds.([]interface{}) {
			n := node.(map[string]interface{})
			s.SeedNodes = append(s.SeedNodes, Node{
				Address: stringValue(n, "address"),
				Key:     stringValue(n, "key"),
			})
		}
	}

	if ans, has := spork["accessNodes"]; has {
		for _, an := range ans.([]interface{}) {
			s.AccessNodes = append(s.AccessNodes, an.(string))
		}
	}

	return s, nil
}

func spork(sporkName string, sporkMap map[string]interface{}) (map[string]interface{}, error) {
	networkName := networkName(sporkName)
	if networkName == "" {
		return nil, ErrNetworkNotFound
	}

	networks := sporkMap["networks"].(map[string]interface{})
	network := networks[networkName].(map[string]interface{})

	// check if sporkName is an exact match
	spork, found := network[sporkName]
	if found {
		return spork.(map[string]interface{}), nil
	}

	// otherwise, it must be one of the known network names
	if sporkName != "mainnet" && sporkName != "testnet" && sporkName != "devnet" {
		return nil, ErrSporkNotFound
	}

	// return the live spork for the network
	for _, data := range network {
		spork := data.(map[string]interface{})
		if boolValue(spork, "live") {
			return spork, nil
		}
	}

	return nil, ErrSporkNotFound
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

func stringValue(container map[string]interface{}, name string) string {
	if value, found := container[name]; found {
		return value.(string)
	}
	return ""
}

func boolValue(container map[string]interface{}, name string) bool {
	if value, found := container[name]; found {
		return value.(bool)
	}
	return false
}
