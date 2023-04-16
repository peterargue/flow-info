package snapshots

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/peterargue/flow-info/internal"
)

// TODO: Add support for verifying a snapshot

// Load loads a snapshot from a local file or url.
func Load(url string) (*Snapshot, error) {
	var data []byte
	var err error

	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		data, err = internal.Download(url)
		if err != nil {
			return nil, fmt.Errorf("error downloading snapshot: %w", err)
		}
	} else {
		data, err = internal.ReadFile(url)
		if err != nil {
			return nil, fmt.Errorf("error reading snapshot: %w", err)
		}
	}

	var snapshot Snapshot
	err = json.Unmarshal(data, &snapshot)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling snapshot json: %w", err)
	}

	return &snapshot, nil
}

// LoadLatestForAN loads the latest snapshot form an access node.
func LoadLatestForAN(ctx context.Context, accessClient *grpc.BaseClient) (*Snapshot, error) {
	data, err := accessClient.GetLatestProtocolStateSnapshot(ctx)
	if err != nil {
		return nil, fmt.Errorf("error downloading latest snapshot: %w", err)
	}

	var snapshot Snapshot
	err = json.Unmarshal(data, &snapshot)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling snapshot json: %w", err)
	}

	return &snapshot, nil
}
