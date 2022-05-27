
const TrainingContract = artifacts.require("TrainingContract");


let location = {
    addressLine1: "test",
    addressLine2: "test1",
    city: "Poltava",
    province: "no",
    countryCode: "380",
    postalCode: "36000",
}


// Variables for tokens minting
let amount = 100;
let data = 0;



contract('TrainingContract', (accounts) => {
    console.log(accounts);

    let user1 = accounts[0];
    let user2 = accounts[1];
    let user3 = accounts[2];


    it('WineBrand setters and getters:', async () => {
        const tc = await TrainingContract.deployed();


        // let tx1 = await tc.setName('Petro', {from: user1});
        // let tx2 = await tc.getName({from: user2});
        // console.log(tx2);
        //
        // assert.equal('Petro', tx2, "Wrong name")

        let tx3 = await tc.setLocation(location, {from: user1});
        let tx4 = await tc.getLocation({from: user2});

        // console.log(tx4);
        for (let i in tx4.logs) {
            if (tx4.logs[i].event == "LOCATION") {
                console.log(tx4.logs[i].args['0']);
                loc = tx4.logs[i].args['0'];
                break;
            }
        }


    });


});

