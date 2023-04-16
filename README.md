# flow-info
This is a library for programmatically accessing the Flow blockchain's spork information.

See [examples/main.go](examples/main.go) for a usage example.

You can also use the included command to download files for a specific spork via the cli.

## CLI Usage
The following command runs the download command that downloads the `node-infos.pub.json` and `root-protocol-state-snapshot.json`
files for the `mainnet16` flow network:
```bash
go run cmd/download/main.go \
	--spork-name mainnet16 \
	--root-protocol-state-snapshot bootstrap/public-root-information/root-protocol-state-snapshot.json \
	--root-protocol-state-snapshot-sig bootstrap/public-root-information/root-protocol-state-snapshot.json.asc \
	--node-info bootstrap/public-root-information/node-infos.pub.json
```

## API Usage
Load spork details for `mainnet16`. The `sporkName` can be either a specific spork name, or the network name (`mainnet`, `testnet`, or `devnet`). If the network name is provided, the current live spork is returned.

```go
// load the latest spork list
info, err := sporks.Load()
if err != nil {
	log.Fatalf("Error loading sporks: %v", err)
}

// load info for the specific spork network
spork, err := info.Spork("mainnet16")
if err != nil {
	log.Fatalf("Error loading spork %s: %v", sporkName, err)
}

// load node initial identities
nodeInfo, err := spork.Identities()
if err != nil {
	log.Fatalf("Error loading spork identities: %v", err)
}

// load the root protocol state snapshot
snapshot, err := spork.ProtocolStateSnapshot()
if err != nil {
	log.Fatalf("Error loading spork protocol state snapshot: %v", err)
}
```

Load node-info details from the spork config
```go
nodeInfo, err := identities.LoadNodeInfo(spork.StateArtefacts.NodeInfo)
if err != nil {
	log.Fatalf("Error loading node info: %v", err)
}
```

Load node-info details from a file
```go
nodeInfo, err := identities.LoadNodeInfo("./node-infos.pub.json")
if err != nil {
	log.Fatalf("Error loading node info: %v", err)
}
```

## Examples
* [examples/bootstrap/main.go](examples/bootstrap/main.go): Bootstrap an observer node.
* [examples/current_identities/main.go](examples/current_identities/main.go): Get the staked nodes for the current epoch from an Access node