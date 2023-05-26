# idp08_selfoptimizing_lp
IDP 08: Self-Optimizing Lending Protocol


LendingPool contract deployed to Sepolia Testnet at:
`0x2C10EDb2480322B6800Fe479d32b5A5e436d3b33`

Contract verified here: https://sepolia.etherscan.io/address/0x2C10EDb2480322B6800Fe479d32b5A5e436d3b33#code

## Deployment and EtherScan Verification

The following scripts were run to deploy and verify the contract to the Sepolia testnet:
```
npx hardhat run scripts/deploy.js --network localhost
```
## Change .env_example to .env 
Change it with your credentials and change the name of it. Instructions are given inside the .env_example file.

## Creating Abi in Go

```

solc --abi --bin Store.sol -o build
abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=store --out=Store.go
```

