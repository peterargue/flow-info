package info

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	SporksJson = "https://raw.githubusercontent.com/onflow/flow/master/sporks.json"
)

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
}

type Node struct {
	Address string
	Key     string
}

func LoadSporkJSON(sporkName string) (*Spork, error) {
	client := http.Client{
		Timeout: time.Second * 120,
	}

	req, err := http.NewRequest(http.MethodGet, SporksJson, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error getting sporks json: %w", err)
	}

	if res.Body == nil {
		return nil, fmt.Errorf("error getting sporks json: empty body")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading sporks json: %w", err)
	}

	var sporkMap map[string]interface{}
	err = json.Unmarshal(body, &sporkMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling sporks json: %w", err)
	}

	networks := sporkMap["networks"].(map[string]interface{})
	network := networks["mainnet"].(map[string]interface{})
	spork := network[sporkName].(map[string]interface{})

	s := &Spork{
		ID:                  uint64(spork["id"].(float64)),
		Live:                spork["live"].(bool),
		Name:                spork["name"].(string),
		RootParentID:        spork["rootParentId"].(string),
		RootStateCommitment: spork["rootStateCommitment"].(string),
		GitCommitHash:       spork["gitCommitHash"].(string),
	}

	s.SporkTime, err = time.Parse(time.RFC3339, spork["sporkTime"].(string))
	if err != nil {
		return nil, fmt.Errorf("error parsing spork time: %w", err)
	}

	s.RootHeight, err = strconv.ParseUint(spork["rootHeight"].(string), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing root height: %w", err)
	}

	gcp := spork["stateArtefacts"].(map[string]interface{})["gcp"].(map[string]interface{})
	s.StateArtefacts = StateArtefacts{
		RootCheckpointFile:                 gcp["rootCheckpointFile"].(string),
		RootProtocolStateSnapshot:          gcp["rootProtocolStateSnapshot"].(string),
		RootProtocolStateSnapshotSignature: gcp["rootProtocolStateSnapshotSignature"].(string),
		NodeInfo:                           gcp["nodeInfo"].(string),
		ExecutionStateBucket:               gcp["executionStateBucket"].(string),
	}

	s.Tags = map[string]string{}
	for tag, value := range spork["tags"].(map[string]interface{}) {
		s.Tags[tag] = value.(string)
	}

	for _, node := range spork["seedNodes"].([]interface{}) {
		n := node.(map[string]interface{})
		s.SeedNodes = append(s.SeedNodes, Node{
			Address: n["address"].(string),
			Key:     n["key"].(string),
		})
	}

	for _, an := range spork["accessNodes"].([]interface{}) {
		s.AccessNodes = append(s.AccessNodes, an.(string))
	}

	return s, nil
}
