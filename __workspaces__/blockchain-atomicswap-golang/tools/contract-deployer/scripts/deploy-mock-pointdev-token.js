// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const hre = require("hardhat");

async function main() {

  const contractName = 'PointDev'
  const owner = "0x872d3d0d6C5c1C0f5E8f9EEc2c4236cc9b5AB823";
  const alice = "0xd5a38dD251Aa8493C03954268FF851290051E634";

  console.log(`ready to deploy ${contractName}.sol`)
  const token = await hre.ethers.deployContract(contractName);

  console.log(`deploying ${contractName}.sol`);
  const tokenDeployed = await token.waitForDeployment();

  console.log(`the deployed ${contractName} contract address is : ` + tokenDeployed.target);

  const deployerBalance = await tokenDeployed.balanceOf(owner);
  console.log(deployerBalance);
  const ownerSigner = (await ethers.getSigners())[0]
  const decimals = await tokenDeployed.decimals();
  const amount = 10n * (10n ** decimals);
  console.log(amount)
  const sended = await tokenDeployed.connect(ownerSigner).transfer(alice, amount);
  console.log(sended)
  const result = await tokenDeployed.balanceOf(alice)
  console.log(result);
  console.log(await tokenDeployed.balanceOf(owner))
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
