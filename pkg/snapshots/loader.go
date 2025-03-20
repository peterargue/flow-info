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

// LoadLatestFromAN loads the latest snapshot form an access node
func LoadLatestFromAN(ctx context.Context, accessClient *grpc.BaseClient) (*Snapshot, error) {
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

// LoadByHeightFromAN loads the latest snapshot form an access node.
func LoadByHeightFromAN(ctx context.Context, accessClient *grpc.BaseClient, height uint64) (*Snapshot, error) {
	data, err := accessClient.GetProtocolStateSnapshotByHeight(ctx, height)
	if err != nil {
		return nil, fmt.Errorf("error downloading snapshot by height: %w", err)
	}

	var snapshot Snapshot
	err = json.Unmarshal(data, &snapshot)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling snapshot json: %w", err)
	}

	return &snapshot, nil
}

// DownloadLatestFromAN downloads the latest snapshot form an access node and save it to the specified path.
func DownloadLatestFromAN(ctx context.Context, accessClient *grpc.BaseClient, path string) error {
	data, err := accessClient.GetLatestProtocolStateSnapshot(ctx)
	if err != nil {
		return fmt.Errorf("error downloading latest snapshot: %w", err)
	}

	return internal.WriteFile(path, data)
}

// DownloadByHeightFromAN downloads a snapshot by height from an access node and save it to the specified path.
// Note: this uses a vanilla grpc client instead of the flow-go-sdk grpc client.
func DownloadByHeightFromAN(
	ctx context.Context,
	accessClient *grpc.BaseClient,
	height uint64,
	path string,
) error {
	data, err := accessClient.GetProtocolStateSnapshotByHeight(ctx, height)
	if err != nil {
		return fmt.Errorf("error downloading snapshot by height: %w", err)
	}

	return internal.WriteFile(path, data)
}
