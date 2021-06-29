const Web3 = require("web3");
const BKCValidatorSetABI = require("./truffle/build/contracts/BKCValidatorSet.json");
const StakePoolABI = require("./truffle/build/contracts/StakePool.json");
const SystemRewardABI = require("./truffle/build/contracts/SystemReward.json");
const ValidatorPoolABI = require("./truffle/build/contracts/ValidatorPool.json");

const BKCValidatorSetAddress = "0x96cC8e9011049eDF9f0468E0d25231bD3db93a91";
const StakePoolAddress = "0x41A1e8E3A61F85396A6e39E3DBFFdcC70D340C36";
const SystemRewardAddress = "0xa3B9A38549fF9DB7F9F2b7b3de45fD9cD1850CF4";
const ValidatorPoolAddress = "0x2ddE8835b6D60911dc5395F07315a7E846d12cdd";

const runTest = async () => {
  const web3 = new Web3("http://localhost:7545");

  const accounts = await web3.eth.getAccounts();

  const BKCValidatorSet = new web3.eth.Contract(
    BKCValidatorSetABI.abi,
    BKCValidatorSetAddress
  );
  const StakePool = new web3.eth.Contract(StakePoolABI.abi, StakePoolAddress);
  const SystemReward = new web3.eth.Contract(
    SystemRewardABI.abi,
    SystemRewardAddress
  );
  const ValidatorPool = new web3.eth.Contract(
    ValidatorPoolABI.abi,
    ValidatorPoolAddress
  );

  await ValidatorPool.methods
    .init(BKCValidatorSetAddress, StakePoolAddress)
    .send({ from: accounts[0] });
  await StakePool.methods
    .init(ValidatorPoolAddress)
    .send({ from: accounts[0] });
  await SystemReward.methods
    .init(BKCValidatorSetAddress, StakePoolAddress)
    .send({ from: accounts[0] });
  await BKCValidatorSet.methods
    .init(ValidatorPoolAddress, SystemRewardAddress)
    .send({ from: accounts[0] });

  const init_status1 = BKCValidatorSet.methods.alreadyInit().call();
  const init_status2 = StakePool.methods.alreadyInit().call();
  const init_status3 = SystemReward.methods.alreadyInit().call();
  const init_status4 = ValidatorPool.methods.alreadyInit().call();

  console.log(init_status1, init_status2, init_status3, init_status4);
};

runTest();
