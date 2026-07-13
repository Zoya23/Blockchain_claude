# Blockchain Practical Exam — Cheat Sheet

## 1. Environment / Tool Verification
```bash
node -v
npm -v
truffle version
docker --version
ipfs --version
```

## 2. Ganache
- Open Ganache → **Quickstart Ethereum**
- RPC Server: `http://127.0.0.1:7545`
- Network ID: `5777` (sometimes shown as `1337`)
- Keep it running for every Truffle practical

## 3. Truffle — full workflow
```bash
truffle init                 # scaffold contracts/ migrations/ test/ truffle-config.js
truffle compile              # compile .sol files
truffle compile --all        # force recompile everything
truffle migrate               # deploy to network in truffle-config.js
truffle migrate --reset      # redeploy from scratch
truffle console              # interactive JS console connected to the network
```

### truffle-config.js skeleton
```js
module.exports = {
  networks: {
    development: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "*"
    }
  },
  compilers: {
    solc: { version: "0.8.19" }
  }
};
```

### Migration file skeleton
```js
const MyContract = artifacts.require("MyContract");

module.exports = function (deployer) {
  deployer.deploy(MyContract);
};
```

### Truffle console cheat sheet
```js
let instance = await MyContract.deployed()          // get deployed instance
let accounts = await web3.eth.getAccounts()          // list Ganache accounts
web3.utils.toWei("1", "ether")                        // convert ETH -> wei
await instance.someFunction(arg1, arg2)               // call a function
await instance.someFunction({ from: accounts[1], value: web3.utils.toWei("1","ether") })
result.toString()                                     // BN -> readable string
```

## 4. Solidity quick syntax
```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Example {
    struct Item { uint id; string name; }
    mapping(uint => Item) public items;
    mapping(address => bool) public seen;

    event ItemAdded(uint id, string name);

    modifier onlyOnce() {
        require(!seen[msg.sender], "Already done");
        _;
        seen[msg.sender] = true;
    }

    function addItem(uint _id, string memory _name) public {
        items[_id] = Item(_id, _name);
        emit ItemAdded(_id, _name);
    }
}
```
Key ideas to remember for viva:
- `public` state variables auto-generate a getter function
- `require(condition, "message")` reverts the transaction if false
- `payable` allows a function/address to receive ETH
- `msg.sender` = caller's address, `msg.value` = ETH sent with the call
- `event` + `emit` = how contracts log data for off-chain listeners

## 5. Hyperledger Fabric — test network
```bash
cd fabric-samples/test-network
./network.sh up                                   # start peers, orderer, CAs
./network.sh createChannel                        # create "mychannel"
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go
./network.sh down                                  # tear down

# Org1 environment (needed before any peer command)
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# Invoke (writes) vs Query (reads)
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com \
  --tls --cafile "<orderer-ca-path>" -C mychannel -n basic \
  --peerAddresses localhost:7051 --tlsRootCertFiles "<peer-ca-path>" \
  -c '{"function":"InitLedger","Args":[]}'

peer chaincode query -C mychannel -n basic -c '{"Args":["GetAllAssets"]}'
```
- **Chaincode** = Fabric's term for a smart contract (usually Go, using `fabric-contract-api-go`)
- **Channel** = a private sub-network/ledger between specific organizations
- **MSP** = Membership Service Provider — Fabric's identity/cert system

## 6. IPFS
```bash
ipfs init                    # create node identity + config
ipfs daemon                  # start the node (keep running)
ipfs add <file>               # store file -> returns CID
ipfs cat <CID>                 # read file contents by CID
ipfs get <CID>                 # download file by CID
ipfs refs local                # list locally stored hashes
ipfs pin add <CID>            # prevent garbage collection
```
- **CID** = Content Identifier — a hash-based address; changes if the file changes (tamper-evidence)

## 7. Python blockchain hashing pattern (used in Practicals 5, 6, 10)
```python
import hashlib

def sha256(text: str) -> str:
    return hashlib.sha256(text.encode()).hexdigest()
```
Standard block anatomy:
```python
class Block:
    def __init__(self, index, data, previous_hash):
        self.index = index
        self.data = data
        self.previous_hash = previous_hash
        self.hash = self.generate_hash()

    def generate_hash(self):
        content = str(self.index) + self.data + self.previous_hash
        return hashlib.sha256(content.encode()).hexdigest()
```
Verification always checks two things per block:
1. `block.previous_hash == previous_block.hash` (chain integrity)
2. `block.hash == block.generate_hash()` recomputed fresh (data integrity)

## 8. Core viva concepts (quick answers)
| Question | Short answer |
|---|---|
| What is a block? | A data container with an index, payload, timestamp, and hash link to the previous block |
| Why SHA-256? | One-way, deterministic, collision-resistant — any data change produces a completely different hash |
| How is tampering detected? | Recompute the hash of current data and compare to the stored hash; any mismatch = tampered |
| What is immutability? | Once a block is chained (its hash referenced by the next block), changing it breaks the chain link, which is detectable |
| Public vs. permissioned blockchain? | Ethereum/Ganache = public/permissionless (anyone can transact); Hyperledger Fabric = permissioned (only known, certified orgs/peers participate) |
| What is a smart contract? | Self-executing code deployed on-chain; its logic and state are enforced by the network, not a single party |
| What is Truffle? | Development framework: compiles, migrates (deploys), and lets you test/interact with Solidity contracts |
| What is Ganache? | A personal local Ethereum blockchain for development/testing, with pre-funded test accounts |
| What is gas? | The fee unit paid to execute operations on Ethereum; prevents infinite loops/spam |
| What is a nonce / duplicate-prevention pattern? | A mapping like `mapping(address => bool) voters` records "has this address already acted" to block repeats |
| What is IPFS used for here? | Off-chain storage of large data, with only the CID (hash) referenced on-chain for verification |
| Why put ML datasets on a blockchain hash instead of the data itself? | Hashing avoids storing large data on-chain (expensive) while still detecting any tampering by comparing hashes |
