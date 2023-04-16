package sporks

import (
	"fmt"
	"strings"
)

// SporkInfo contains information about all sporks.
type SporkInfo struct {
	Networks map[string]Sporks
}

// Sporks contains a list of spork information.
type Sporks struct {
	Sporks map[string]Spork
}

func newSporkInfo() *SporkInfo {
	return &SporkInfo{
		Networks: make(map[string]Sporks),
	}
}

func newSporks() *Sporks {
	return &Sporks{
		Sporks: make(map[string]Spork),
	}
}

// Spork returns the spork with the given name.
func (info *SporkInfo) Spork(sporkName string) (*Spork, error) {
	networkName := networkName(sporkName)
	if networkName == "" {
		return nil, fmt.Errorf("invalid spork name: %s", sporkName)
	}

	network, ok := info.Networks[networkName]
	if !ok {
		return nil, fmt.Errorf("network %s not found", networkName)
	}

	spork, ok := network.Sporks[sporkName]
	if !ok {
		return nil, fmt.Errorf("spork %s not found", sporkName)
	}

	return &spork, nil
}

// LatestSpork returns the most recent spork for the given network.
func (info *SporkInfo) LatestSpork(network string) (*Spork, error) {
	sporks, ok := info.Networks[network]
	if !ok {
		return nil, fmt.Errorf("network %s not found", network)
	}

	var latestSpork *Spork
	for _, spork := range sporks.Sporks {
		// if there is a live spork, return that
		if spork.Live {
			return &spork, nil
		}

		// otherwise, return find the most recent spork
		if latestSpork == nil || spork.SporkTime.After(latestSpork.SporkTime) {
			latestSpork = &spork
		}
	}

	return latestSpork, nil
}

// Print prints the spork info to stdout.
func (info *SporkInfo) Print() {
	for networkName, network := range info.Networks {
		fmt.Printf("\n%s:\n", networkName)
		fmt.Println(strings.Repeat("=", len(networkName)+1))
		for _, spork := range network.Sporks {
			spork.Print()
		}
	}
}
