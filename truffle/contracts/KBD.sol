pragma solidity <=0.8.6;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract KBD is ERC20 {
    address private admin;

    uint256 public exchangeRate = 5000;

    address payable private pool =
        payable(0x60bB6c1B4a0F5B1ea820be6c762384982b8a658c);

    modifier onlyAdmin() {
        require(msg.sender == admin, "admin only");
        _;
    }

    constructor() ERC20("Sakon", "WIN") {
        _mint(msg.sender, 1000000 * (10**uint256(decimals())));
        admin = msg.sender;
    }

    function fund() public payable {
        require(msg.value > 0, "can't fund with 0 amount");
    }

    function burn(address account, uint256 amount) public onlyAdmin {
        _burn(account, amount * (10**uint256(decimals())));
    }

    function mint(address account, uint256 amount) public onlyAdmin {
        _mint(account, amount * (10**uint256(decimals())));
    }

    function this_balance() public view onlyAdmin returns (uint256) {
        return address(this).balance;
    }

    function buy() external payable {
        require(msg.value > 0, "can't buy any with zero ETH");
        uint256 getAmount = msg.value * exchangeRate;

        if (balanceOf(pool) > getAmount) {
            _burn(pool, getAmount);
            _mint(msg.sender, getAmount);
        } else {
            _mint(msg.sender, getAmount);
            _mint(pool, getAmount);
        }
        updateExchangeRate();
    }

    function sell(uint256 amount) external {
        require(balanceOf(msg.sender) >= amount, "not enough balance to sell");
        uint256 getAmount = amount / (exchangeRate);
        address payable recipient = payable(msg.sender);
        if (address(this).balance < getAmount) {
            revert();
        }
        recipient.transfer(getAmount);
        transfer(payable(pool), amount);
    }

    function updateExchangeRate() private {
        exchangeRate = exchangeRate + exchangeRate / 10;
    }
}
