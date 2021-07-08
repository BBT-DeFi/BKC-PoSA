pragma solidity <=0.8.6;

contract Test {
    mapping(string => uint256) public pair;

    uint256 public num;

    string public name;

    address public addr;

    function setValue(string calldata key, uint256 value) external {
        pair[key] = value;
    }

    function setName(string calldata _name) external {
        name = _name;
    }

    function setAddress(address a) external {
        addr = a;
    }

    function setNum(uint256 _num) external {
        num = _num;
    }
}
