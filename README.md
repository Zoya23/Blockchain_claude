# Blockchain Technologies — Practical Journal

M.Sc. Data Science, Semester IV — University of Mumbai
Source: converted from the handwritten/typed journal PDF into copy-paste-ready code.

📄 **[EXAM_CHEATSHEET.md](EXAM_CHEATSHEET.md)** — one-page reference for the practical exam.

## Index

| # | Practical | Folder | Status |
|---|---|---|---|
| 1 | Installation of Ethereum, Truffle, Ganache, and other tools | [Practical-01-Environment-Setup](Practical-01-Environment-Setup) | Steps only (no code) |
| 2 | Writing and deploying basic smart contracts on Ethereum | [Practical-02-HelloWorld-Contract](Practical-02-HelloWorld-Contract) | ✅ Verified |
| 3 | Configuring and running a Hyperledger Fabric network | [Practical-03-Hyperledger-Fabric-Network](Practical-03-Hyperledger-Fabric-Network) | ⚠️ Corrected (see README) |
| 4 | Storing and retrieving data using IPFS | [Practical-04-IPFS-Storage](Practical-04-IPFS-Storage) | ✅ As journal |
| 5 | Using blockchain to ensure data integrity in ML models | [Practical-05-ML-Data-Integrity](Practical-05-ML-Data-Integrity) | ✅ Verified (ran successfully) |
| 6 | Implementing a blockchain-based data provenance system | [Practical-06-Data-Provenance](Practical-06-Data-Provenance) | ✅ Verified (ran successfully) |
| 7 | Developing a decentralized marketplace for data exchange | [Practical-07-Data-Marketplace](Practical-07-Data-Marketplace) | ✅ As journal |
| 8 | Writing and deploying chaincode on Hyperledger Fabric | [Practical-08-Chaincode-Fabric](Practical-08-Chaincode-Fabric) | ⚠️ Journal content + added real Fabric chaincode (see README) |
| 9 | Implementing a secure voting system using blockchain | [Practical-09-Voting-System](Practical-09-Voting-System) | ⚠️ Reconstructed from journal description (see README) |
| 10 | Combining blockchain with IoT for secure data management | [Practical-10-IoT-Blockchain](Practical-10-IoT-Blockchain) | ⚠️ Reconstructed (verified, ran successfully) |

## How to use this repo
Each `Practical-XX-*` folder is self-contained:
- Source files (`.sol`, `.py`, `.js`, `.go`) — copy directly into your own project/journal
- `README.md` — the Aim, step-by-step instructions, expected output, and any corrections made vs. the original journal

## Things flagged during conversion (read before your exam)
1. **Practical 2** — the journal's `truffle-config.js` was cut off mid-file; completed here.
2. **Practical 3** — the journal's Fabric download link and `peer chaincode` commands were incomplete/outdated for a real TLS-enabled test network; corrected commands included, with the original simplified version also explained.
3. **Practical 8** — the journal's actual content duplicates Practical 7's Solidity contract rather than real Fabric chaincode. Both the as-written version and a real Go chaincode example are included.
4. **Practicals 9 & 10** — the journal only *describes* what the code should contain (no code was pasted in the original document). Working code was written to match the description and journal's expected outputs exactly, and was tested where a Python interpreter was available (P10 verified; P9 needs Solidity/Truffle to test, but the logic directly mirrors the console session and error message already in your journal).

## Publishing to GitHub
```bash
cd Blockchain-Journal
git init
git add .
git commit -m "Blockchain Technologies practical journal"
git branch -M main
git remote add origin <your-repo-url>
git push -u origin main
```
