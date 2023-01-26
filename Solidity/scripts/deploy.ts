/* global ethers hre */
//@ts-ignore
import { ethers } from "hardhat";

export async function deployContract() {

  const Oracle = await ethers.getContractFactory('Oracle')
  const oracleContract = await Oracle.deploy()
  await oracleContract.deployed();

  const Proxy = await (await ethers.getContractFactory('Proxy'))
  const proxy = await Proxy.deploy(oracleContract.address);
  await proxy.deployed();

  let contractC = await (await ethers.getContractAt("Oracle", proxy.address))
  console.log("Owner account", await contractC.owner())

  console.log(`Proxy address ${proxy.address} implementation address ${oracleContract.address}`);

  return(proxy.address)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
if (require.main === module) {
  deployContract()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
}