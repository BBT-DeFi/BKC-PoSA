const BKCValidatorSet = artifacts.require("BKCValidatorSet");
const StakePool = artifacts.require("StakePool");
const SystemReward = artifacts.require("SystemReward");
const ValidatorPool = artifacts.require("ValidatorPool");

var BKCValidatorSetAddress;
var StakePoolAddress;
var SystemRewardAddress;
var ValidatorPoolAddress;

var BKCValidatorSetInstance;
var StakePoolInstance;
var SystemRewardInstance;
var ValidatorPoolInstance;

var validatorTest1;
var validatorTest2;
var validatorTest3;
var validatorTest4;
var validatorTest5;

var delegatorTest1;
var delegatorTest2;
var delegatorTest3;
var delegatorTest4;
var delegatorTest5;

contract("BKCValidatorSet", (accounts) => {
  validatorTest1 = accounts[0];
  validatorTest2 = accounts[1];
  validatorTest3 = accounts[2];
  validatorTest4 = accounts[3];
  validatorTest5 = accounts[5];

  validators = [
    validatorTest1,
    validatorTest2,
    validatorTest3,
    validatorTest4,
    validatorTest5,
  ];

  validators_amount = [10e18, 20e18, 25e18, 22e18, 1e18];

  delegatorTest1 = accounts[0];
  delegatorTest2 = accounts[1];
  delegatorTest3 = accounts[2];
  delegatorTest4 = accounts[3];
  delegatorTest5 = accounts[4];

  delegators = [
    delegatorTest1,
    validatorTest2,
    validatorTest3,
    validatorTest4,
    validatorTest5,
  ];

  it("deployed all contract", async () => {
    BKCValidatorSetInstance = await BKCValidatorSet.deployed();
    StakePoolInstance = await StakePool.deployed();
    SystemRewardInstance = await SystemReward.deployed();
    ValidatorPoolInstance = await ValidatorPool.deployed();
    BKCValidatorSetAddress = BKCValidatorSetInstance.address;
    StakePoolAddress = StakePoolInstance.address;
    SystemRewardAddress = SystemRewardInstance.address;
    ValidatorPoolAddress = ValidatorPoolInstance.address;

    const init_status1 = await BKCValidatorSetInstance.alreadyInit.call();
    const init_status2 = await StakePoolInstance.alreadyInit.call();
    const init_status3 = await SystemRewardInstance.alreadyInit.call();
    const init_status4 = await ValidatorPoolInstance.alreadyInit.call();

    assert.equal(
      init_status1,
      false,
      "the init status of BKCValidatorSet should be false"
    );
    assert.equal(
      init_status2,
      false,
      "the init status of StakePool should be false"
    );
    assert.equal(
      init_status3,
      false,
      "the init status of SystemReward should be false"
    );
    assert.equal(
      init_status4,
      false,
      "the init status of ValidatorPool should be false"
    );
  });

  it("initialize all the contracts", async () => {
    await ValidatorPoolInstance.init(BKCValidatorSetAddress, StakePoolAddress, {
      from: accounts[0],
    });
    await StakePoolInstance.init(ValidatorPoolAddress, { from: accounts[1] });

    await SystemRewardInstance.init(BKCValidatorSetAddress, StakePoolAddress, {
      from: accounts[2],
    });
    await BKCValidatorSetInstance.init(
      ValidatorPoolAddress,
      SystemRewardAddress,
      {
        from: accounts[3],
      }
    );
    const init_status1 = await BKCValidatorSetInstance.alreadyInit.call();
    const init_status2 = await StakePoolInstance.alreadyInit.call();
    const init_status3 = await SystemRewardInstance.alreadyInit.call();
    const init_status4 = await ValidatorPoolInstance.alreadyInit.call();

    assert.equal(
      init_status1,
      true,
      "the init status of BKCValidatorSet should be true"
    );
    assert.equal(
      init_status2,
      true,
      "the init status of StakePool should be true"
    );
    assert.equal(
      init_status3,
      true,
      "the init status of SystemReward should be true"
    );
    assert.equal(
      init_status4,
      true,
      "the init status of ValidatorPool should be true"
    );
  });

  it("register validator (include fail attempts with not enough stake amount)", async () => {
    for (var i = 0; i < 4; i++) {
      await ValidatorPoolInstance.registerValidator({
        from: validators[i],
        value: validators_amount[i],
      });
      const v = await ValidatorPoolInstance.validators.call(i);
      assert.equal(
        v[0],
        validators[i],
        "the validator consensusAddress must be equal"
      );
    }
    try {
      await ValidatorPoolInstance.registerValidator({
        from: validators[4],
        value: validators_amount[4],
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err
          .toString()
          .includes(
            "can't register to be a validator not enough staking amount sent"
          )
      );
    }
    validators_amount[4] = 50e18;
    await ValidatorPoolInstance.registerValidator({
      from: validators[4],
      value: validators_amount[4],
    });
    const v = await ValidatorPoolInstance.validators.call(4);
    assert.equal(
      v[0],
      validators[4],
      "the validator consensusAddress must be equal"
    );

    const number_of_validators =
      await ValidatorPoolInstance.NumberOfValidator.call();
    //console.log(number_of_validators);
    assert.equal(
      number_of_validators,
      5,
      "the number of validators in the pool must be 5"
    );
  });

  it("update the validator set", async () => {
    try {
      await BKCValidatorSetInstance.currentValidatorSet.call(0);
      assert.fail();
    } catch (err) {
      assert.ok(err.toString());
    }
    let map = await BKCValidatorSetInstance.currentValidatorSetMap.call(
      validators[0]
    );
    assert.equal(
      map,
      0,
      "the index must be 0 for all address before the call of updateValidatorSet"
    );
    await BKCValidatorSetInstance.updateValidatorSet({ from: validators[0] });
    let num = await BKCValidatorSetInstance.number_of_validators.call();
    assert.equal(num, 2, "the number of validators must be 2");
    const v1 = await BKCValidatorSetInstance.currentValidatorSet.call(0);
    assert.equal(
      v1[0],
      validators[4],
      "the leading validator must be the 5th validator"
    );
    assert.equal(v1[1], 50e18, "with 50 ether amount");
    const v2 = await BKCValidatorSetInstance.currentValidatorSet.call(1);
    assert.equal(
      v2[0],
      validators[2],
      "the second leading validator must be the third validator"
    );
    assert.equal(v2[1], 25e18, "with 25 ether amount");
  });

  it("validator can top-up", async () => {
    ValidatorPoolInstance.validatorTopUp(validators[0], {
      from: validators[0],
      value: 16e18,
    });
    const index = await ValidatorPoolInstance.validatorsMap.call(validators[0]);
    const v = await ValidatorPoolInstance.validators.call(index - 1);
    //console.log(v);
    assert.equal(v[1], 26e18, "the amount must be updated");

    await BKCValidatorSetInstance.updateValidatorSet({ from: validators[0] });
    const v1 = await BKCValidatorSetInstance.currentValidatorSet.call(0);
    assert.equal(
      v1[0],
      validators[4],
      "the first validator should be the same"
    );
    assert.equal(v1[1], 50e18, "the first validator should be the same");
    const v2 = await BKCValidatorSetInstance.currentValidatorSet.call(1);
    assert.equal(
      v2[0],
      validators[0],
      "the second validator now should be the just-top-up validator"
    );
    assert.equal(
      v2[1],
      26e18,
      "the second validator now should be the just-top-up validator"
    );
  });
});
