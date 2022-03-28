# flow-info
This is a library for interacting with the Flow blockchain's spork information.

See [examples/main.go](examples/main.go) for a usage example.

You can also use the included command to download files for a specific spork via the cli.

### CLI Usage
The following command runs the download command that downloads the `node-infos.pub.json` and `root-protocol-state-snapshot.json`
files for the `mainnet16` flow network:
```bash
go run cmd/download/main.go \
	--spork-name mainnet16 \
	--root-protocol-state-snapshot bootstrap/public-root-information/root-protocol-state-snapshot.json \
	--root-protocol-state-snapshot-sig bootstrap/public-root-information/root-protocol-state-snapshot.json.asc \
	--node-info bootstrap/public-root-information/node-infos.pub.json
```

### API Usage
Load spork details for `mainnet16`. The `sporkName` can be either a specific spork name, or the network name (`mainnet`, `testnet`, or `devnet`). If the network name is provided, the current live spork is returned.

```go
spork, err := info.LoadSpork("mainnet16")
if err != nil {
	log.Fatalf("Error loading spork: %v", err)
}
```

Load node-info details from the spork config
```go
nodeInfo, err := info.LoadNodeInfos(spork.StateArtefacts.NodeInfo)
if err != nil {
	log.Fatalf("Error loading node info: %v", err)
}
```

Load node-info details from a file
```go
nodeInfo, err := info.LoadNodeInfos("./node-infos.pub.json")
if err != nil {
	log.Fatalf("Error loading node info: %v", err)
}
```