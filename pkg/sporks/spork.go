package sporks

import (
	"fmt"
	"time"

	"github.com/peterargue/flow-info/pkg/identities"
	"github.com/peterargue/flow-info/pkg/snapshots"
)

// Spork contains information about a spork.
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

// StateArtefacts contains information about the state artefacts for a spork.
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

// Node contains information about a seed node.
type Node struct {
	Address string
	Key     string
}

// Identities returns the initial identities for the spork.
func (s *Spork) Identities() (identities.IdentityList, error) {
	return identities.LoadNodeInfo(s.StateArtefacts.NodeInfo)
}

// ProtocolStateSnapshot returns the protocol state snapshot for the spork.
func (s *Spork) ProtocolStateSnapshot() (*snapshots.Snapshot, error) {
	return snapshots.Load(s.StateArtefacts.RootProtocolStateSnapshot)
}

// Print prints the spork details.
func (s *Spork) Print() {
	fmt.Printf("%s:\n", s.Name)
	fmt.Printf("  ID: %v\n", s.ID)
	fmt.Printf("  Live: %v\n", s.Live)
	fmt.Printf("  Name: %v\n", s.Name)
	fmt.Printf("  SporkTime: %v\n", s.SporkTime)
	fmt.Printf("  RootHeight: %v\n", s.RootHeight)
	fmt.Printf("  RootParentID: %v\n", s.RootParentID)
	fmt.Printf("  RootStateCommitment: %v\n", s.RootStateCommitment)
	fmt.Printf("  GitCommitHash: %v\n", s.GitCommitHash)
	fmt.Printf("  StateArtefacts:\n")
	fmt.Printf("    RootCheckpointFile: %v\n", s.StateArtefacts.RootCheckpointFile)
	fmt.Printf("    RootProtocolStateSnapshot: %v\n", s.StateArtefacts.RootProtocolStateSnapshot)
	fmt.Printf("    RootProtocolStateSnapshotSignature: %v\n", s.StateArtefacts.RootProtocolStateSnapshotSignature)
	fmt.Printf("    NodeInfo: %v\n", s.StateArtefacts.NodeInfo)
	fmt.Printf("    ExecutionStateBucket: %v\n", s.StateArtefacts.ExecutionStateBucket)
	fmt.Printf("    ProtocolDBArchive: %v\n", s.StateArtefacts.ProtocolDBArchive)
	fmt.Printf("    ProtocolDBArchiveChecksum: %v\n", s.StateArtefacts.ProtocolDBArchiveChecksum)
	fmt.Printf("    ExecutionStateArchive: %v\n", s.StateArtefacts.ExecutionStateArchive)
	fmt.Printf("    ExecutionStateArchiveChecksum: %v\n", s.StateArtefacts.ExecutionStateArchiveChecksum)
	if len(s.Tags) > 0 {
		fmt.Printf("  Tags:\n")
		for k, v := range s.Tags {
			fmt.Printf("    %s: %s\n", k, v)
		}
	}
	if len(s.SeedNodes) > 0 {
		fmt.Printf("  SeedNodes:\n")
		for _, node := range s.SeedNodes {
			fmt.Printf("  - Address: %s\n", node.Address)
			fmt.Printf("    Key: %s\n", node.Key)
		}
	}
	if len(s.AccessNodes) > 0 {
		fmt.Printf("  AccessNodes:\n")
		for _, node := range s.AccessNodes {
			fmt.Printf("    %s\n", node)
		}
	}
}
