var HDWalletProvider = require("@truffle/hdwallet-provider");
var PrivateKeyProvider = require("truffle-privatekey-provider");

require('dotenv').config();

productPrivateKey = process.env.PRIVATE_KEY;
infuraProjectKey = process.env.INFURA_PROJECT_KEY;

module.exports = {
    contracts_directory: "./contracts",
    contracts_build_directory: "./output",
    migrations_directory: "./migrations",
    networks: {
        development: {
            host: "127.0.0.1",
            port: 8545, //7545
            network_id: "4447", //5777
            from: "0xbE7b24346122B58D678BD4E12E51077214e62032", //need to be reset for your local dev network
            gasPrice: 875000000,

            // optional config values:
            // gas
            // gasPrice
            // from - default address to use for any transaction Truffle makes during migrations
            // provider - web3 provider instance Truffle should use to talk to the Ethereum network.
            //          - function that returns a web3 provider instance (see below.)
            //          - if specified, host and port are ignored.
            // skipDryRun: - true if you don't want to test run the migration locally before the actual migration (default is false)
            // confirmations: - number of confirmations to wait between deployments (default: 0)
            // timeoutBlocks: - if a transaction is not mined, keep waiting for this number of blocks (default is 50)
            // deploymentPollingInterval: - duration between checks for completion of deployment transactions
            // disableConfirmationListener: - true to disable web3's confirmation listener
        },
        rinkeby: {
            host: "localhost",
            provider: function () {
//                 return new HDWalletProvider(mnemonicFirefox, "https://rinkeby.infura.io/v3/"+infuraProjectKey);
                return new PrivateKeyProvider(productPrivateKey, "https://rinkeby.infura.io/v3/"+infuraProjectKey);
            },
            network_id:4,
            gas : 25000000,
            gasPrice : 10000000000
        },
        mainnet: {
            provider: function () {
                return new HDWalletProvider(mnemonic, "https://mainnet.infura.io/v3/" + infuraProjectKey);
            },
            network_id: 1,
            gas: 2600000,
            gasPrice: 15000000000
        },
        fantom: {
            provider: function () {
                return new HDWalletProvider(mnemonic, "https://rpcapi.fantom.network");
            },
            network_id: 250,
            gas: 3000000,
            gasPrice: 50000000000
        },
        bsc: {
            provider: () => new HDWalletProvider(mnemonic, "https://bsc-dataseed1.binance.org"),
            network_id: 56,
            gas: 3000000,
            gasPrice: 5000000000,
            confirmations: 10,
            timeoutBlocks: 200,
            skipDryRun: true
        },
    },
    compilers: {
        solc: {
            version: "^0.8.11",
            settings: {
                // optimizer: {
                //     enabled: true,
                //     runs: 1500
                // }
            }
        }
    }
};