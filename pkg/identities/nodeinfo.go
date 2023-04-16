package identities

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/peterargue/flow-info/internal"
)

type NodeInfo struct {
	Role          string
	Address       string
	NodeID        string
	Stake         uint64
	NetworkPubKey string
	StakingPubKey string
}

// LoadNodeInfo loads node infos from a file or url.
func LoadNodeInfo(url string) (IdentityList, error) {
	var data []byte
	var err error

	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		data, err = internal.Download(url)
		if err != nil {
			return nil, fmt.Errorf("error downloading node info: %w", err)
		}
	} else {
		data, err = internal.ReadFile(url)
		if err != nil {
			return nil, fmt.Errorf("error reading node info: %w", err)
		}
	}

	var nodeInfos []NodeInfo
	err = json.Unmarshal(data, &nodeInfos)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling node-infos json: %w", err)
	}

	return nodeInfos, nil
}
