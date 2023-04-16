package snapshots

import (
	"github.com/peterargue/flow-info/pkg/identities"
)

// Snapshot represents a flow protocol state snapshot
// TODO: use the flow-go's implementation
type Snapshot struct {
	Head              Header
	Identities        identities.IdentityList
	LatestSeal        map[string]interface{}
	LatestResult      map[string]interface{}
	SealingSegment    map[string]interface{}
	QuorumCertificate QuorumCertificate
	Phase             uint64
	Epochs            struct {
		Previous Epoch
		Current  Epoch
		Next     Epoch
	}
	Params Params
}

type Header struct {
	ChainID            string
	ParentID           string
	Height             uint64
	PayloadHash        string
	Timestamp          string
	View               uint64
	ParentView         uint64
	ParentVoterIndices interface{}
	ParentVoterSigData interface{}
	ProposerID         string
	ProposerSigData    interface{}
	LastViewTC         interface{}
	ID                 string
}

type QuorumCertificate struct {
	View          int
	BlockID       string
	SignerIndices string
	SigData       string
}

type Epoch struct {
	Counter            uint64
	FirstView          uint64
	DKGPhase1FinalView uint64
	DKGPhase2FinalView uint64
	DKGPhase3FinalView uint64
	FinalView          uint64
	RandomSource       string
	InitialIdentities  identities.IdentityList
	Clustering         []identities.IdentityList
	Clusters           []struct {
		Index     uint64
		Counter   uint64
		Members   identities.IdentityList
		RootBlock struct {
			Header  Header
			Payload struct{}
		}
		RootQC QuorumCertificate
	}
	DKG struct {
		GroupKey     string
		Participants map[string]struct {
			Index    uint64
			KeyShare string
		}
	}
}

type Params struct {
	ChainID                    string
	SporkID                    string
	SporkRootBlockHeight       uint64
	ProtocolVersion            uint64
	EpochCommitSafetyThreshold uint64
}
