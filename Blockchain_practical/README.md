# Practical 1 — Installation of Ethereum, Truffle, Ganache, and Other Tools

## Aim
Set up a complete local Ethereum development environment: Node.js, Truffle, Ganache, MetaMask, VS Code and Git Bash.

## Step 1: Install Node.js
1. Go to the official Node.js download page: https://nodejs.org
2. Download the **LTS** version for Windows and run the installer.
3. Verify the install:
```bash
node -v
npm -v
```

## Step 2: Install Truffle
```bash
npm install -g truffle
truffle version
```

## Step 3: Install Ganache
1. Download Ganache from https://trufflesuite.com/ganache
2. Install it, open it, and select **Quickstart Ethereum**.
3. Note down (you'll need these later):
   - **RPC URL** — usually `http://127.0.0.1:7545`
   - Test accounts, private keys, and balances shown in the UI

## Step 4: Install MetaMask
1. Install the [MetaMask](https://metamask.io/) browser extension.
2. Create a wallet and **save the recovery phrase somewhere safe**.
3. Add Ganache as a custom network in MetaMask:
   - RPC URL: the Ganache RPC URL noted above
   - Chain ID: the chain ID shown by Ganache
4. Import one Ganache account into MetaMask using its private key.

## Step 5: Install VS Code
1. Download from https://code.visualstudio.com/
2. Install the **Solidity** extension from the Extensions marketplace.

## Step 6: Install Git Bash
Install [Git for Windows](https://git-scm.com/download/win) so you have a bash-style terminal for blockchain commands.

## Verification Checklist
| Tool | Verify with |
|---|---|
| Node.js | `node -v` |
| npm | `npm -v` |
| Truffle | `truffle version` |
| Ganache | UI opens, Quickstart workspace running |
| MetaMask | Connected to Ganache network, one account imported |

> This practical has no source code — it's an environment setup checklist. Practical 2 onward uses the tools installed here.
