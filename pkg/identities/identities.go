package identities

import "strings"

type IdentityList []NodeInfo

func (l IdentityList) ByNodeID(nodeID string) *NodeInfo {
	for _, i := range l {
		if i.NodeID == nodeID {
			return &i
		}
	}
	return nil
}

func (l IdentityList) ByAddress(address string) *NodeInfo {
	for _, i := range l {
		if i.Address == address {
			return &i
		}

		// for convenience, allow searching by just the host part of an IPv4/DNS address
		parts := strings.Split(i.Address, ":")
		if len(parts) == 2 && address == parts[0] {
			return &i
		}
	}
	return nil
}

func (l IdentityList) ByNetworkPubKey(key string) *NodeInfo {
	for _, i := range l {
		if i.NetworkPubKey == key {
			return &i
		}
	}
	return nil
}

func (l IdentityList) ByRole(role string) IdentityList {
	var result IdentityList
	for _, i := range l {
		if i.Role == role {
			result = append(result, i)
		}
	}
	return result
}
