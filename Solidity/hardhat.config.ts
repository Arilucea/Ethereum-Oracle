import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@nomiclabs/hardhat-ethers";
import dotenv from "dotenv";

dotenv.config({path: __dirname + "/.env"})

module.exports = {
  solidity: {
    version: "0.8.17",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    }
  },
  networks: {
    mumbai: {
      url: process.env.MUMBAI_URL,
      accounts: [process.env.SECRET]
    }
  }
};
