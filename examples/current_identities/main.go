package main

import (
	"context"
	"fmt"
	"log"

	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/peterargue/flow-info/pkg/snapshots"
)

// This example get the current network identities from an Access node.

const accessURL = "access.mainnet.nodes.onflow.org:9000"

func main() {
	ctx := context.Background()

	accessClient, err := client.New(accessURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(20*1024*1024)),
	)
	if err != nil {
		log.Fatalf("Error creating access node client: %v", err)
	}

	snapshot, err := snapshots.LoadLatestForAN(ctx, accessClient)
	if err != nil {
		log.Fatalf("Error loading snapshot: %v", err)
	}

	fmt.Printf("Current Identities:\n")
	for _, identity := range snapshot.Identities {
		fmt.Printf("NodeID: %s\n", identity.NodeID)
		fmt.Printf("  Address: %s\n", identity.Address)
		fmt.Printf("  Role: %s\n", identity.Role)
		fmt.Printf("  Stake: %d\n", identity.Stake)
		fmt.Printf("  NetworkPubKey: %s\n", identity.NetworkPubKey)
		fmt.Printf("  StakingPubKey: %s\n", identity.StakingPubKey)
	}
}
