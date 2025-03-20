package snapshots

import (
	"github.com/peterargue/flow-info/pkg/identities"
)

// RootProtocolStateSnapshot struct
type Snapshot struct {
	SealingSegment      SealingSegment    `json:"SealingSegment"`
	QuorumCertificate   QuorumCertificate `json:"QuorumCertificate"`
	Params              Params            `json:"Params"`
	SealedVersionBeacon interface{}       `json:"SealedVersionBeacon"`
}

func (s Snapshot) CurrentEpochSetup() EpochSetup {
	state := s.SealingSegment.ProtocolStateEntry()
	return state.EpochEntry.CurrentEpochSetup
}

func (s Snapshot) NextEpochSetup() EpochSetup {
	state := s.SealingSegment.ProtocolStateEntry()
	return state.EpochEntry.NextEpochSetup
}

// SealingSegment struct
type SealingSegment struct {
	Blocks               []Block                       `json:"Blocks"`
	ExtraBlocks          []Block                       `json:"ExtraBlocks"`
	ExecutionResults     []ExecutionResult             `json:"ExecutionResults"`
	LatestSeals          map[string]string             `json:"LatestSeals"`
	FirstSeal            FirstSeal                     `json:"FirstSeal"`
	ProtocolStateEntries map[string]ProtocolStateEntry `json:"ProtocolStateEntries"`
}

func (s SealingSegment) ProtocolStateEntry() ProtocolStateEntry {
	for _, v := range s.ProtocolStateEntries {
		return v
	}
	return ProtocolStateEntry{}
}

// Block struct
type Block struct {
	Header  Header  `json:"Header"`
	Payload Payload `json:"Payload"`
}

// Header struct
type Header struct {
	ChainID            string      `json:"ChainID"`
	ParentID           string      `json:"ParentID"`
	Height             uint64      `json:"Height"`
	PayloadHash        string      `json:"PayloadHash"`
	Timestamp          string      `json:"Timestamp"`
	View               uint64      `json:"View"`
	ParentView         uint64      `json:"ParentView"`
	ParentVoterIndices string      `json:"ParentVoterIndices"`
	ParentVoterSigData string      `json:"ParentVoterSigData"`
	ProposerID         string      `json:"ProposerID"`
	ProposerSigData    string      `json:"ProposerSigData"`
	LastViewTC         interface{} `json:"LastViewTC"`
	ID                 string      `json:"ID"`
}

// Payload struct
type Payload struct {
	Guarantees      []Guarantee       `json:"Guarantees"`
	Seals           []Seal            `json:"Seals"`
	Receipts        []Receipt         `json:"Receipts"`
	Results         []ExecutionResult `json:"Results"`
	ProtocolStateID string            `json:"ProtocolStateID"`
}

type Guarantee struct {
	CollectionID     string  `json:"CollectionID"`
	ReferenceBlockID string  `json:"ReferenceBlockID"`
	ChainID          string  `json:"ChainID"`
	SignerIndices    string  `json:"SignerIndices"`
	Signature        *string `json:"Signature"`
}

type Seal struct {
	BlockID                string                  `json:"BlockID"`
	ResultID               string                  `json:"ResultID"`
	FinalState             string                  `json:"FinalState"`
	AggregatedApprovalSigs []AggregatedApprovalSig `json:"AggregatedApprovalSigs"`
	ID                     string                  `json:"ID"`
}

type AggregatedApprovalSig struct {
	VerifierSignatures []string `json:"VerifierSignatures"`
	SignerIDs          []string `json:"SignerIDs"`
}

type Receipt struct {
	ExecutorID        string   `json:"ExecutorID"`
	ResultID          string   `json:"ResultID"`
	Spocks            []string `json:"Spocks"`
	ExecutorSignature string   `json:"ExecutorSignature"`
	ID                string   `json:"ID"`
}

// Chunk struct
type Chunk struct {
	CollectionIndex      uint64 `json:"CollectionIndex"`
	StartState           string `json:"StartState"`
	EventCollection      string `json:"EventCollection"`
	BlockID              string `json:"BlockID"`
	TotalComputationUsed uint64 `json:"TotalComputationUsed"`
	NumberOfTransactions uint64 `json:"NumberOfTransactions"`
	Index                uint64 `json:"Index"`
	EndState             string `json:"EndState"`
}

// ServiceEvent struct
type ServiceEvent struct {
	Type  string      `json:"Type"`
	Event interface{} `json:"Event"`
}

// ExecutionResult struct
type ExecutionResult struct {
	PreviousResultID string         `json:"PreviousResultID"`
	BlockID          string         `json:"BlockID"`
	Chunks           []Chunk        `json:"Chunks"`
	ServiceEvents    []ServiceEvent `json:"ServiceEvents"`
	ExecutionDataID  string         `json:"ExecutionDataID"`
	ID               string         `json:"ID"`
}

// FirstSeal struct
type FirstSeal struct {
	BlockID                string                  `json:"BlockID"`
	ResultID               string                  `json:"ResultID"`
	FinalState             string                  `json:"FinalState"`
	AggregatedApprovalSigs []AggregatedApprovalSig `json:"AggregatedApprovalSigs"`
	ID                     string                  `json:"ID"`
}

// KVStore struct
type KVStore struct {
	Version uint64 `json:"Version"`
	Data    string `json:"Data"`
}

type Epoch struct {
	SetupID          string `json:"SetupID"`
	CommitID         string `json:"CommitID"`
	ActiveIdentities []struct {
		NodeID  string `json:"NodeID"`
		Ejected bool   `json:"Ejected"`
	} `json:"ActiveIdentities"`
	EpochExtensions interface{} `json:"EpochExtensions"`
}

type Identity struct {
	identities.NodeInfo
	ParticipationStatus string `json:"ParticipationStatus,omitempty"`
}

type EpochSetup struct {
	Counter            uint64     `json:"Counter"`
	FirstView          uint64     `json:"FirstView"`
	DKGPhase1FinalView uint64     `json:"DKGPhase1FinalView"`
	DKGPhase2FinalView uint64     `json:"DKGPhase2FinalView"`
	DKGPhase3FinalView uint64     `json:"DKGPhase3FinalView"`
	FinalView          uint64     `json:"FinalView"`
	Participants       []Identity `json:"Participants"`
	Assignments        [][]string `json:"Assignments"`
	RandomSource       string     `json:"RandomSource"`
	TargetDuration     uint64     `json:"TargetDuration"`
	TargetEndTime      uint64     `json:"TargetEndTime"`
}

func (e EpochSetup) Identities() identities.IdentityList {
	list := make(identities.IdentityList, len(e.Participants))
	for i, v := range e.Participants {
		list[i] = v.NodeInfo
	}
	return list
}

func (e EpochSetup) Clusters() []identities.IdentityList {
	participants := make(map[string]identities.NodeInfo, len(e.Participants))
	for _, identity := range e.Participants {
		participants[identity.NodeID] = identity.NodeInfo
	}

	clusters := make([]identities.IdentityList, len(e.Assignments))
	for i, assignment := range e.Assignments {
		cluster := make(identities.IdentityList, 0, len(assignment))
		for _, nodeID := range assignment {
			identity, ok := participants[nodeID]
			if !ok {
				continue
			}
			cluster = append(cluster, identity)
		}
		clusters[i] = cluster
	}
	return clusters
}

type EpochCommit struct {
	Counter    uint64 `json:"Counter"`
	ClusterQCs []struct {
		SigData  string   `json:"SigData"`
		VoterIDs []string `json:"VoterIDs"`
	} `json:"ClusterQCs"`
	DKGGroupKey        string   `json:"DKGGroupKey"`
	DKGParticipantKeys []string `json:"DKGParticipantKeys"`
}

// EpochEntry struct
type EpochEntry struct {
	PreviousEpoch             Epoch       `json:"PreviousEpoch"`
	CurrentEpoch              Epoch       `json:"CurrentEpoch"`
	NextEpoch                 Epoch       `json:"NextEpoch"`
	EpochFallbackTriggered    bool        `json:"EpochFallbackTriggered"`
	PreviousEpochSetup        EpochSetup  `json:"PreviousEpochSetup"`
	PreviousEpochCommit       EpochCommit `json:"PreviousEpochCommit"`
	CurrentEpochSetup         EpochSetup  `json:"CurrentEpochSetup"`
	CurrentEpochCommit        EpochCommit `json:"CurrentEpochCommit"`
	NextEpochSetup            EpochSetup  `json:"NextEpochSetup"`
	NextEpochCommit           EpochCommit `json:"NextEpochCommit"`
	CurrentEpochIdentityTable []Identity  `json:"CurrentEpochIdentityTable"`
	NextEpochIdentityTable    []Identity  `json:"NextEpochIdentityTable"`
}

func (e *EpochEntry) CurrentEpochInitialIdentities() identities.IdentityList {
	list := make(identities.IdentityList, len(e.CurrentEpochIdentityTable))
	for i, v := range e.CurrentEpochIdentityTable {
		list[i] = v.NodeInfo
	}
	return list
}

// ProtocolStateEntry struct
type ProtocolStateEntry struct {
	KVStore    KVStore    `json:"KVStore"`
	EpochEntry EpochEntry `json:"EpochEntry"`
}

// QuorumCertificate struct
type QuorumCertificate struct {
	View          uint64 `json:"View"`
	BlockID       string `json:"BlockID"`
	SignerIndices string `json:"SignerIndices"`
	SigData       string `json:"SigData"`
}

// Params struct
type Params struct {
	ChainID              string `json:"ChainID"`
	SporkID              string `json:"SporkID"`
	SporkRootBlockHeight uint64 `json:"SporkRootBlockHeight"`
	ProtocolVersion      int    `json:"ProtocolVersion"`
}
