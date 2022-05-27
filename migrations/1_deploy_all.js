const TrainingContract = artifacts.require("TrainingContract");
// const TrainingERC20 = artifacts.require("TrainingERC20");
// const TrainingNFT1155 = artifacts.require("TrainingNFT1155");
// const TrainingNFT721 = artifacts.require("TrainingNFT721");


// const BigNumber = require('bignumber.js');
// initialSupply = new BigNumber('1000000000000000000000')

module.exports = async function(deployer, network, accounts) {
    let options = {};

    if (network === "development") {
        options.from = accounts[0];
    }

    await deployer.deploy(TrainingContract, options).then(async () => {
        console.log('Deployed TrainingContract');
    });

    // await deployer.deploy(TrainingNFT1155, options).then(async () => {
    //     console.log('Deployed TrainingNFT1155');
    // });
    //
    // await deployer.deploy(TrainingNFT721).then(async () => {
    //     console.log('Deployed TrainingNFT721');
    // });
    //
    // await deployer.deploy(TrainingERC20, initialSupply).then(async () => {
    //     console.log('Deployed TrainingERC20');
    // });
};