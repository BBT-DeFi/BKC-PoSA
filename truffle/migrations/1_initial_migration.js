const BKCValidatorSet = artifacts.require("BKCValidatorSet");
const StakePool = artifacts.require("StakePool");
const SystemReward = artifacts.require("SystemReward");
const ValidatorPool = artifacts.require("ValidatorPool");

module.exports = function (deployer) {
  deployer.deploy(BKCValidatorSet);
  deployer.deploy(StakePool);
  deployer.deploy(SystemReward);
  deployer.deploy(ValidatorPool);
};
