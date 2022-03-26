# flow-info
Tools for downloading flow network information

See [examples/main.go](examples/main.go) for a usage example.

You can also use the included command to download files for a specific spork via the cli.

### Usage
The following command runs the download command that downloads the `node-infos.pub.json` and `root-protocol-state-snapshot.json`
files for the `mainnet16` flow network:
```
go run cmd/download/main.go \
	--spork-name mainnet16 \
	--root-protocol-state-snapshot bootstrap/public-root-information/root-protocol-state-snapshot.json \
	--root-protocol-state-snapshot-sig bootstrap/public-root-information/root-protocol-state-snapshot.json.asc \
	--node-info bootstrap/public-root-information/node-infos.pub.json
```
