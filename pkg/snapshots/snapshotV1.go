package snapshots

import (
	"github.com/peterargue/flow-info/pkg/identities"
)

// SnapshotV1 represents a flow protocol state snapshot from before v0.33 and earlier
type SnapshotV1 struct {
	Head              Header
	Identities        identities.IdentityList
	LatestSeal        map[string]interface{}
	LatestResult      map[string]interface{}
	SealingSegment    map[string]interface{}
	QuorumCertificate QuorumCertificate
	Phase             uint64
	Epochs            struct {
		Previous EpochV1
		Current  EpochV1
		Next     EpochV1
	}
	Params Params
}

type EpochV1 struct {
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
