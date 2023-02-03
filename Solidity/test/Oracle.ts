/* global ethers hre */
//@ts-ignore
import { ethers } from "hardhat";
import {deployContract} from "../scripts/deploy";
import { expect } from "chai";
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'; 
import {Oracle, Oracle__factory, Proxy, Proxy__factory, PetitionExample, PetitionExample__factory } from "../typechain-types"

describe("Oracle", function () {

    let oracle: Oracle; 
    let proxy: Proxy;
    let requesterContract: PetitionExample;

    let owner: SignerWithAddress;
    let user: SignerWithAddress;

    let proxyAddress: string;


    before(async function () {        
        proxyAddress = await deployContract();
        proxy = await ethers.getContractAt('Proxy', proxyAddress);
        oracle = await ethers.getContractAt('Oracle', proxyAddress);

        [owner, user] = await ethers.getSigners();
    })

    describe("Deployment and implmentation", async function () {
        it('Should set the right owner', async function () {
            expect(await oracle.owner()).to.equal(owner.address);
        });

        it('Should be possible to update the implementation', async function () {
            const newImplementation = await new Oracle__factory(owner).deploy();
            await newImplementation.deployed();
            
            await proxy.setImplementation(newImplementation.address);
            expect(await proxy.getImplementation()).to.equal(newImplementation.address);
        })

        it('Should only be possible to update the implementation by the owner', async function () {
            const newImplementation = await new Oracle__factory(owner).deploy();
            await newImplementation.deployed();
            
            await expect(proxy.connect(user).setImplementation(newImplementation.address)).to.be.revertedWith('Not the owner of the contract');
        });
    });

    describe('Oracle petitions', async function () {
        it("Should receive petitions only from contracts", async function () {
            const message = "SomeURL";
            const type = "String";
            await expect(oracle.makePetition(message, type)).to.be.revertedWith('The caller is not a contract');
        });

        it("When a request is donde should emit an event with the petition data", async function () {
            const message = "SomeURL";
            const type = "String";
            requesterContract = await new PetitionExample__factory(owner).deploy(proxyAddress);

            await expect(requesterContract.callInt(message, type, {value: 1000000})).to.emit(oracle, "Petition").withArgs(requesterContract.address, message, type);
        });
    });

});