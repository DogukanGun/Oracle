const hre = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();
  
    console.log("Deploying contracts with the account:", deployer.address);
  
    console.log("Account balance:", (await deployer.getBalance()).toString());

    const Oracle = await ethers.getContractFactory("Oracle"); //Replace with name of your smart contract
    const oracle = await Oracle.deploy();
  
    console.log("Contract address for Oracle:", oracle.address);

    const LendingPool = await ethers.getContractFactory("LendingPool"); //Replace with name of your smart contract
    const lendingPool = await LendingPool.deploy(10);
  
    console.log("Contract address for Lending Pool :", lendingPool.address);

    const Wallet = await ethers.getContractFactory("Wallet"); //Replace with name of your smart contract
    const wallet = await Wallet.deploy();
  
    console.log("Contract address for Wallet :", wallet.address);

    console.log("Account balance:", (await deployer.getBalance()).toString());

  }
  
  main()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });