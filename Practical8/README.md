# Practical 8 — Writing and Deploying Chaincode on Hyperledger Fabric

## Aim (as titled in the journal)
Write and deploy chaincode (a smart contract) on a Hyperledger Fabric network.

## ⚠️ Note on the original journal content
The journal's Practical 8 pages actually re-paste the **same Solidity `DataMarketplace` contract, Truffle config, and Ganache steps from Practical 7** — not real Hyperledger Fabric chaincode (which is written in Go/JavaScript/Java and deployed with `peer chaincode` / `network.sh deployCC`, as shown in Practical 3). This looks like a copy-paste mix-up in the original journal between the Ganache/Truffle practicals and the Fabric practicals.

Two things are included here so you're covered either way:

### Option A — submit exactly what's in your journal
Use the identical files from [`../Practical-07-Data-Marketplace`](../Practical-07-Data-Marketplace) (`DataMarketplace.sol`, migration, `truffle-config.js`, console commands). Safest if your journal is graded as originally submitted/signed.

### Option B — real Hyperledger Fabric chaincode (matches the actual practical title)
[`chaincode-go/`](chaincode-go) contains a proper Fabric chaincode implementation of the same "data marketplace" idea, written in Go using `fabric-contract-api-go` (the same library `asset-transfer-basic` from Practical 3 uses):

- `smartcontract.go` — `Dataset` asset with `InitLedger`, `CreateDataset`, `ReadDataset`, `PurchaseDataset`, `GetAllDatasets`
- `main.go` — chaincode entrypoint
- `go.mod` — module definition

#### Deploying it
1. Copy the `chaincode-go` folder into your `fabric-samples` directory (e.g. next to `asset-transfer-basic`)
2. From `fabric-samples/test-network`:
```bash
./network.sh up
./network.sh createChannel
./network.sh deployCC -ccn datamarketplace -ccp ../chaincode-go -ccl go
```
3. Set the Org1 environment variables (see [Practical 3's README](../Practical-03-Hyperledger-Fabric-Network/README.md) for the complete set)
4. Initialize and query:
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com \
  --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" \
  -C mychannel -n datamarketplace \
  --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" \
  -c '{"function":"InitLedger","Args":[]}'

peer chaincode query -C mychannel -n datamarketplace -c '{"Args":["GetAllDatasets"]}'
```

> ⚠️ This Go code follows the standard `fabric-contract-api-go` pattern used throughout the official Fabric samples, but there's no Go toolchain in this environment to compile it — review it (or run `go build ./...`) before your exam if you plan to demo it live.
