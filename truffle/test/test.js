const Web3 = require("web3");
const colors = require("colors");
const web3 = new Web3("http://localhost:7545");

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

const BondStatus = {
  BONDED: 0,
  UNBONDING: 1,
  UNBONDED: 2,
};

const BondStatusMap = {
  0: "BONDED",
  1: "UNBONDING",
  2: "UNBONDED",
};

async function getBalance(address) {
  const balance = await web3.eth.getBalance(address);
  return balance;
}

function printValidator(v) {
  console.log(
    "validator" +
      validators.indexOf(v[0]) +
      " address: " +
      v[0] +
      " stakeAmount: " +
      v[1] / Math.pow(10, 18) +
      " bond status: " +
      BondStatusMap[v[2]] +
      " isJail: " +
      v[3]
  );
}

async function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function printUserUnDelegateQueue() {
  console.log(colors.green("\nDelegators in User Undelegate Queue"));
  for (var j = 0; j < delegators.length; j++) {
    for (var i = 0; i < validators.length; i++) {
      const amount =
        await StakePoolInstance.getUserDelegationUnbondingAmountCallable(
          delegators[j],
          validators[i]
        );
      if (amount > 0) {
        const Amount = await StakePoolInstance.getUnbondingQueueAmount.call(
          delegators[j]
        );
        const Time = await StakePoolInstance.getUnbondingQueueTime.call(
          delegators[j]
        );
        const Validator =
          await StakePoolInstance.getUnbondingQueueValidator.call(
            delegators[j]
          );
        console.log("for delegator" + j + " : ");
        for (var k = 0; k < Amount.length; k++) {
          console.log(
            k +
              ". Amount : " +
              Amount[k] / Math.pow(10, 18) +
              " ether " +
              "will finish unbonding at: " +
              new Date(Time[k].toNumber() * 1000).toLocaleString() +
              " Validator : validator" +
              validators.indexOf(Validator[k])
          );
        }
        break;
      }
    }
  }
}

async function printState() {
  async function BKCValidatorSetState() {
    const num = await BKCValidatorSetInstance.number_of_validators.call();
    console.log(colors.yellow("\nValidators in currentValidatorSet"));
    for (var i = 0; i < num; i++) {
      const v = await BKCValidatorSetInstance.currentValidatorSet.call(i);
      printValidator(v);
    }
  }
  async function ValidatorPoolState() {
    console.log(colors.cyan("\nValidators in Validator pool"));
    for (var i = 0; i < validators.length; i++) {
      try {
        const v = await ValidatorPoolInstance.validators.call(i);
        printValidator(v);
      } catch (err) {}
    }
    console.log(colors.blue("\nValidators in Validator Unbonding Queue"));
    for (var i = 0; i < validators.length; i++) {
      try {
        const v = await ValidatorPoolInstance.validators.call(i);
        if (v[2] == BondStatus.UNBONDING && !v[3]) {
          // not jail
          printValidator(v);
          const queue = await ValidatorPoolInstance.validatorUnBondQueue(v[0]);
          console.log(
            "will finish unbonding at: " +
              new Date(queue.toNumber() * 1000).toLocaleString()
          );
        }
      } catch (err) {}
    }
    console.log(colors.red("\nValidators in Validator Jail Queue"));
    for (var i = 0; i < validators.length; i++) {
      try {
        const v = await ValidatorPoolInstance.validators.call(i);
        if (v[3]) {
          // Jail
          printValidator(v);
          const queue = await ValidatorPoolInstance.validatorJailQueue(v[0]);
          console.log(
            "will finish jailing at: " +
              new Date(queue.toNumber() * 1000).toLocaleString()
          );
        }
      } catch (err) {}
    }
  }
  async function StakePoolState() {
    console.log(colors.magenta("\nDelegations in Stake Pool"));
    for (var i = 0; i < validators.length; i++) {
      try {
        const v = await ValidatorPoolInstance.validators.call(i);
        const all_delegators = await StakePoolInstance.getDelegators(v[0]);
        console.log("for validator" + validators.indexOf(v[0]) + " : ");
        for (var j = 0; j < all_delegators.length; j++) {
          const totalDelegate =
            await StakePoolInstance.getUserTotalDelegationAmountCallable(
              all_delegators[j],
              v[0]
            );
          console.log(
            "total delegation from delegator" +
              delegators.indexOf(all_delegators[j]) +
              " is " +
              web3.utils.fromWei(totalDelegate.toString()) +
              "e18"
          );
        }
      } catch (err) {}
    }
    try {
      await printUserUnDelegateQueue();
    } catch (err) {}
  }
  async function SystemRewardState() {
    const num = await BKCValidatorSetInstance.number_of_validators.call();
    console.log(
      colors.brightBlue(
        "\nReward in reward mapping for the currently active validators"
      )
    );
    for (var i = 0; i < num; i++) {
      const v = await BKCValidatorSetInstance.currentValidatorSet.call(i);
      const reward = await SystemRewardInstance.rewardMapping.call(v[0]);
      console.log(
        "for validator" +
          validators.indexOf(v[0]) +
          " : the pending reward is " +
          reward / Math.pow(10, 18) +
          " ether"
      );
    }
  }
  await BKCValidatorSetState();
  await ValidatorPoolState();
  await StakePoolState();
  await SystemRewardState();
}

async function checkDelegation(
  delegatorNum,
  validatorNum,
  bondedAmount,
  unbondingAmount,
  totalAmount
) {
  const delegateBondedBalance =
    await StakePoolInstance.getUserDelegationBondedAmountCallable(
      delegators[delegatorNum],
      validators[validatorNum]
    );
  assert.equal(
    delegateBondedBalance,
    bondedAmount,
    "the bonded delegation balance must be " +
      bondedAmount / Math.pow(10, 18) +
      " ether"
  );
  const delegateUnbondingBalance =
    await StakePoolInstance.getUserDelegationUnbondingAmountCallable(
      delegators[delegatorNum],
      validators[validatorNum]
    );
  assert.equal(
    delegateUnbondingBalance,
    unbondingAmount,
    "the unbonding balance must be " +
      unbondingAmount / Math.pow(10, 18) +
      " ether"
  );
  const delegateTotalBalance =
    await StakePoolInstance.getUserTotalDelegationAmountCallable(
      delegators[delegatorNum],
      validators[validatorNum]
    );
  assert.equal(
    delegateTotalBalance,
    totalAmount,
    "the total delegation must be " + totalAmount / Math.pow(10, 18) + " ether"
  );
}

contract("All Contracts", (accounts) => {
  validatorTest1 = accounts[0];
  validatorTest2 = accounts[1];
  validatorTest3 = accounts[2];
  validatorTest4 = accounts[3];
  validatorTest5 = accounts[4];
  validatorTest6 = accounts[10];

  validators = [
    validatorTest1,
    validatorTest2,
    validatorTest3,
    validatorTest4,
    validatorTest5,
    validatorTest6,
  ];

  validators_amount = [10e18, 20e18, 25e18, 22e18, 1e18];

  delegatorTest1 = accounts[5];
  delegatorTest2 = accounts[6];
  delegatorTest3 = accounts[7];
  delegatorTest4 = accounts[8];
  delegatorTest5 = accounts[9];

  delegators = [
    delegatorTest1,
    delegatorTest2,
    delegatorTest3,
    delegatorTest4,
    delegatorTest5,
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
    await printState();
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

  // validator
  it("register validator (include fail attempts with not enough stake amount : [10e18, 20e18, 25e18, 22e18, 1e18] => [10e18, 20e18, 25e18, 22e18, 50e18])", async () => {
    await printState();
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

  it("update the validator set : [10e18, 20e18, 25e18, 22e18, 50e18] active validators => (4,2)", async () => {
    await printState();
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

  it("validator can top-up : validator 0 do top up 16e18 [10e18, 20e18, 25e18, 22e18, 50e18] => [26e18, 20e18, 25e18, 22e18, 50e18]", async () => {
    await printState();
    await ValidatorPoolInstance.validatorTopUp({
      from: validators[0],
      value: 16e18,
    });
    const index = await ValidatorPoolInstance.validatorsMap.call(validators[0]);
    const v = await ValidatorPoolInstance.validators.call(index - 1);
    //console.log(v);
    assert.equal(v[1], 26e18, "the amount must be updated");
  });

  it("update of the validator set : [26e18, 20e18, 25e18, 22e18, 50e18] active validators => (4,0)", async () => {
    await printState();
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

  it("the unbonding validator is in the validatorUnBondQueue", async () => {
    await printState();
    const queue = await ValidatorPoolInstance.validatorUnBondQueue.call(
      validators[2]
    );
    var now = new Date();
    var tomorrow = new Date();

    assert.equal(
      Math.floor(now / 1000) < queue,
      true,
      "the queue should be set for 1 day after"
    );
    tomorrow.setDate(tomorrow.getDate() + 1);
    assert.equal(
      Math.floor(tomorrow / 1000) > queue,
      true,
      "it must be able to dequeue the validator tomorrow"
    );
  });

  it("still can't dequeue the unbonding validator now", async () => {
    await printState();
    try {
      await ValidatorPoolInstance.removeUnBondingValidatorFromQueue({
        from: validators[2],
      });
      assert.fail();
    } catch (err) {
      assert.ok(err.toString().includes("unbonding still in progress"));
    }
  });

  it("after the update of the validator set : validator 2's bondStatus must be unbonding ", async () => {
    await printState();
    const v2 = await ValidatorPoolInstance.validators.call(2);
    assert.equal(
      v2[2],
      BondStatus.UNBONDING,
      "the bond status of the second validator must be unbonding"
    );
  });

  it("validator can withdraw fund : validator 1 do withdraw fund 8e18 [26e18, 20e18, 25e18, 22e18, 50e18] => [26e18, 12e18, 25e18, 22e18, 50e18]", async () => {
    await printState();
    const balanceBefore = await getBalance(validators[1]);
    await ValidatorPoolInstance.withdrawFund((8e18).toString(), {
      from: validators[1],
    });
    const v1 = await ValidatorPoolInstance.validators.call(1);
    assert.equal(
      v1[1],
      12e18,
      "the stakeAmount of the validator 1 must be 12 ether"
    );
    const balanceAfter = await getBalance(validators[1]);
    assert.equal(
      balanceAfter - balanceBefore >= 7e18,
      true,
      "the withdraw is returned"
    );
  });

  it("update of the validator set nothing changes: [26e18, 12e18, 25e18, 22e18, 21e18] active validators => (4,0)", async () => {
    await printState();
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

  it("validator can't withdraw fund or top-up if bonded : validator 0 try to withdraw and top-up", async () => {
    await printState();
    const v0 = await ValidatorPoolInstance.validators.call(0);
    assert.equal(v0[2], BondStatus.BONDED);
    try {
      await ValidatorPoolInstance.withdrawFund((8e18).toString(), {
        from: validators[0],
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err.toString().includes("can't do operation when unbonding or bonded")
      );
    }
    try {
      await ValidatorPoolInstance.validatorTopUp({
        from: validators[0],
        value: 10e18,
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err.toString().includes("can't do operation when unbonding or bonded")
      );
    }
  });

  it("validator can't withdraw fund or top-up if unbonding : validator 2 try to withdraw and top-up", async () => {
    await printState();
    const v2 = await ValidatorPoolInstance.validators.call(2);
    assert.equal(v2[2], BondStatus.UNBONDING);
    try {
      await ValidatorPoolInstance.withdrawFund((8e18).toString(), {
        from: validators[2],
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err.toString().includes("can't do operation when unbonding or bonded")
      );
    }
    try {
      await ValidatorPoolInstance.validatorTopUp({
        from: validators[2],
        value: 10e18,
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err.toString().includes("can't do operation when unbonding or bonded")
      );
    }
  });

  it("validator withdraw too much fund that its fund is less than the minimum stake amount for validators : validator 1 do withdraw fund 8e18 [26e18, 12e18, 25e18, 22e18, 21e18] => [26e18, 25e18, 22e18, 21e18]", async () => {
    await printState();
    const balanceBefore = await getBalance(validators[1]);
    await ValidatorPoolInstance.withdrawFund((8e18).toString(), {
      from: validators[1],
    });
    const v1 = await ValidatorPoolInstance.validators.call(1); // the validator now should change.
    assert.equal(
      v1[1],
      25e18,
      "the stakeAmount of the validator 1 must be 25 ether"
    );
    const balanceAfter = await getBalance(validators[1]);
    assert.equal(
      balanceAfter - balanceBefore >= 11e18, // 12 ether returned
      true,
      "the balance is all returned"
    );
    const num = await ValidatorPoolInstance.NumberOfValidator.call();
    assert.equal(num, 4);
  });

  it("renew registering validator : register validator1 with registering fund 20e18 [26e18, 25e18, 22e18, 21e18] => [26e18, 25e18, 22e18, 21e18, 20e18])", async () => {
    await printState();
    await ValidatorPoolInstance.registerValidator({
      from: validators[1],
      value: validators_amount[1],
    });
    const v4 = await ValidatorPoolInstance.validators.call(4);
    assert.equal(
      v4[0],
      validators[1],
      "the validator consensusAddress must be equal"
    );
    assert.equal(v4[1], 20e18, "the stake amount must be 20 ether");

    const number_of_validators =
      await ValidatorPoolInstance.NumberOfValidator.call();
    //console.log(number_of_validators);
    assert.equal(
      number_of_validators,
      5,
      "the number of validators in the pool must be 5"
    );
  });

  it("try to register the same validator : registering validator1 again", async () => {
    await printState();
    try {
      await ValidatorPoolInstance.registerValidator({
        from: validators[1],
        value: validators_amount[1],
      });
      assert.fail();
    } catch (err) {
      assert.ok(err.toString().includes("can't register same validator"));
    }

    const number_of_validators =
      await ValidatorPoolInstance.NumberOfValidator.call();
    //console.log(number_of_validators);
    assert.equal(
      number_of_validators,
      5,
      "the number of validators in the pool must be 5"
    );
  });

  // delegation
  it("can't delegate to a bonded validator : delegator2 try to delegate to a bonded validator0", async () => {
    await printState();
    try {
      await StakePoolInstance.delegate(validators[0], {
        from: delegators[2],
        value: 1e18,
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err
          .toString()
          .includes(
            "can't delegate to an bonding delegator or an unbonding delegator"
          )
      );
    }
  });

  it("can't delgate to an unbonding validator : delegator1 try to delegate to an unbonding validator2", async () => {
    await printState();
    try {
      await StakePoolInstance.delegate(validators[2], {
        from: delegators[1],
        value: 1e18,
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err
          .toString()
          .includes(
            "can't delegate to an bonding delegator or an unbonding delegator"
          )
      );
    }
  });

  it("can't delegate to a non-validator : delegator1 try to delegate to delegator2", async () => {
    await printState();
    try {
      await StakePoolInstance.delegate(delegators[2], {
        from: delegators[1],
        value: 1e18,
      });
      assert.fail();
    } catch (err) {
      assert.ok(err.toString().includes("can't delegate to a non-validator"));
    }
  });

  it("can delegate to an unbonded validator : delegator2 delegates 1e18 to validator3", async () => {
    await printState();
    await StakePoolInstance.delegate(validators[3], {
      from: delegators[2],
      value: 1e18,
    });
    await checkDelegation(2, 3, 1e18, 0, 1e18);
  });

  it("can delegate to the same validator : delegator2 delegates another 4e18 to validator3", async () => {
    await printState();
    await StakePoolInstance.delegate(validators[3], {
      from: delegators[2],
      value: 4e18,
    });
    await checkDelegation(2, 3, 5e18, 0, 5e18);
  });

  it("can delegate to multiple validator at the same time : delegator2 delegates 1e18 to validator1", async () => {
    await printState();
    await StakePoolInstance.delegate(validators[1], {
      from: delegators[2],
      value: 1e18,
    });
    await checkDelegation(2, 1, 1e18, 0, 1e18);
  });

  it("multiple delegators can delegate to the same validator : delegator1 delegates 1e18 to validator3", async () => {
    await printState();
    await StakePoolInstance.delegate(validators[3], {
      from: delegators[1],
      value: 1e18,
    });
    await checkDelegation(1, 3, 1e18, 0, 1e18);
  });

  it("multiple delegators can delegate to the same validator : delegator4 delegates 2e18 to validator1", async () => {
    await printState();
    await StakePoolInstance.delegate(validators[1], {
      from: delegators[4],
      value: 2e18,
    });
    await checkDelegation(4, 1, 2e18, 0, 2e18);
  });

  it("if validator withdraw fund that it's kicked out of the validator pool, then the delegations also disappear(returned) : validator 1 withdraw 12e18 triggering return of delegation of delegator2 and delegator4", async () => {
    await printState();
    const balanceBefore = await getBalance(validators[1]);
    const balanceDelegator2Before = await getBalance(delegators[2]);
    const balanceDelegator4Before = await getBalance(delegators[4]);
    await ValidatorPoolInstance.withdrawFund((12e18).toString(), {
      from: validators[1],
    });
    const balanceAfter = await getBalance(validators[1]);
    const balanceDelegator2After = await getBalance(delegators[2]);
    const balanceDelegator4After = await getBalance(delegators[4]);
    assert.equal(
      balanceAfter - balanceBefore >= 19e18, // 20 ether returned
      true,
      "the balance is all returned"
    );
    assert.equal(
      balanceDelegator2After - balanceDelegator2Before >= 9e17 &&
        balanceDelegator2After - balanceDelegator2Before <= 11e17, // 1 ether returned
      true,
      "the returned balance should be 1e18"
    ); // 1 ether returned.
    assert.equal(
      balanceDelegator4After - balanceDelegator4Before >= 19e17 &&
        balanceDelegator4After - balanceDelegator4Before <= 21e17, // 2 ether returned
      true,
      "the returned baalnce should be 2e18"
    );
    const num = await ValidatorPoolInstance.NumberOfValidator.call();
    assert.equal(num, 4); // only 4 validators remain.
    await checkDelegation(2, 1, 0, 0, 0);
    await checkDelegation(4, 1, 0, 0, 0);
  });

  it("the validation of delegator2 to other validator must not change", async () => {
    await printState();
    await checkDelegation(2, 3, 5e18, 0, 5e18);
  });

  it("update of validator set : active validator is then be (4,3) due to total delegation of 6 ether of validator3", async () => {
    await printState();
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
      validators[3],
      "the second validator now should be the just-delegated validator"
    );
    assert.equal(
      v2[1],
      22e18,
      "the second validator now should be the j validator"
    );
  });

  it("re-entering of validator1 with registering fund 60e18", async () => {
    await printState();
    await ValidatorPoolInstance.registerValidator({
      from: validators[1],
      value: 60e18,
    });
    const v = await ValidatorPoolInstance.validators.call(4);
    assert.equal(
      v[0],
      validators[1],
      "the validator consensusAddress must be equal"
    );
  });

  it("the re-delegation of the ex-delegated delegator works : delegator2 delegates 2e18 to validator1", async () => {
    await printState();
    await StakePoolInstance.delegate(validators[1], {
      from: delegators[2],
      value: 2e18,
    });
    await checkDelegation(2, 1, 2e18, 0, 2e18);
  });

  it("the re-delegation of the ex-delegated delegator works : delegator4 delegates 1e18 to validator1", async () => {
    await printState();
    await StakePoolInstance.delegate(validators[1], {
      from: delegators[4],
      value: 1e18,
    });
    await checkDelegation(4, 1, 1e18, 0, 1e18);
    const power = await ValidatorPoolInstance.getTotalPower.call(validators[1]);
    assert.equal(power, 63e18, "the total power must be 63 ether");
  });

  it("update of validator set : active validator is then be (4,1)", async () => {
    await printState();
    await BKCValidatorSetInstance.updateValidatorSet({ from: validators[0] });
    const v1 = await BKCValidatorSetInstance.currentValidatorSet.call(0);
    assert.equal(
      v1[0],
      validators[1],
      "the first validator should be the just-entering validator (validator1)"
    );
    assert.equal(
      v1[1],
      60e18,
      "the first validator should have stake amount of 60 ether"
    );
    const v2 = await BKCValidatorSetInstance.currentValidatorSet.call(1);
    assert.equal(
      v2[0],
      validators[4],
      "the second validator now should be the 4th validator"
    );
    assert.equal(
      v2[1],
      50e18,
      "the second validator should have stake amount of 50 ether"
    );
  });

  it("undelegate from a bonded validator will result in UserUnDelegationQueue with nothing changes to the delegation object : delegator2 undelegates 1e18 from validator1 which is an active validator", async () => {
    await printState();
    await StakePoolInstance.undelegate(validators[1], (1e18).toString(), {
      from: delegators[2],
    });
    await checkDelegation(2, 1, 1e18, 1e18, 2e18);
  });

  it("undelegate from an unbonding validator will result in UserUnDelegationQueue with nothing changes to the delegation object : delegator2 undelegates 1e18 from an unbonding validator3", async () => {
    await printState();
    await StakePoolInstance.undelegate(validators[3], (1e18).toString(), {
      from: delegators[2],
    });
    await checkDelegation(2, 3, 4e18, 1e18, 5e18);
  });

  it("register validator5 with registering fund of 12e18", async () => {
    await printState();
    await ValidatorPoolInstance.registerValidator({
      from: validators[5],
      value: 12e18,
    });
    const v = await ValidatorPoolInstance.validators.call(5);
    assert.equal(
      v[0],
      validators[5],
      "the validator consensusAddress must be equal"
    );
  });

  it("delegator3 delegates 3e18 to validator5", async () => {
    await printState();
    await StakePoolInstance.delegate(validators[5], {
      from: delegators[3],
      value: 3e18,
    });
    await checkDelegation(3, 5, 3e18, 0, 3e18);
  });

  it("undelegation from an unbonded validator will undelegate and change the delegation object immediately : delegator3 undelegates 2e18 from an unbonded validator5", async () => {
    await printState();
    const balanceBefore = await getBalance(delegators[3]);
    await StakePoolInstance.undelegate(validators[5], (2e18).toString(), {
      from: delegators[3],
    });
    const balanceAfter = await getBalance(delegators[3]);
    await checkDelegation(3, 5, 1e18, 0, 1e18);
    assert.equal(
      balanceAfter - balanceBefore >= 19e17,
      true,
      "the returned fund must be 2 ether"
    );
  });

  // reward
  it("can't add reward to a non-active validator", async () => {
    await printState();
    try {
      await SystemRewardInstance.addReward(validators[3], {
        from: accounts[19],
        value: 30e18,
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err.toString().includes("can't add reward to a non-active validator")
      );
    }
  });

  it("it can add reward to a specific active validator : 70 ether added to validator1 map", async () => {
    await printState();
    await SystemRewardInstance.addReward(validators[1], {
      from: accounts[19],
      value: 70e18,
    });
  });

  it("the reward added is split with the system at the rate of 1:9 or 10% : added 70 ether to validator1's rewardMapping 63 must be in the mapping", async () => {
    await printState();
    const reward = await SystemRewardInstance.rewardMapping.call(validators[1]);
    assert.equal(reward, 63e18, "the total reward is deducted by 7e18");
    const contractBalance = await SystemRewardInstance.getBalance.call();
    assert.equal(contractBalance, 70e18, "the contract balance must be 70e18");
  });

  it("it can add another reward to another specific active validator : added 50 ether to validator4's map", async () => {
    await printState();
    await SystemRewardInstance.addReward(validators[4], {
      from: accounts[19],
      value: 50e18,
    });
    const reward = await SystemRewardInstance.rewardMapping.call(validators[4]);
    assert.equal(reward, 45e18, "the total reward is deducted by 5e18");
    const contractBalance = await SystemRewardInstance.getBalance.call();
    assert.equal(
      contractBalance,
      120e18,
      "the contract balance must be 120e18"
    );
  });

  it("it can distribute reward of validators to the validator itself and the delegators (this call is via the BKCValidatorSetInstance.updateValidatorSet function)", async () => {
    await printState();
    const balanceValidator1Before = await getBalance(validators[1]);
    const balanceDelegator2Before = await getBalance(delegators[2]);
    const balanceDelegator4Before = await getBalance(delegators[4]);
    await BKCValidatorSetInstance.updateValidatorSet({ from: validators[0] }); // the update validatorset will trigger the reward distribution
    // the reward must already be distributed.
    // validator1 must get 60 ether
    // delegator2 must get 2 ether
    // delegator4 must get 1 ether
    const balanceValidator1After = await getBalance(validators[1]);
    const balanceDelegator2After = await getBalance(delegators[2]);
    const balanceDelegator4After = await getBalance(delegators[4]);
    assert.equal(
      balanceValidator1After - balanceValidator1Before >= 59e18,
      true,
      "60 ether rewarded"
    );
    assert.equal(
      balanceDelegator2After - balanceDelegator2Before >= 19e17,
      true,
      "2 ether rewarded"
    );
    assert.equal(
      balanceDelegator4After - balanceDelegator4Before >= 9e17,
      true,
      "1 ether rewarded"
    );
    const reward = await SystemRewardInstance.rewardMapping.call(validators[1]);
    assert.equal(reward, 0, "the reward must now be 0");
  });

  it("it can distribute reward of another validator as well", async () => {
    await printState();
    const reward = await SystemRewardInstance.rewardMapping.call(validators[4]);
    assert.equal(reward, 0, "the reward must now be 0");
  });

  it("should able to dequeue the unbonding validator now", async () => {
    await printState();
    await ValidatorPoolInstance.removeUnBondingValidatorFromQueue({
      from: validators[2],
    });
  });

  // jail
  it("can't jail non-active validators", async () => {
    await printState();
    try {
      await ValidatorPoolInstance.jailValidator(validators[2], {
        from: validators[1],
      });
      assert.fail();
    } catch (err) {
      assert.ok(err.toString().includes("can't jail non-active validator"));
    }
  });

  it("only active validators can jail other validators : delegator3 try to jail validator4, validator2 try to jail validator1", async () => {
    await printState();
    try {
      await ValidatorPoolInstance.jailValidator(validators[4], {
        from: delegators[3],
      });
      assert.fail();
    } catch (err) {
      assert.ok(err.toString().includes("not allow for non-active validator"));
    }
    try {
      await ValidatorPoolInstance.jailValidator(validators[1], {
        from: validators[2],
      });
      assert.fail();
    } catch (err) {
      assert.ok(err.toString().includes("not allow for non-active validator"));
    }
  });

  it("validator can't jail itself : validator1 try to jail itself", async () => {
    try {
      await ValidatorPoolInstance.jailValidator(validators[1], {
        from: validators[1],
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err.toString().includes("only other validators can do this operation")
      );
    }
  });

  it("add reward to imitate jailing : 100 ether added to validator1 map", async () => {
    await printState();
    await SystemRewardInstance.addReward(validators[1], {
      from: accounts[19],
      value: 100e18,
    });
  });

  it("an active validator can only be jailed from another validator and after consensus", async () => {
    await printState();
    await ValidatorPoolInstance.jailValidator(validators[1], {
      from: validators[4],
    });
  }); // ????

  it("validator can't unbond itself when jailed", async () => {
    try {
      await ValidatorPoolInstance.removeUnBondingValidatorFromQueue({
        from: validators[1],
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err
          .toString()
          .includes(
            "can't unbond, validator is jailed, also the validator should not be in the unbonding queue"
          )
      );
    }
  });

  it("jailed validator can't get any reward (reward go to 0 first) : validator 1 is jailed, hence validator1's reward will be 0", async () => {
    await printState();
    const reward1 = await SystemRewardInstance.rewardMapping.call(
      validators[1]
    );
    assert.equal(reward1, 0, "the reward should be 0 ether");
  });

  it("the jailed validator's reward will then be distributed among other validators and the system itself at rate 90:10", async () => {
    await printState();
    const reward4 = await SystemRewardInstance.rewardMapping.call(
      validators[4]
    );
    assert.equal(reward4, 81e18, "the reward should be 81 ether");
  });

  it("the jailed validator is penalized 5e18, and is removed from the validator set", async () => {
    await printState();
    const index =
      (await ValidatorPoolInstance.validatorsMap.call(validators[1])) - 1;
    const v = await ValidatorPoolInstance.validators.call(index);
    assert.equal(v[1], 55e18, "the balance is deducted by 5 ether");
    const val = await BKCValidatorSetInstance.currentValidatorSetMap.call(
      validators[1]
    );
    assert.equal(
      val,
      0,
      "the jailed validator must not be in the validator set"
    );
  });

  it("already jailed validator can't become active for 2 days(in the test for 10 seconds) : update validator set, (4,3) not (4,1)", async () => {
    await printState();
    await BKCValidatorSetInstance.updateValidatorSet({ from: validators[0] });
    const v1 = await BKCValidatorSetInstance.currentValidatorSet.call(0);
    assert.equal(
      v1[0],
      validators[4],
      "the first validator must be validator 4"
    );
    const v2 = await BKCValidatorSetInstance.currentValidatorSet.call(1);
    assert.equal(
      v2[0],
      validators[3],
      "the second validator must be validator 3"
    );
  });

  // this also implies that delegating to a jailed validator is not possible, and withdrawal or topping-up are also not possible
  it("validator bonding status is UNBONDING but not in the validatorUnbondQueue", async () => {
    await printState();
    const queue = await ValidatorPoolInstance.validatorUnBondQueue.call(
      validators[1]
    );
    console.log(new Date(queue.toNumber() * 1000).toLocaleString());
    assert.equal(
      queue,
      0,
      "the validator must not be in the unbond queue (the time should be 0)"
    );
    const index =
      (await ValidatorPoolInstance.validatorsMap.call(validators[1])) - 1;
    const v = await ValidatorPoolInstance.validators.call(index);
    assert.equal(
      v[2],
      BondStatus.UNBONDING,
      "the bonding status should be unbonding"
    );
  });

  it("the jailed validator can't receive any further reward : try to add reward 10 ether to validator1 reward map", async () => {
    await printState();
    try {
      await SystemRewardInstance.addReward(validators[1], {
        from: validators[19],
        value: 10e18,
      });
      assert.fail();
    } catch (err) {
      assert.ok(
        err.toString().includes("can't add reward to a non-active validator")
      );
    }
  });

  it("validator can unjail itself out of the validatorJailQueue after the period : validator1 unjail itself after 10 seconds of waiting", async () => {
    await printState();
    await sleep(10 * 1000); // 10 seconds
    await ValidatorPoolInstance.removeJailValidatorFromQueue({
      from: validators[1],
    });
    const index =
      (await ValidatorPoolInstance.validatorsMap.call(validators[1])) - 1;
    const v = await ValidatorPoolInstance.validators.call(index);
    assert.equal(v[3], false, "the jail status should be false"); // not jail
    assert.equal(v[2], BondStatus.UNBONDED, ""); // not unbonding
  });

  it("ends", async () => {
    await printState();
  });
});
